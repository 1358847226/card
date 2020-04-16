package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type UserClaims struct {
	UserID    string `json:"userId"`
	jwt.StandardClaims
}

func NewToken(uid string) (string , error){
	expTime := time.Now().Add(time.Hour * 24).Unix()
	claims := &UserClaims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := "kasdnjklasn"
	return token.SignedString(([]byte(key)))
}



func Validate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.Request.Header.Get("Access-Token")

		log.Println("秘钥", key)
		callFun := func(token *jwt.Token) (interface{}, error) {
			return []byte("kasdnjklasn"), nil
		}

		log.Println("验证密钥", key)
		if token, err := jwt.ParseWithClaims(key, &UserClaims{}, callFun); err == nil {
			if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
				log.Printf("数据 %v", claims)
				ctx.Set("userid", claims.UserID)
				ctx.Next()
			} else {
				ctx.JSON(401, gin.H{
					"code":401,
					"message": "登录过期",
				})
				ctx.Abort()
			}
		} else {
			ctx.JSON(401, gin.H{
				"code": 401,
				"message": "登录过期",
			})
			ctx.Abort()
		}


	}
}