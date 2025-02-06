package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/par4m/url-shortener/api/database"
)

type TagRequest struct {
	ShortID string `json:"shortID"`
	Tag     string `json:"tag"`
}

func AddTag(c *gin.Context) {
	var tagRequest TagRequest

	if err := c.ShouldBindJSON(&tagRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body"})
		return
	}

	shortID := tagRequest.ShortID
	tag := tagRequest.Tag

	r := database.CreateClient(0)
	defer r.Close()

	// check if shortID already exists or not
	val, err := r.Get(database.Ctx, shortID).Result()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found for the given ShortID"})
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		// if data is not a JSON object, assume it as plain string
		data = make(map[string]interface{})
		data["data"] = val
	}

	// check if "tags" field already exists - slice of string
	var tags []string
	if existingTags, ok := data["tags"].([]interface{}); ok {
		for _, t := range existingTags {
			if strTag, ok := t.(string); ok {
				tags = append(tags, strTag)
			}
		}
	}

	// check for duplicate tags
	for _, existingTag := range tags {
		if existingTag == tag {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tag already Exists"})
			return
		}
	}

	// Add new tag to the tag slice
	tags = append(tags, tag)
	data["tags"] = tags

	// Marshall updated data back to JSON
	updatedData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Marshall updated data"})
		return
	}
	err = r.Set(database.Ctx, shortID, updatedData, 0).Err()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Update Database"})
		return
	}

	// Respond with updated data
	c.JSON(http.StatusOK, data)

}
