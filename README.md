
# URL Shortener (Go)

This is a simple URL shortener service built with Go. It allows users to register long URLs and receive a shortened version that redirects to the original address.

---

## 🛠 Features

- Shorten long URLs
- Redirect shortened URLs to original
- In-memory storage (easy to upgrade to a database)
- RESTful API structure

---

## 📁 Project Structure

url-shortener/
├── cmd/
│ └── main.go # Entry point
├── internal/
│ ├── handler/ # HTTP handlers
│ │ └── handler.go
│ ├── model/ # Request/response models
│ │ └── url.go
│ ├── service/ # Business logic
│ │ └── url_service.go
│ └── storage/ # In-memory storage logic
│ └── store.go
├── go.mod
├── go.sum
└── README.md





---

## 🚀 How to Run

1. Clone the repository:

```bash
git clone https://github.com/Enenche124/url-shortener.git
cd url-shortener
# URL Shortener (Go)

This is a simple URL shortener service built with Go. It allows users to register long URLs and receive a shortened version that redirects to the original address.

---

## 🛠 Features

- Shorten long URLs
- Redirect shortened URLs to original
- In-memory storage (easy to upgrade to a database)
- RESTful API structure

---

## 📁 Project Structure

url-shortener/  
├── cmd/  
│   └── main.go # Entry point  
├── internal/  
│   ├── handler/ # HTTP handlers  
│   │   └── handler.go  
│   ├── model/ # Request/response models  
│   │   └── url.go  
│   ├── service/ # Business logic  
│   │   └── url_service.go  
│   └── storage/ # In-memory storage logic  
│       └── store.go  
├── go.mod  
├── go.sum  
└── README.md  

---

## 🚀 How to Run

1. Clone the repository:

```bash
git clone https://github.com/Enenche124/url-shortener.git
cd url-shortener
