package routes

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/par4m/url-shortener/api/database"
	"github.com/par4m/url-shortener/api/models"
	"github.com/par4m/url-shortener/api/utils"
)

func ShortenURL(c *gin.Context) {
	var body models.Request

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse JSON"})
		return
	}

	r2 := database.CreateClient(1)
	defer r2.Close()

	val, err := r2.Get(database.Ctx, c.ClientIP()).Result()
	if err == redis.Nil {
		if err := r2.Set(database.Ctx, c.ClientIP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set initial rate limit"})
			return
		}
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error accessing rate limit data"})
		return
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error converting rate limit value"})
		return
	}

	if valInt <= 0 {
		limit, err := r2.TTL(database.Ctx, c.ClientIP()).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting TTL for rate limit"})
			return
		}
		c.JSON(http.StatusTooManyRequests, gin.H{
			"error":            "Rate limit exceeded",
			"rate_limit_reset": limit / time.Nanosecond / time.Minute,
		})
		return
	}

	if !govalidator.IsURL(body.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	if utils.IsDifferentDomain(body.URL) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cross-domain URL shortening is not allowed"})
		return
	}

	body.URL = utils.EnsureHTTPPrefix(body.URL)
	var id string
	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	r := database.CreateClient(0)
	defer r.Close()

	val, _ = r.Get(database.Ctx, id).Result()
	if val != "" {
		c.JSON(http.StatusConflict, gin.H{"error": "URL custom short already exists"})
		return
	}

	if body.Expiry == 0 {
		body.Expiry = 24
	}

	if err := r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to connect to the Redis server"})
		return
	}

	resp := models.Response{
		Expiry:          body.Expiry,
		XRateLimitReset: 30,
		XRateRemaining:  10,
		URL:             body.URL,
		CustomShort:     os.Getenv("DOMAIN") + "/" + id,
	}

	r2.Decr(database.Ctx, c.ClientIP())
	val, _ = r2.Get(database.Ctx, c.ClientIP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)

	ttl, _ := r2.TTL(database.Ctx, c.ClientIP()).Result()
	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute

	c.JSON(http.StatusOK, resp)
}
