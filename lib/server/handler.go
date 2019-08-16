package server

import (
	"log"
	"strconv"
	"time"

	"github.com/mclellac/amity/lib/api"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Service struct
type Service struct {
	db *gorm.DB
}

// CreatePost takes the request and adds puts it in the
// database.
func (s *Service) CreatePost(c *gin.Context) {
	var post api.Post

	if err := c.Bind(&post); err != nil {
		c.JSON(400, api.NewError("HTTP: 400 Bad Request"))
		return
	}

	post.Created = int32(time.Now().Unix())

	s.db.Save(&post)

	c.JSON(201, post)
}

// DeletePost finds a post entry in the database based on the ID
// and removes it from the table.
func (s *Service) DeletePost(c *gin.Context) {
	id, err := s.GetID(c)
	if err != nil {
		c.JSON(400, api.NewError("Unable to delete post."))
		return
	}

	var post api.Post

	if s.db.First(&post, id).RecordNotFound() {
		c.JSON(404, api.NewError("Unable to find the post."))
	} else {
		s.db.Delete(&post)
		c.Data(204, "application/json", make([]byte, 0))
	}
}

// GetAllPosts retrieves all posts in the database
func (s *Service) GetAllPosts(c *gin.Context) {
	var posts []api.Post

	s.db.Order("created desc").Find(&posts)

	c.JSON(200, posts)
}

// GetPost takes the ID of the post entry from the
// context, finds it in the database, and returns the
// entry formatted as JSON.
func (s *Service) GetPost(c *gin.Context) {
	id, err := s.GetID(c)
	if err != nil {
		c.JSON(400, api.NewError("Unable to locate post."))
		return
	}

	var post api.Post

	if s.db.First(&post, id).RecordNotFound() {
		c.JSON(404, gin.H{"error": "HTTP 404 not found."})
	} else {
		c.JSON(200, post)
	}
}

// GetID retrieves the entry from the database by ID.
func (s *Service) GetID(c *gin.Context) (int32, error) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	return int32(id), nil
}

// UpdatePost takes the ID of the post entry from the
// context, and updates the database entry.
func (s *Service) UpdatePost(c *gin.Context) {
	id, err := s.GetID(c)
	if err != nil {
		c.JSON(400, api.NewError("Problem decoding ID sent"))
		return
	}

	var post api.Post

	if err := c.Bind(&post); err != nil {
		c.JSON(400, api.NewError("Problem decoding article"))
		return
	}

	post.ID = int32(id)

	var existing api.Post

	if s.db.First(&existing, id).RecordNotFound() {
		c.JSON(404, api.NewError("HTTP 404 not found."))
	} else {
		s.db.Save(&post)
		c.JSON(200, post)
	}
}
