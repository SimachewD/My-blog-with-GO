package handlers

import (

	"database/sql"
	"myblog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPosts handles GET requests to retrieve all posts
func GetPosts(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        rows, err := db.Query("SELECT id, title, content FROM posts")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer rows.Close()

        var posts []models.Post
        for rows.Next() {
            var post models.Post
            if err := rows.Scan(&post.ID, &post.Title, &post.Content); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            posts = append(posts, post)
        }

        c.JSON(http.StatusOK, posts)
    }
}


// CreatePost handles POST requests to create a new post
func CreatePost(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var newPost models.Post
        if err := c.ShouldBindJSON(&newPost); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        _, err := db.Exec("INSERT INTO posts (title, content) VALUES ($1, $2)", newPost.Title, newPost.Content)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, newPost)
    }
}

// UpdatePost handles PUT requests to update an existing post
func UpdatePost(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
            return
        }

        var updatedPost models.Post
        if err := c.ShouldBindJSON(&updatedPost); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        _, err = db.Exec("UPDATE posts SET title=$1, content=$2 WHERE id=$3", updatedPost.Title, updatedPost.Content, id)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, updatedPost)
    }
}

// DeletePost handles DELETE requests to delete a post
func DeletePost(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
            return
        }

        _, err = db.Exec("DELETE FROM posts WHERE id=$1", id)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
    }
}