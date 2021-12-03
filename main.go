package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/babelcoder-courses/bookstore/config"
	"github.com/babelcoder-courses/bookstore/route"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setup() error {
	if err := config.Load(); err != nil {
		return err
	}

	if err := config.InitDB(); err != nil {
		return err
	}

	return nil
}

func setCORS(r *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	r.Use(cors.New(corsConfig))
}

func setUploadFolders(r *gin.Engine) {
	r.Static("/uploads", "./uploads")
	uploadDirs := [...]string{"author", "users", "books"}
	for _, dir := range uploadDirs {
		os.MkdirAll(filepath.Join("uploads", dir), 0755)
	}
}

func main() {
	if err := setup(); err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()
	setCORS(r)
	setUploadFolders(r)

	route.Serve(r)
	r.Run(":" + config.Env.Port)
}
