# Student API on Go

Простой REST API для управления студентами.  
Реализован на **Go** + **Gin** + **PostgreSQL**.

---

## Технологии
- Go (Golang)
- Gin (web framework)
- GORM (ORM)
- PostgreSQL
- Docker (позже для контейнеризации)

---

## Как запустить
```bash
git clone https://github.com/USERNAME/go-student-api
cd go-student-api
go run main.go

Сервер поднимется на: http://localhost:8080

Примеры API
POST /students → создать студента
GET /students → список студентов
PUT /students/{id} → обновить данные
DELETE /students/{id} → удалить

👤 Автор
6upmanoff
[GitHub:](https://github.com/6upmanoff)
