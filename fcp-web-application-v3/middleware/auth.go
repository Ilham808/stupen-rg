package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		tokenStr, err := ctx.Cookie("session_token")
		content := ctx.GetHeader("Content-Type")
		if err != nil {
			if err == http.ErrNoCookie && content == "application/json" {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			} else {
				ctx.JSON(http.StatusSeeOther, gin.H{"error": "unauthorized"})
			}
			ctx.Abort()
			return
		}

		claims := model.Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return model.JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token signature"})
				ctx.Abort()
				return
			}

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}

		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Next()
	})
}
