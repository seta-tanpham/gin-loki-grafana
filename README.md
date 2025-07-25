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
├── Dockerfile
├── main.go
├── go.mod
├── go.sum
├── promtail-config.yml
├── logs
│   └── server.log (auto-generated)
└── README.md
```

---

## ⚙️ Setup Instructions

### 1. Clone the repo

```bash
git clone https://github.com/seta-tanpham/gin-loki-grafana.git
cd gin-loki-grafana
```

### 2. Start all services

```bash
#first time
docker compose up --build

#rerun
docker compose up

#if using docker compose v1 then the command should be docker-compose
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

- `GET /ping` — returns `pong` (Success 200)
- `GET /fail` — simulates failure with 500
- `GET /random_value` — simulates not found with 404

---

## 📊 Grafana Dashboard

### Dashboard Metrics via LogQL

--Config data source
Loki URL: http://loki:3100

--Import dashboard
Use `grafana-import-1753270140645.json`

--Refresh each panel in case of errors

```logql
✅ Total Success Requests (200):  
  
sum by(job) (
  count_over_time({job="gin-app"} | json | status=~"2.." [$__range])
)
or vector(0)
```

```logql
❌ Total Failed Requests (500):
sum(
  count_over_time(
    {job="gin-app"} 
    | json 
    | status=~"5.." 
    [$__range]
  )
)
or vector(0)
```
```logql
✅ Success Rate:
(
  sum(
    count_over_time(
      {job="gin-app"} 
      | json 
      | status=~"2.." 
      [$__range]
    )
  )
  /
  sum(
    count_over_time(
      {job="gin-app"} 
      | json 
      | method="GET" 
      [$__range]
    )
  )
) * 100
```

> Adjust the time picker for different time windows.

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
