package handlers

import (
	"encoding/json"
	"fmt"
	"go-scheduler/internal/models"
	"go-scheduler/internal/services"
	"net/http"
)

type ApiHandler struct {
	Service *services.EmailService
}

func NewApiHandler(service *services.EmailService) *ApiHandler {
	return &ApiHandler{Service: service}
}

func (h *ApiHandler) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", h.IndexHandler)
	mux.HandleFunc("/index.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.js")
	})
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
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.EmailResponse{
		Message: "Email sent successfully",
		EmailID: emailID,
	})
}
