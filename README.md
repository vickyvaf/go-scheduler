# Go Email Scheduler 📧

A lightweight, minimal web application for scheduling and sending emails using Go and Resend API. Designed for simplicity and performance, the frontend uses plain HTML while the backend provides flexible handlers for both form and JSON requests.

## ✨ Features

- **Scheduled Delivery**: Specify a future date and time for emails to be sent automatically using Resend's scheduling capabilities.
- **Multiple Recipients**: Support for sending to one or more recipients via comma-separated email addresses.
- **Dual-Mode Handlers**: Backend supports both traditional HTML forms and modern JSON API submissions.
- **Resend Integration**: Seamless delivery through the high-performance [Resend](https://resend.com/) API.

## 🚀 Tech Stack

- **Backend**: [Go (Golang)](https://go.dev/)
- **Frontend**: Plain HTML5 & Vanilla JavaScript
- **Email Service**: [Resend API](https://resend.com/)

## 🛠️ Installation & Setup

### Prerequisites

- Go 1.20 or later installed.
- A Resend API Key (Get one at [resend.com](https://resend.com/)).

### Steps

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd go-scheduler
   ```

2. **Setup environment variables**
   Create a `.env` file from the example:
   ```bash
   cp .env.example .env
   ```
   Edit `.env` and fill in your details:
   ```env
   RESEND_CLIENT=your_resend_api_key
   RESEND_FROM=onboarding@resend.dev
   PORT=8080
   ```

3. **Install dependencies**
   ```bash
   go mod tidy
   ```

4. **Run the application**
   ```bash
   go run main.go
   ```
   The server will start at `http://localhost:8080`.

## 📁 Project Structure

```text
├── internal/
│   ├── handlers/    # HTTP Route handlers (Form & JSON parsing)
│   ├── models/      # Data structures for email requests
│   └── services/    # Resend API service implementation
├── main.go          # Server initialization and env loading
├── index.html       # Minimal frontend with scheduling support
├── go.mod           # Go dependencies
└── .env             # Configuration (API keys, ports)
```

## 📝 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | `/`      | Serves the main UI (index.html) |
| POST   | `/email` | Sends/Schedules an email (Accepts JSON or Form data) |

## 💡 Usage

To schedule an email, use the "Schedule At" field in the form. The system will automatically convert local browser time to the ISO 8601 UTC format required by the Resend API.

Leave the schedule field empty to send the email immediately.
