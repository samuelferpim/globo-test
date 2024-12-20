package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

func CaptchaMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			c.JSON(400, gin.H{"error": "Unable to read request body"})
			c.Abort()
			return
		}

		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		var body struct {
			CaptchaCompleted bool `json:"captcha_completed"`
		}

		if err := json.Unmarshal(bodyBytes, &body); err != nil {
			log.Printf("Error unmarshalling JSON: %v", err)
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			c.Abort()
			return
		}

		if !body.CaptchaCompleted {
			c.JSON(401, gin.H{"error": "CAPTCHA not completed"})
			c.Abort()
			return
		}

		c.Next()
	}
}