package middleware

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CaptchaMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			logger.Error("Error reading body", zap.Error(err))
			c.JSON(400, gin.H{"error": "Unable to read request body"})
			c.Abort()
			return
		}

		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		var body struct {
			CaptchaCompleted bool `json:"captcha_completed"`
		}

		if err := json.Unmarshal(bodyBytes, &body); err != nil {
			logger.Error("Error unmarshalling JSON", zap.Error(err))
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			c.Abort()
			return
		}

		if !body.CaptchaCompleted {
			logger.Warn("CAPTCHA not completed")
			c.JSON(401, gin.H{"error": "CAPTCHA not completed"})
			c.Abort()
			return
		}

		c.Next()
	}
}
