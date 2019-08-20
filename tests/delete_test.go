package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/mclellac/amity/lib/client"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

func TestDeletePost(t *testing.T) {
	// given
	client := client.Client{Host: "http://localhost:3001"}
	post, _ := client.CreatePost("foo", "bar")
	ID := post.ID

	// when
	err := client.DeletePost(ID)

	// then
	if err != nil {
		t.Error(err)
	}

	_, err = client.GetPost(ID)
	if err == nil {
		t.Error(err)
	}
}

func TestDeleteNotFound(t *testing.T) {
	// given
	client := client.Client{Host: "http://localhost:3001"}
	ID := int32(3)
	// when
	err := client.DeletePost(ID)

	// then
	if err == nil {
		t.Error(err)
	}

}
