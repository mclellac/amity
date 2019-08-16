package server

import (
	"fmt"
	"log"
	"os"

	"github.com/mclellac/amity/lib/api"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Config struct {
	Server
	Database
}

type Server struct {
	DomainName string
}

type Database struct {
	Username     string
	Password     string
	Hostname     string
	DatabaseName string
}

type Daemon struct {
}

func (d *Daemon) getDB(cfg Config) (*gorm.DB, error) {

	connectionString := "postgres://" + cfg.Database.Username + ":" +
		cfg.Database.Password + "@" +
		cfg.Database.Hostname + ":5432/" +
		cfg.Database.DatabaseName + "?sslmode=disable"

	fmt.Println(connectionString)
	return gorm.Open("postgres", connectionString)
}

func (d *Daemon) Migrate(cfg Config) error {
	db, err := d.getDB(cfg)
	if err != nil {
		return err
	}
	// Disable table name's pluralization
	db.SingularTable(true)
	// Enable Logger
	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))

	db.AutoMigrate(api.Post{}, api.OAuth2{}, api.User{})
	return nil
}

func (d *Daemon) Run(cfg Config) error {
	// TODO: Make the logfile location/name a configurable item in the server config. For now we'll just hardcode it.
	logfile, err := os.OpenFile("amityd.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer logfile.Close()

	db, err := d.getDB(cfg)
	if err != nil {
		return err
	}
	defer db.Close()

	// Disable table name's pluralization
	db.SingularTable(true)

	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = logfile

	handler := &Service{db: db}

	r := gin.New()

	// Global middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Resources
	r.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "ハローワールド"}) })
	r.POST("/post/new", handler.CreatePost)
	r.GET("/posts", handler.GetAllPosts)
	r.GET("/post/:id", handler.GetPost)
	r.DELETE("/post/:id/", handler.DeletePost)
	r.PUT("/post/:id", handler.UpdatePost)

	// Run
	r.Run(cfg.Server.DomainName)

	return nil
}
