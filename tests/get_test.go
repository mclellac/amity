package main

import (
	"github.com/mclellac/amity/lib/client"

	"fmt"
	"log"
	"testing"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

func TestGetPost(t *testing.T) {
	// given
	client := client.Client{Host: "http://localhost:3001"}
	post, _ := client.CreatePost("foo", "bar")
	ID := post.ID

	// when
	post, err := client.GetPost(ID)

	// then
	if err != nil {
		t.Error(err)
	}

	if post.Title != "foo" && post.Article != "bar" {
		t.Error("returned post not right")
	}

	// cleanup
	_ = client.DeletePost(post.ID)
}

func TestGetNotFound(t *testing.T) {
	// given
	client := client.Client{Host: "http://localhost:3001"}
	ID := int32(3)

	// when
	_, err := client.GetPost(ID)

	// then
	if err == nil {
		t.Error(err)
	}
}

func TestGetAllPosts(t *testing.T) {
	// given
	client := client.Client{Host: "http://localhost:3001"}
	client.CreatePost("foo", "bar")
	client.CreatePost("baz", "bing")

	// when
	posts, err := client.GetAllPosts()

	// then
	if err != nil {
		t.Error(err)
	}

	if len(posts) != 2 {
		t.Error("Wrong number of posts")
	}

	if posts[0].Title != "foo" && posts[0].Article != "bar" {
		t.Error("Returned post not right")
	}

	if posts[1].Title != "baz" && posts[1].Article != "bing" {
		t.Error("Returned post not right")
	}

	// cleanup
	_ = client.DeletePost(posts[0].ID)
	_ = client.DeletePost(posts[1].ID)
}
