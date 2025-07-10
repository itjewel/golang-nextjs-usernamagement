package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/md-jewel-mia/user-service/internal/service" // ✅ Update this path to match your go.mod module
)

type UserHandler struct {
    service service.UserService
}

// Constructor function for the handler
func NewUserHandler(s service.UserService) *UserHandler {
    return &UserHandler{service: s}
}

// HTTP GET: /users → returns list of users
func (h *UserHandler) GetUsers(c *gin.Context) {
    users := h.service.GetAll()
    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "data":    users,
    })
}
