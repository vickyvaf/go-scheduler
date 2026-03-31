# Go Scheduler 📧

A lightweight, modern web application for scheduling and sending emails using Go and Resend API. Featuring a clean UI built with HTMX and Tailwind CSS.

## ✨ Features

- **Modern UI**: Clean and minimal interface using Tailwind CSS v4.
- **Rich Text Editor**: Integrated Quill editor for composing HTML emails.
- **Real-time Feedback**: HTMX-powered form submission with loading states and success/error notifications.
- **Backend Stability**: Built with Go's robust standard library and modern architectural patterns.
- **Resend Integration**: Seamless email delivery via [Resend](https://resend.com/).

## 🚀 Tech Stack

- **Backend**: [Go (Golang)](https://go.dev/)
- **Frontend Interactivity**: [HTMX](https://htmx.org/)
- **Styling**: [Tailwind CSS v4](https://tailwindcss.com/)
- **Rich Text Editor**: [Quill.js](https://quilljs.com/)
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
   Copy the example environment file and fill in your Resend credentials.
   ```bash
   cp .env.example .env
   ```
   Edit `.env`:
   ```env
   RESEND_CLIENT=your_resend_api_key
   RESEND_FROM=sender@yourdomain.com
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
│   ├── handlers/    # HTTP Route handlers
│   ├── models/      # Data structures for requests/responses
│   └── services/    # Business logic (Email Service)
├── main.go          # Application entry point
├── index.html       # Single-page frontend template
└── go.mod           # Go module definition
```

## 📝 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | `/`      | Serves the main UI (index.html) |
| POST   | `/email` | Sends an email (Accepts JSON or Form via HTMX) |

