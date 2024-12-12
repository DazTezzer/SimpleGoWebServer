package security

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.Request.Header.Get("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
            c.Abort()
            return
        }

        claims, err := ValidateToken(token)
        if err != nil {
            if claims != nil && err.Error() == "token is expired" {
                newToken, err := GenerateToken(claims.Subject)
                if err != nil {
                    c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate new token"})
                    c.Abort()
                    return
                }

                c.JSON(http.StatusUnauthorized, gin.H{
                    "error":     "Token expired",
                    "new_token": newToken,
                })
                c.Abort()
                return
            }

            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
    
        c.Set("claims", claims)
        c.Next()
    }
}