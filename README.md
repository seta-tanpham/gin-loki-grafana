# 🚀 Gin REST API with Zerolog + Loki + Grafana Dashboard

This project is a minimal REST API built using the [Gin](https://github.com/gin-gonic/gin) framework in Go. It demonstrates how to:

- Log API activity using [Zerolog](https://github.com/rs/zerolog)
- Push logs to [Loki](https://grafana.com/oss/loki/)
- View and analyze logs via [Grafana](https://grafana.com/)
- Build a dashboard to monitor API success/failure rate

---

## 🛠️ Tech Stack

- Go (Gin Framework)
- Zerolog (structured logging)
- Docker Compose
- Grafana + Loki + Promtail (log pipeline)

---

## 📦 Folder Structure

```
.
├── docker-compose.yml
├── app
│   ├── main.go
│   └── logger.go
├── promtail
│   └── config.yml
├── logs
│   └── server.log (auto-generated)
└── README.md
```

---

## ⚙️ Setup Instructions

### 1. Clone the repo

```bash
git clone https://github.com/your-org/gin-zerolog-loki-demo.git
cd gin-zerolog-loki-demo
```

### 2. Start all services

```bash
docker-compose up --build
```

### 3. Access the services

| Service   | URL                             |
|-----------|----------------------------------|
| API       | http://localhost:8080/ping      |
| Grafana   | http://localhost:3000           |
| Loki      | http://localhost:3100           |

**Default Grafana credentials:**
- **User:** `admin`
- **Password:** `admin`

---

## 🧪 Example Endpoints

- `GET /ping` — returns `pong` (Success)
- `GET /error` — simulates failure with 500
- `GET /status/:code` — customizable status code (e.g., `/status/404`)

---

## 📊 Grafana Dashboard

### Dashboard Metrics via LogQL

- ✅ Success Requests:  
  ```logql
  count_over_time({job="gin-app"} | json | status=~"2.." [1m])
  ```

- ❌ Failed Requests:  
  ```logql
  count_over_time({job="gin-app"} | json | status!~"2.." [1m])
  ```

> Adjust `[1m]` to `[5m]`, `[1h]`, etc. for different time windows.

---

## 📓 Notes

- Logs are written to `logs/server.log`, which `promtail` reads and pushes to `Loki`.
- Log format is structured JSON using `zerolog` for easy parsing in Grafana.
- Make sure Docker Desktop has access to your file system if you encounter volume issues.

---

## 📌 TODO (optional improvements)

- [ ] Add Prometheus + Gin middleware for metrics
- [ ] Add healthcheck endpoint
- [ ] Add alert rules in Grafana
- [ ] Deploy to Kubernetes with Loki Stack
