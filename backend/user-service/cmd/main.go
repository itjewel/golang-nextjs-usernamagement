package main

import (
    "log"
    "user-service/internal/handler"
    "user-service/internal/repository"
    "user-service/internal/service"
    "user-service/pkg/server"
)

func main() {
    repo := repository.NewUserRepository()
    svc := service.NewUserService(repo)
    h := handler.NewUserHandler(svc)

    srv := server.NewServer()
    srv.RegisterRoutes(h)
    
    log.Fatal(srv.Run(":8080"))
}
