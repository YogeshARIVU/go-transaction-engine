# 💳 Go Transaction Processing Engine

A simple backend service built in Go that simulates a card transaction system.

## 🚀 Features

* Card validation (ACTIVE / BLOCKED)
* Secure PIN authentication (SHA-256 hashing)
* Transaction processing (Withdraw / Topup)
* Balance tracking
* Transaction history logging
* REST APIs

---

## 🛠️ Tech Stack

* Go (Golang)
* Standard HTTP library
* In-memory storage (map)

---

## 📂 Project Structure

```
cmd/
internal/
  ├── handler/
  ├── model/
  ├── repository/
  ├── service/
  └── util/
```

---

## ▶️ Run Locally

```bash
go mod tidy
go run ./cmd
```

Server runs on:

```
http://localhost:8080
```

---

## 📡 API Endpoints

### 🔹 Transaction

POST `/api/transaction`

```json
{
  "cardNumber": "4123456789012345",
  "pin": "1234",
  "type": "withdraw",
  "amount": 200
}
```

---

### 🔹 Get Balance

GET `/api/card/balance/{cardNumber}`

---

### 🔹 Transaction History

GET `/api/card/transactions/{cardNumber}`

---

## 🧪 Sample Curl

```bash
curl -X POST http://localhost:8080/api/transaction \
-H "Content-Type: application/json" \
-d '{"cardNumber":"4123456789012345","pin":"1234","type":"withdraw","amount":200}'
```

---

## 🔐 Security

* PIN stored using SHA-256 hashing
* No plaintext PIN storage
* No PIN logging

---


