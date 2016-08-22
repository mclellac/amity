package server

import (
	"log"
	"os"

	"github.com/mclellac/amity/lib/api"

	"github.com/gin-gonic/gin"
	//"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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
	connectionString := cfg.Database.Username + ":" +
		cfg.Database.Password + "@tcp(" +
		cfg.Database.Hostname + ":3306)/" +
		cfg.Database.DatabaseName + "?charset=utf8&parseTime=True"

	return gorm.Open("mysql", connectionString)
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
		log.Fatalf("error opening file: %v", err)
	}
	defer logfile.Close()

	db, err := d.getDB(cfg)
	if err != nil {
		return err
	}
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
	r.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Hello"}) })
	r.POST("/post/new", handler.CreatePost)
	r.GET("/posts", handler.GetAllPosts)
	r.GET("/post/:id", handler.GetPost)
	r.DELETE("/post/:id", handler.DeletePost)
	r.PUT("/post/:id", handler.UpdatePost)

	// Run
	r.Run(cfg.Server.DomainName)

	return nil
}
