package services

import (
	"fmt"
	"go-scheduler/internal/models"
	"strings"

	"github.com/resend/resend-go/v3"
)

type EmailService struct {
	Client *resend.Client
	From   string
}

func NewEmailService(apiKey, from string) *EmailService {
	if from == "" {
		from = "Acme <onboarding@resend.dev>"
	}
	client := resend.NewClient(apiKey)
	return &EmailService{
		Client: client,
		From:   from,
	}
}

func (s *EmailService) Send(req models.EmailRequest) (string, error) {
	toList, err := s.parseRecipients(req.To)
	if err != nil {
		return "", err
	}

	params := &resend.SendEmailRequest{
		From:    s.From,
		To:      toList,
		Html:    req.Html,
		Subject: req.Subject,
	}

	sent, err := s.Client.Emails.Send(params)
	if err != nil {
		return "", fmt.Errorf("failed to send: %w", err)
	}

	return sent.Id, nil
}

func (s *EmailService) parseRecipients(to any) ([]string, error) {
	var toList []string
	switch v := to.(type) {
	case string:
		parts := strings.Split(v, ",")
		for _, p := range parts {
			trimmed := strings.TrimSpace(p)
			if trimmed != "" {
				toList = append(toList, trimmed)
			}
		}
	case []any:
		for _, p := range v {
			if s, ok := p.(string); ok {
				toList = append(toList, s)
			}
		}
	case []string:
		toList = v
	default:
		return nil, fmt.Errorf("invalid type for 'to' field")
	}

	if len(toList) == 0 {
		return nil, fmt.Errorf("recipient email is required")
	}
	return toList, nil
}
