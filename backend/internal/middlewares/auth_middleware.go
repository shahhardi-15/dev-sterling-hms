package middlewares

import (
	"sterling-hms/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			utils.Unauthorized(ctx, "authorization header is required")
			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Unauthorized(ctx, "invalid authorization format")
			ctx.Abort()
			return
		}

		claims, err := utils.ValidateToken(parts[1])
		if err != nil {
			utils.Unauthorized(ctx, "invalid or expired token")
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}

func RequireRoles(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims := ctx.MustGet("claims").(*utils.JWTClaims)
		for _, r := range roles {
			if claims.Role == r {
				ctx.Next()
				return
			}
		}
		utils.Forbidden(ctx, "insufficient permissions")
		ctx.Abort()
	}
}
