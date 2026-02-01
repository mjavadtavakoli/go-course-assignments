# go-course-assignments 

This repository contains exercises, examples, and challenges completed during a comprehensive Go (Golang) learning course.
The project is organized in chapters and sections for a step-by-step learning experience.

⸻

🐳 Dockerized Projects

(If Dockerfile is included)

docker build -t app-name .
docker run -p 8080:8080 app-name


⸻


🧑‍💻 Author

Mohammad Javad Tavakoli

⸻

Project Structure

The repository is organized into chapters, each containing exercises and challenges:
```bash 
├── chapter1-section6
│   ├── challenge-ecommerce-orders
│   │   ├── Dockerfile
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── handlers
│   │   │   ├── orders_test.go
│   │   │   └── orders.go
│   │   ├── main.go
│   │   ├── Makefile
│   │   ├── models
│   │   │   └── order.go
│   │   ├── README.md
│   │   └── utils
│   │       └── validator.go
│   ├── exercise-01
│   │   └── helloword.go
│   ├── exercise-02
│   │   ├── go-vet.go
│   │   └── issue-AND-fix.go
│   ├── exercise-03
│   │   └── movie-api
│   │       ├── go.mod
│   │       ├── go.sum
│   │       ├── main.go
│   │       └── README.md
│   ├── exercise-04
│   │   ├── go.mod
│   │   ├── hello-cli
│   │   ├── hello-cli-linux
│   │   ├── hello-cli-macos-arm
│   │   ├── hello-cli-macos-intel
│   │   ├── hello-cli-windows.exe
│   │   ├── mian.go
│   │   └── README.md
│   └── exercise-05
│       ├── main.go
│       └── Makefile
├── chapter2-section6
│   ├── challenge-01
│   │   └── main.go
│   ├── challenge-02
│   │   └── main.go
│   ├── challenge-03
│   │   └── main.go
│   ├── exercise-01
│   │   └── main.go
│   ├── exercise-02
│   │   └── main.go
│   ├── exercise-03
│   │   └── main.go
│   └── exercise-04
│       └── main.go
├── chapter3-section5
│   └── main.go
├── chapter3-section6
│   ├── challenge-01
│   │   └── main.go
│   ├── challenge-02
│   │   └── main.go
│   ├── exercise-01
│   │   └── main.go
│   ├── exercise-02
│   │   └── main.go
│   └── exercise-03
│       └── main.go
├── chapter4
│   └── main.go
├── chapter5-section6
│   ├── challenge-01
│   │   └── main.go
│   ├── challenge-02
│   │   └── main.go
│   ├── challenge-03
│   │   └── main.go
│   ├── challenge-04
│   │   └── main.go
│   ├── exercise-01
│   │   └── main.go
│   ├── exercise-02
│   │   └── main.go
│   ├── exercise-03
│   │   └── main.go
│   ├── exercise-04
│   │   └── main.go
│   ├── exercise-05
│   │   └── main.go
│   └── exercise-06
│       └── main.go
└── README.md

39 directories, 52 files
```
⸻

📚 Prerequisites
	•	Go 1.21+
	•	Git
	•	Docker (optional)

⸻

✍️ Final Note

This repository is not just a collection of exercises; it represents a practical Go learning path — from simple main.go files to a fully working API.

It is ideal for learners aiming to build backend applications, microservices, or production-ready projects in Go.

