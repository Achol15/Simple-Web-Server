# Simple Go Web Server

This repository contains a simple web server written in Go. The server listens on port 9090 and responds with "Hi! Everyone Whatsup? Welcome to my very first golang web server." to HTTP requests.

## Features

- Listens on port `9090`.
- Handles HTTP GET requests.
- Responds with a "Hi! Everyone Whatsup? Welcome to my very first golang web server." message.

---

## Prerequisites

- [Go](https://golang.org/) (version 1.16 or later recommended)

---

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/Achol15/simple-go-web-server.git
cd simple-go-web-server
```

### 2. Run the Web Server
```
go run main.go
```

### Access the Web Server
Open a web browser and visit:
```
http://localhost:9090
```
You will see the following response:
```
Hi! Everyone Whatsup? Welcome to my very first golang web server.
```
### Customization
* Change Port: Replace 8080 in http.ListenAndServe(":9090", nil).
* Modify Response: Edit the handler function.

### Troubleshooting
If port 9090 is in use, kill the process or use a different port.
### License
This project is licensed under the MIT License. See the LICENSE file for details.
### Author
# Jerin Nur Khan Achol
