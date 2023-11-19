# Go Clean Project Layout

Go (Golang) Backend Project Layout with Gin and SQLite

## Layers of the project

- Router
- Controller
- Service
- Repository
- Domain

![Go Clean Project Layout](https://github.com/nevzatcirak/go-gin-project-layout/blob/main/diagram.png?raw=true)

### The Project Folder Structure

```
.
├── api
│   ├── controller
│   │   ├── user_controller.go
│   │   └── video_controller.go
│   └── route
│       ├── ui_route.go
│       ├── user_route.go
│       ├── route.go
│       └── video_route.go
├── bootstrap
│   ├── app.go
│   ├── database.go
│   └── env.go
├── cmd
│   └── main.go
├── domain
│   ├── error_response.go
│   ├── user.go
│   └── video.go
├── go.mod
├── go.sum
├── internal
│   └── .gitkeep
├── middleware
│   └── logger.go
├── mongo
│   └── mongo.go
├── repository
│   ├── video_repository.go
│   └── user_repository.go
├── service
│   ├── user_service.go
│   └── video_service.go
├── template
│   ├── css
│   │   └── index.css
│   ├── footer.html
│   ├── header.html
│   └── index.html
└── validator
    └── validators.go
```

### If this project helps you in anyway, show your love ❤️ by putting a ⭐ on this project ✌️

#### This project layout is inspired by [Go backend clean architecture](https://amitshekhar.me/blog/go-backend-clean-architecture)

