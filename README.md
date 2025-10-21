# 🚀 GoPost API

[![](https://img.shields.io/badge/Go-1.2x-blue?style=for-the-badge&logo=go)](https://go.dev/)
[![](https://img.shields.io/badge/Fiber-v2-cyan?style=for-the-badge&logo=go)](https://gofiber.io/)
[![](https://img.shields.io/badge/Docker-Build-blue?style=for-the-badge&logo=docker)](https://www.docker.com/)
[![](https://img.shields.io/badge/Prometheus-Monitoring-orange?style=for-the-badge&logo=prometheus)](https://prometheus.io/)
[![](https://img.shields.io/badge/Grafana-Dashboard-brightgreen?style=for-the-badge&logo=grafana)](https://grafana.com/)
[![](https://img.shields.io/badge/PostgreSQL-Database-blue?style=for-the-badge&logo=postgresql)](https://www.postgresql.org/)

GoPost is a high-performance backend API built with **Go (Fiber)**, featuring user authentication, post management, and a complete, pre-configured monitoring stack.

The project is fully containerized with Docker and uses Google's Wire for dependency injection, JWT for authentication, and Air for a seamless live-reload development experience.

## ✨ Features

* **Full User Authentication:** Secure registration and login using **JWT** (JSON Web Tokens).
* **Post Management:** Full CRUD (Create, Read, Update, Delete) operations for posts.
* **Dependency Injection:** Clean and maintainable code using **Google's Wire**.
* **Live-Reload:** Automatic hot-reloading in development using **Air**.
* **Fully Containerized:** Includes a multi-service `docker-compose.yaml` for the app, database, and monitoring.
* **Integrated Monitoring:**
    * **Prometheus:** Pre-configured to scrape application metrics.
    * **Grafana:** Pre-configured with a provisioned data source to visualize Prometheus data.
* **Health & Metrics:** Dedicated `/healthcheck` and `/metrics` endpoints.
* **Secure by Default:** Includes security-focused middlewares for CORS and headers.

## 🖼️ Screenshots

Here's a look at the monitoring stack in action.

### Grafana Dashboard
![Grafana Dashboard Screenshot](/.github/assets/grafana.png "Grafana Dashboard")

### API in Postman
![API Request Screenshot](/.github/assets/api-example.png "API Example")

## 🛠️ Tech Stack

* **Backend:** Go (Golang)
* **Framework:** Fiber
* **Database:** PostgreSQL 15
* **Authentication:** JWT (JSON Web Tokens)
* **Dependency Injection:** Google Wire
* **Monitoring:** Prometheus & Grafana
* **Containerization:** Docker & Docker Compose
* **Development Tools:** Air (Hot Reload), Makefile

## 🏁 Getting Started (Docker)

The recommended way to run this project is with Docker Compose, which handles the application, database, and monitoring stack all at once.

### Prerequisites

* [Git](https://git-scm.com/downloads)
* [Docker](https://www.docker.com/get-started) & [Docker Compose](https://docs.docker.com/compose/install/)

---

### 1. Clone the Repository

```sh
git clone [https://github.com/SamedArslan28/GoPost.git](https://github.com/SamedArslan28/GoPost.git)
cd GoPost