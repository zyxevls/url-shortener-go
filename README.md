# URL Shortener Service

Production-ready URL Shortener built with **Golang + PostgreSQL + Redis** using **Clean Architecture** principles.

---

# Features

## ✅ Core Features

* Generate short URL
* Redirect to original URL
* Custom alias support
* Expired link handling
* Click tracking

## ⚡ Advanced Features

* Redis caching (fast redirect)
* Rate limiting (anti spam)
* Async click tracking
* Clean Architecture (scalable & maintainable)

---

# 🏗️ Tech Stack

| Layer        | Technology         |
| ------------ | ------------------ |
| Backend      | Golang             |
| Database     | PostgreSQL         |
| Cache        | Redis              |
| Architecture | Clean Architecture |

---

# 📁 Project Structure

```
url-shortener/
├── cmd/
│   └── main.go
│
├── internal/
│   ├── domain/
│   ├── usecase/
│   ├── delivery/http/
│   ├── repository/
│   │   ├── postgres/
│   │   └── redis/
│   └── config/
│
├── pkg/utils/
└── go.mod
```

---

# ⚙️ Setup & Installation

## 1. Clone Project

```bash
git clone https://github.com/zyxevls/url-shortener-go.git
cd url-shortener
```

## 2. Install Dependencies

```bash
go mod tidy
```

## 3. Setup PostgreSQL

Create database:

```sql
CREATE DATABASE urlshort;
```

Create table:

```sql
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    short_code VARCHAR(20) UNIQUE NOT NULL,
    custom_alias VARCHAR(50),
    click_count INT DEFAULT 0,
    expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 4. Setup Redis

Run Redis locally:

```bash
docker run -d -p 6379:6379 redis
```

## 5. Run App

```bash
go run cmd/main.go
```

Server running at:

```
http://localhost:8080
```

---

# 🌐 API Documentation

## 🔹 Create Short URL

**POST** `/api/v1/shorten`

### Request Body

```json
{
  "url": "https://google.com",
  "custom_alias": "zay",
  "expire_at": "2026-12-01"
}
```

### Response

```json
{
  "short_url": "http://localhost:8080/zay"
}
```

---

## 🔹 Redirect

**GET** `/{code}`

### Example

```
http://localhost:8080/zay
```

### Behavior

* Redirect to original URL
* Return **404** if not found
* Return **410** if expired

---

## 🔹 Rate Limit

* Max **10 requests/minute per IP**
* Exceed → `429 Too Many Requests`

---

# ⚡ System Flow

## 🔹 Create URL Flow

```
Client → API → Usecase → PostgreSQL
                         ↓
                      Redis Cache
```

---

## 🔹 Redirect Flow

```
Client → API
        ↓
    Rate Limit (Redis)
        ↓
    Check Cache (Redis)
        ↓
   Hit → Redirect ⚡
        ↓
   Miss → PostgreSQL
        ↓
   Cache → Redirect
        ↓
   Increment Click (Async)
```

---

# 🧠 Cache Strategy

## URL Cache

```
Key   : short:<code>
Value : original_url
TTL   : follow expired time
```

## Click Counter

```
Key   : click:<code>
Value : increment
```

## Rate Limit

```
Key   : rate:<ip>
TTL   : 60s
Limit : 10 req/min
```

---

# 🔥 Clean Architecture Flow

```
Handler → Usecase → Repository → Database
                 ↓
               Redis
```

---

# 🧪 Testing

## Create URL

```bash
curl -X POST http://localhost:8080/api/v1/shorten \
-H "Content-Type: application/json" \
-d '{"url":"https://google.com"}'
```

## Access URL

```
http://localhost:8080/abc123
```

---

# 📊 Example Scenario

1. User creates short link
2. Data stored in PostgreSQL
3. Cached in Redis
4. User accesses link
5. Redirect happens instantly (Redis)
6. Click counted async

---

# 🚀 Future Improvements

* JWT Authentication
* User dashboard (React)
* QR Code generator
* Analytics (daily clicks)
* Geo tracking
* Distributed system (Kafka)

---

# 🧑‍💻 Author

Built for learning & production-ready backend practice 🚀

---

# ⭐ Notes

This project demonstrates:

* Clean Architecture implementation
* Scalable backend design
* High-performance caching strategy

---

