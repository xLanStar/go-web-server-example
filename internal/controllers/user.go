package controllers

import (
	"go-web-server-example/internal/exception"
	"go-web-server-example/internal/services/auth"
	"go-web-server-example/internal/services/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var data struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}

	c.BindJSON(&data)

	// TODO 檢查是否已存在相同的帳號
	originalUser := db.GetUserByAccount(data.Account)

	if originalUser != nil {
		c.Status(http.StatusNotAcceptable)
		return
	}

	// 建立資料
	user := db.CreateUser(data.Account, data.Password)

	// 產生 JWT Token
	token, err := auth.GenerateToken(data.Account, data.Password)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user, // TODO response user data
	})
}

func Login(c *gin.Context) {
	var data struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}

	c.BindJSON(&data)

	user := db.GetUserByAccount(data.Account)

	if user == nil || data.Password != user.Password {
		c.Status(http.StatusUnauthorized)
		return
	}

	// 產生 Token
	token, err := auth.GenerateToken(data.Account, data.Password)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": string(token),
		"user":  user,
	})
}

func Validate(c *gin.Context) {
	user, ok := c.Get("user")
	if ok && user != nil {
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": exception.UNKNOWN,
		})
	}
}
