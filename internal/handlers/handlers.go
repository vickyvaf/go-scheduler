package handlers

import (
	"encoding/json"
	"fmt"
	"go-scheduler/internal/models"
	"go-scheduler/internal/services"
	"net/http"
	"strings"
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
		// Fallback to form data
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
			return
		}
		req.To = r.FormValue("to")
		req.Subject = r.FormValue("subject")
		req.Html = r.FormValue("html")
		req.ScheduledAt = r.FormValue("scheduled_at")
	}

	// Normalize ScheduledAt to ISO 8601 if needed
	if req.ScheduledAt != "" {
		// Browser datetime-local sends "YYYY-MM-DDTHH:MM" or "YYYY-MM-DDTHH:MM:SS"
		// Resend expects ISO 8601 e.g. "2024-03-22T10:00:00Z"
		if !strings.Contains(req.ScheduledAt, "Z") && !strings.Contains(req.ScheduledAt, "+") {
			if len(req.ScheduledAt) == 16 {
				req.ScheduledAt += ":00Z"
			} else if len(req.ScheduledAt) == 19 {
				req.ScheduledAt += "Z"
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

	// For standard form submission, we can redirect or show a simple message
	w.Header().Set("Content-Type", "text/html")
	if req.ScheduledAt != "" {
		fmt.Fprintf(w, "<h2>Email berhasil dijadwalkan!</h2><p>ID: %s</p><a href='/'>Kembali</a>", emailID)
	} else {
		fmt.Fprintf(w, "<h2>Email berhasil dikirim!</h2><p>ID: %s</p><a href='/'>Kembali</a>", emailID)
	}
}
