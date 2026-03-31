package handlers

import (
	"encoding/json"
	"fmt"
	"go-scheduler/internal/models"
	"go-scheduler/internal/services"
	"net/http"
	"strings"
	"time"
)

type ApiHandler struct {
	Service *services.EmailService
}

func NewApiHandler(service *services.EmailService) *ApiHandler {
	return &ApiHandler{Service: service}
}

func (h *ApiHandler) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", h.IndexHandler)
	mux.HandleFunc("/email", h.EmailHandler)
}

func (h *ApiHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func (h *ApiHandler) EmailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.EmailRequest
	contentType := r.Header.Get("Content-Type")

	if strings.Contains(contentType, "application/json") {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
			return
		}
		req.To = r.FormValue("to")
		req.Subject = r.FormValue("subject")
		req.Html = r.FormValue("html")
		req.ScheduledAt = r.FormValue("scheduled_at")
	}

	if req.ScheduledAt != "" {
		customLayout := "02-01-2006 15:04"
		isoLayout := "2006-01-02T15:04"

		loc := time.FixedZone("WIB", 7*3600)

		var t time.Time
		var err error

		t, err = time.ParseInLocation(customLayout, req.ScheduledAt, loc)
		if err != nil {
			t, err = time.ParseInLocation(isoLayout, req.ScheduledAt, loc)
		}

		if err == nil {
			req.ScheduledAt = t.UTC().Format(time.RFC3339)
		} else {
			if !strings.Contains(req.ScheduledAt, "Z") && !strings.Contains(req.ScheduledAt, "+") {
				if len(req.ScheduledAt) == 16 {
					req.ScheduledAt += ":00Z"
				} else if len(req.ScheduledAt) == 19 {
					req.ScheduledAt += "Z"
				}
			}
		}
	}

	emailID, err := h.Service.Send(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send email: %v", err), http.StatusInternalServerError)
		return
	}

	if r.Header.Get("HX-Request") == "true" {
		w.Header().Set("Content-Type", "text/html")
		if req.ScheduledAt != "" {
			fmt.Fprintf(w, "Email berhasil dijadwalkan!")
		} else {
			fmt.Fprintf(w, "Email berhasil dikirim!")
		}
		return
	}


	w.Header().Set("Content-Type", "text/html")
	if req.ScheduledAt != "" {
		fmt.Fprintf(w, "<h2>Email berhasil dijadwalkan!</h2><p>ID: %s</p><a href='/'>Kembali</a>", emailID)
	} else {
		fmt.Fprintf(w, "<h2>Email berhasil dikirim!</h2><p>ID: %s</p><a href='/'>Kembali</a>", emailID)
	}
}
