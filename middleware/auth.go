package middleware

import (
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go-cloud-disk/auth"
	"go-cloud-disk/model"
	"go-cloud-disk/serializer"
	"go-cloud-disk/utils"
)

// JWTAuth check jwt auth and save jwt info
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token format Authorization: "Bearer [token]"
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.JSON(200, serializer.NotLogin("Need Token"))
			c.Abort()
			return
		}

		parts := strings.Split(authorization, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(200, serializer.NotLogin("Token format error"))
			c.Abort()
			return
		}

		// parse token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(200, serializer.NotLogin("Token error"))
			c.Abort()
			return
		}

		// check if the token has expired
		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			c.JSON(200, serializer.NotLogin("Token expiration"))
			c.Abort()
			return
		}

		c.Set("UserId", claims.UserId)
		c.Set("UserName", claims.UserName)
		c.Set("Status", claims.Status)

		c.Next()
	}
}

func CasbinAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get obj and act
		userStatus := c.MustGet("Status").(string)
		method := c.Request.Method
		path := c.Request.URL.Path
		object := strings.TrimPrefix(path, "/api/v1/")

		if ok, _ := auth.Casbin.Enforce(userStatus, object, method); !ok {
			c.JSON(200, serializer.NotAuthErr("not auth"))
			c.Abort()
		}
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.MustGet("UserId").(string)
		var user model.User
		if err := model.DB.Where("uuid = ?", uuid).Find(&user).Error; err != nil {
			log.Println("get user info err when check admin auth", err)
			c.Abort()
			return
		}
		if user.Status != c.MustGet("Status").(string) {
			c.JSON(200, serializer.NotAuthErr("change jwt!!!"))
			c.Abort()
			return
		}
	}
}
