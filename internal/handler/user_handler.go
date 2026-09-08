package handler

import (
	"github.com/gin-gonic/gin"
	"internal/service"
	"net/http"
	"strconv"
	"internal/model"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.userService.CreateUser(&user); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
		c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
		c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	user, err := h.userService.GetUserByID(uint(id))
	if err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
		c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	var user model.User
	if err := c.ShouldBindJSON(&user); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = uint(id)
	if err := h.userService.UpdateUser(&user); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
		c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	if err := h.userService.DeleteUser(uint(id)); err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
		c.JSON(http.StatusNoContent, nil)
}