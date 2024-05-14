package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

var jwtKey = []byte("secret")

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		tokenStr, err := ctx.Cookie("session_token")
		// if err == http.ErrNoCookie {
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		// 	ctx.Abort()
		// 	return
		// }

		if err != nil {
			if ctx.GetHeader("Content-Type") == "application/json" {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			} else {

				ctx.JSON(http.StatusSeeOther, gin.H{"error": "unauthorized"})
			}
			ctx.Abort()
			return
		}

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtKey, nil
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

		ctx.Set("id", claims.UserID)
		ctx.Next()
	})
}
