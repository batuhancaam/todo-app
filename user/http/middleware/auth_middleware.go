package middleware

import (
	"net/http"
	"strings"

	"github.com/batuhancaam/todo-app/helper"
	"github.com/batuhancaam/todo-app/user/service"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	userService service.UserService
}

// TODO: Refactor
func NewAuthMiddleware(userService service.UserService) gin.HandlerFunc {
	return (&AuthMiddleware{
		userService: userService,
	}).Handle
}

func (m *AuthMiddleware) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	aheaderParts := strings.Split(authHeader, " ")
	if len(aheaderParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if aheaderParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := m.userService.ParseToken(c.Request.Context(), aheaderParts[1])

	helper.ErrorPanic(err)

	c.Set("user", user)

}
