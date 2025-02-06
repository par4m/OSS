package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/par4m/url-shortener/api/database"
)

func GetByShortID(c *gin.Context) {
	shortID := c.Param("shortID")

	r := database.CreateClient(0)

	defer r.Close()
	val, err := r.Get(database.Ctx, shortID).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found for given shortID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": val})

}
