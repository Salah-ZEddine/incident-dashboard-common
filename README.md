# Incident-Dashboard-Common

**Purpose:**
Shared Go models and validation logic for the IncidentPilot microservices.

---

## ⚙️ Tech Stack
- Go 1.24+

---

## 🧱 Features
- Common data models (Log, Alert)
- Log validation utilities
- Used by log-ingestor, log-processor, and incident-dashboard

---

## 📦 Usage
Import as a Go module in your microservices:
```go
import "github.com/your-org/incident-dashboard-common/models"
```

---

## 📁 Structure
- `models/Alert.go`   — Alert struct
- `models/log.go`     — Log struct
- `models/validate_log.go` — Log validation

---

## 🤝 Contributing
- Fork, branch, PR
- Write tests for new features
- Follow Go best practices

---

## 📄 License
MIT (or specify) 