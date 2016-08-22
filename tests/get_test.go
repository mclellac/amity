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
	client := client.Client{Host: "http://localhost:3000"}
	post, _ := client.CreatePost("foo", "bar")
	id := post.Id

	// when
	post, err := client.GetPost(id)

	// then
	if err != nil {
		t.Error(err)
	}

	if post.Title != "foo" && post.Article != "bar" {
		t.Error("returned post not right")
	}

	// cleanup
	_ = client.DeletePost(post.Id)
}

func TestGetNotFound(t *testing.T) {
	// given
	client := client.Client{Host: "http://localhost:3000"}
	id := int32(3)

	// when
	_, err := client.GetPost(id)

	// then
	if err == nil {
		t.Error(err)
	}
}

func TestGetAllPosts(t *testing.T) {
	// given
	client := client.Client{Host: "http://localhost:3000"}
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
	_ = client.DeletePost(posts[0].Id)
	_ = client.DeletePost(posts[1].Id)
}
