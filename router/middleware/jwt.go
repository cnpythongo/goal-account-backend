package middleware

import (
	"github.com/cnpythongo/goal/pkg/common/config"
	"github.com/cnpythongo/goal/pkg/response"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var jwtSecret = []byte(config.GetConfig().App.Secret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(15 * 24 * time.Hour) // 15天
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "goal-console",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = response.SuccessCode
		token := c.GetHeader("token")
		if token == "" {
			response.FailJsonResp(c, response.JWTTokenEmptyError, nil)
			c.Abort()
			return
		}

		claims, err := ParseToken(token)
		if err != nil || claims == nil {
			code = response.JWTTokenParseError
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = response.JWTTokenExpiredError
		}
		if code != 0 {
			response.FailJsonResp(c, code, nil)
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Next()
	}
}

func SetupJwtTokenToHeader(c *gin.Context, token string) {
	c.Header("token", token)
}

func CleanupJwtToken(c *gin.Context) {
	c.Header("token", "")
}
