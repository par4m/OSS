package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/par4m/url-shortener/api/database"
	"github.com/par4m/url-shortener/api/models"
)

func EditURl(c *gin.Context) {

	shortID := c.Param("shortID")

	var body models.Request

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot Pass JSON"})
	}

	r := database.CreateClient(0)
	defer r.Close()

	// check if the short id exists or not
	val, err := r.Get(database.Ctx, shortID).Result()
	if err != nil || val == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short Not Found"})
		return
	}
	r.Set(database.Ctx, shortID, body.URL, body.Expiry*3600*time.Second).Err()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update the shortened content"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "content has been updated"})

}
