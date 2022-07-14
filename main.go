package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	router = gin.Default()
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var exampleUser = User{
	ID:       1,
	Username: "admin",
	Password: "admin",
}

func video(c *gin.Context) {
	videoid := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id":            videoid,
		"episode_count": 12,
		"title":         "Kancolle!",
		"video":         "/video/バカとテストと召喚獣 - S01E01.mkv",
		"subtitle":      "/video/バカとテストと召喚獣 - S01E01.zh-CN.vtt",
	})
}

func videoinfo(c *gin.Context) {
	// videoid := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"title":       "Kancolle!",
		"description": "Kancolle description",
	})
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			// 可将将* 替换为指定的域名
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func main() {
	router.Use(Cors())
	router.GET("/api/v1/video/:id", video)
	router.GET("/api/v1/videoinfo/:id", videoinfo)
	log.Fatal(router.Run(":8080"))
}
