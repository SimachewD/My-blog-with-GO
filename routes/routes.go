package routes

import (
	"database/sql"
	"myblog/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

)

func SetupRoutes(r *gin.Engine, db *sql.DB) {

	r.GET("/posts", handlers.GetPosts(db))
	r.POST("/posts", handlers.CreatePost(db))
	r.PUT("/posts/:id", handlers.UpdatePost(db))
	r.DELETE("/posts/:id", handlers.DeletePost(db))
}
