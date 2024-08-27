package infrastructure

import (
	domain "LoanAPI/LoanTrackerAPI-Go/Domain"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth_middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		headers := strings.Split(header, " ")
		if len(headers) != 2 || headers[0] != "bearer" {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{
				Message: "Authorization header is not valid",
				Status:  http.StatusBadRequest,
			})
			c.Abort()
			return
		} else {

			user_id, err := VerifyRefreshToken(headers[1])
			if err == nil {
				c.Set("user_id", user_id)
				c.Next()
			} else {
				c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
					Message: err.Error(),
					Status:  http.StatusInternalServerError,
				})
				c.Abort()
				return
			}
		}
	}
}
