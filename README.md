# 🛒 Kharido  
A lightweight, modular Go-based backend template for building e‑commerce or transactional applications.

Kharido is a clean, structured Go project designed with separation of concerns in mind.  
It provides a ready‑to‑extend foundation with controllers, models, routing, configuration, and utility layers.

---

## 📁 Project Structure
kharido/ │ ├── cmd/            # Application entry point(s) ├── config/         # Configuration loading (env, constants, settings) ├── controllers/    # HTTP handlers for various features ├── models/         # Data models and domain structures ├── routes/         # Route definitions and API grouping ├── store/          # Database/store layer (DB operations, repositories) ├── utils/          # Helper utilities (logging, responses, validation, etc.) │ └── go.mod          # Go module definitio

This structure follows a clean, maintainable pattern suitable for production APIs.

---

## 🚀 Features

- **Modular folder architecture** for scalable development  
- **Controllers + Routes** pattern for clean API design  
- **Models layer** for domain-driven struct definitions  
- **Store layer** for database abstraction  
- **Config package** for centralized configuration  
- **Utility helpers** for common tasks  
- Fully written in **Go (100%)**

---

## 🧩 Tech Stack

| Component | Technology |
|----------|------------|
| Language | Go (Golang) |
| Architecture | Modular, MVC-inspired |
| API Style | REST-ready structure |
| Config | Environment-based configuration |

---

## ▶️ Getting Started

### 1️⃣ Clone the repository

```bash
git clone https://github.com/itsabhipro/kharido.git
cd kharido
