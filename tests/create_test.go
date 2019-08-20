package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/mclellac/amity/lib/client"
)

var _ = fmt.Print
var _ = log.Print

func TestCreatePost(t *testing.T) {
	// given
	client := client.Client{Host: "http://localhost:3001"}

	// when
	post, err := client.CreatePost("foo", "bar")

	// then
	if err != nil {
		t.Error(err)
	}

	if post.Title != "foo" && post.Article != "bar" {
		t.Error("Post title and article doesn't match")
	}

	// cleanup
	_ = client.DeletePost(post.ID)
}
