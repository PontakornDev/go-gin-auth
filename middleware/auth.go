package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PontakornDev/ginAuth/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	jwt.StandardClaims
}

// GenerateToken generates a JWT token for the given username
func GenerateToken(user *models.Users) (string, error) {
	var secretKey []byte = []byte(os.Getenv("SECRET_KEY"))
	expirationTime := time.Now().Add(5 * time.Minute) // Token expiration time

	claims := &Claims{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// fmt.Println("secretKey => ", secretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken validates the given token and returns the claims if valid
func ValidateToken(tokenString string) (*Claims, error) {
	var secretKey []byte = []byte(os.Getenv("SECRET_KEY"))
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract JWT token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		tokenString = tokenString[len("Bearer "):]
		// Validate the token
		claims, err := ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Token is valid, you can now use the claims
		c.Set("username", claims.Username)
		c.Next()
	}
}
