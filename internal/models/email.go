package models

type EmailRequest struct {
	To      any    `json:"to"`
	Subject string `json:"subject"`
	Html    string `json:"html"`
}

type EmailResponse struct {
	Message string `json:"message,omitempty"`
	EmailID string `json:"email_id,omitempty"`
}
