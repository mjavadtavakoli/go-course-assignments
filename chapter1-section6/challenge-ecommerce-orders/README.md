# Ecommerce Orders API 🛒

A simple e-commerce order management system built with Go and Gin.
This project demonstrates a clean Go project structure, RESTful APIs,
unit testing, automation with Makefile, and Docker support.

---

## ✨ Features

- RESTful API using Gin
- Create and list orders (CRUD-ready)
- Input validation
- Unit tests with high coverage
- Makefile for automation
- Dockerized application


---

## 📁 Project Structure
```
.
├── Dockerfile
├── go.mod
├── go.sum
├── handlers
│   ├── orders_test.go
│   └── orders.go
├── main.go
├── Makefile
├── models
│   └── order.go
├── README.md
└── utils
    └── validator.go

4 directories, 10 files
```

---

▶️ Run Project

```
go run main.go
```

or using Makefile:
```
make run
```

---

## 🛠 Tech Stack

- Go (Golang)
- Gin Web Framework
- Go testing (`testing`, `httptest`)
- Makefile
- Docker
  
