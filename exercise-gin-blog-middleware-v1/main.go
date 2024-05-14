package main

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var Posts = []Post{
	{ID: 1, Title: "Judul Postingan Pertama", Content: "Ini adalah postingan pertama di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Title: "Judul Postingan Kedua", Content: "Ini adalah postingan kedua di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{
	{Username: "user1", Password: "pass1"},
	{Username: "user2", Password: "pass2"},
}

func AuthMiddleware(users []User) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		authParts := strings.Split(auth, " ")
		if len(authParts) != 2 || authParts[0] != "Basic" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		payload, err := base64.StdEncoding.DecodeString(authParts[1])
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		credentials := strings.Split(string(payload), ":")
		if len(credentials) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		username := credentials[0]
		password := credentials[1]

		valid := false
		for _, user := range users {
			if user.Username == username && user.Password == password {
				valid = true
				break
			}
		}

		if !valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Continue if authentication successful
		c.Next()
	} // TODO: replace this
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//Set up authentication middleware here // TODO: replace this
	r.Use(AuthMiddleware(users))
	r.GET("/posts", func(c *gin.Context) {
		idParam := c.Query("id")
		if idParam != "" {
			id, err := strconv.Atoi(idParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
				return
			}
			post := findPostByID(id)
			if post == nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Postingan tidak ditemukan"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"post": post})
			return
		}
		c.JSON(http.StatusOK, gin.H{"posts": Posts})
	})

	r.POST("/posts", func(c *gin.Context) {
		var newPost Post
		if err := c.ShouldBindJSON(&newPost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		newPost.ID = len(Posts) + 1
		newPost.CreatedAt = time.Now()
		newPost.UpdatedAt = time.Now()
		Posts = append(Posts, newPost)
		c.JSON(http.StatusCreated, gin.H{"message": "Postingan berhasil ditambahkan", "post": newPost})
	})

	return r
}

func findPostByID(id int) *Post {
	for _, post := range Posts {
		if post.ID == id {
			return &post
		}
	}
	return nil
}

func main() {
	r := SetupRouter()

	r.Run(":8080")
}
