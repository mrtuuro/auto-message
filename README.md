# Auto-Messager

A self-contained Go service that:

* polls MongoDB every **2 minutes**,
* pulls **two** unsent SMS records,
* POSTs them to a configurable webhook (mock SMS provider),
* marks rows as `sent=true` so they never resend,
* exposes HTTP endpoints to **start/stop** the loop and **list** sent messages,
* (bonus) caches `messageId → sent_at` in Redis,
* ships with **Swagger UI**, a health-check, and Docker-Compose for local dev.

---

## Features

| Feature                           | Endpoint / Mechanism                  |
|-----------------------------------|---------------------------------------|
| Health-check                      | `GET /v1/healthz`                     |
| Start auto-sender (2 min ticker)  | `POST /v1/auto/start`                 |
| Stop auto-sender                  | `POST /v1/auto/stop`                  |
| List sent messages (paginated)    | `GET /v1/messages/list?limit=&offset=`|
| Liveness / Swagger docs           | `/swagger/index.html`                 |
| Graceful shutdown                 | SIGINT → stop ticker → stop HTTP      |

---


## 🔧 Tech stack

| Layer      | Lib / Image                               |
|------------|-------------------------------------------|
| Go         | v1.24
| Web        | [Echo v4](https://echo.labstack.com/)     |
| DB         | MongoDB 7                                 |
| Driver     | `go.mongodb.org/mongo-driver/v2`          |
| Swagger    | `swaggo/echo-swagger` + `swag` CLI        |

---

## Quick start (Docker)

```bash
git clone https://github.com/mrtuuro/auto-message.git
cd auto-message

touch .env                    # adjust system environments

# PORT=:<enter your port>
# MONGO_URI=mongodb+srv://<user>:<pass>@<host>/?retryWrites=true&w=majority
# DATABASE_NAME=<database-name>
# COLLECTION_NAME=<collection-name>
# SECRET_KEY=<enter your jwt secret key>
# WEBHOOK_URL=<your webhook url>
# WEBHOOK_KEY=<your webhook key>
vim .env

# additional to see the Makefile commands
make help

# run the service
make run                      # regenerates Swagger, builds, starts on <your-port>
