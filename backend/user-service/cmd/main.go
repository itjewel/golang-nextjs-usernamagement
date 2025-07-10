package main

import (
    "log"

    "github.com/md-jewel-mia/user-service/internal/handler"
    "github.com/md-jewel-mia/user-service/internal/repository"
    "github.com/md-jewel-mia/user-service/internal/service"
    "github.com/md-jewel-mia/user-service/pkg/server"
)

func main() {
    repo := repository.NewUserRepository()
    svc := service.NewUserService(repo)
    h := handler.NewUserHandler(svc)

    srv := server.NewServer()
    srv.RegisterRoutes(h)

    log.Fatal(srv.Run(":8080"))
}
