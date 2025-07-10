package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "user-service/internal/service"
)

type UserHandler struct {
    service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
    return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
    users := h.service.GetAll()
    c.JSON(http.StatusOK, users)
}
