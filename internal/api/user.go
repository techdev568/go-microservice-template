package api

import (
	"net/http"

	"github.com/techdev568/go-microservice-template/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *gin.Engine, db *gorm.DB) {
	users := r.Group("/users")
	{
		users.POST("/", func(c *gin.Context) {
			var user models.User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := db.Create(&user).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, user)
		})

		users.GET("/", func(c *gin.Context) {
			var users []models.User
			db.Find(&users)
			c.JSON(http.StatusOK, users)
		})

		users.PUT("/:id", func(c *gin.Context) {
			var user models.User
			id := c.Param("id")
			if err := db.First(&user, id).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			db.Save(&user)
			c.JSON(http.StatusOK, user)
		})

		users.DELETE("/:id", func(c *gin.Context) {
			id := c.Param("id")
			if err := db.Delete(&models.User{}, id).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
		})
	}
}
