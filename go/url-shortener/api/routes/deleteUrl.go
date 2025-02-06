package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/par4m/url-shortener/api/database"
)

func DeleteUrl(c *gin.Context) {

	shortID := c.Param("shortID")

	r := database.CreateClient(0)
	defer r.Close()

	err := r.Del(database.Ctx, shortID).Err()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to Delete SHortened Link"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})

}
