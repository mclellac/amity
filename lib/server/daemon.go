package server

import (
	"fmt"
	"log"
	"os"

	"github.com/mclellac/amity/lib/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// Using gorm with a postgres db
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Config struct
type Config struct {
	Server
	Database
}

// Server struct
type Server struct {
	DomainName string
}

// Database config struct
type Database struct {
	Username     string
	Password     string
	Hostname     string
	DatabaseName string
}

// Daemon struct
type Daemon struct {
}

// getDB establishes the Postgres DB connection string from the config file.
func (d *Daemon) getDB(cfg Config) (*gorm.DB, error) {

	connectionString := "postgres://" + cfg.Database.Username + ":" +
		cfg.Database.Password + "@" +
		cfg.Database.Hostname + ":5432/" +
		cfg.Database.DatabaseName + "?sslmode=disable"

	fmt.Println(connectionString)
	return gorm.Open("postgres", connectionString)
}

// Migrate the database
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

// Run the server
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

	router := gin.Default()

	gin.ForceConsoleColor()

	// Global middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	// App routes
	router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "ハローワールド"}) })

	// v1 route group
	v1 := router.Group("/v1")
	{
		v1.POST("/post/new", handler.CreatePost)
		v1.GET("/posts", handler.GetAllPosts)
		v1.GET("/post/:id", handler.GetPost)
		v1.DELETE("/post/:id/", handler.DeletePost)
		v1.PUT("/post/:id", handler.UpdatePost)
	}
	
	
	// Run
	router.Run(cfg.Server.DomainName)

	return nil
}
