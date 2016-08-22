package main

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/mclellac/amity/lib/client"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

func TestUpdatePost(t *testing.T) {

	// given
	client := client.Client{Host: "http://localhost:3000"}
	post, _ := client.CreatePost("foo", "bar")

	// when
	post.Title = "foo"
	post.Article = "bar"
	_, err := client.UpdatePost(post)

	// then
	if err != nil {
		t.Error(err)
	}

	postResult, _ := client.GetPost(post.Id)

	if !reflect.DeepEqual(post, postResult) {
		t.Error("returned post not updated")
	}

	// cleanup
	_ = client.DeletePost(post.Id)
}

func TestUpdateNonExistant(t *testing.T) {

	// given
	client := client.Client{Host: "http://localhost:3000"}
	post, _ := client.CreatePost("foo", "bar")
	_ = client.DeletePost(post.Id)

	// when
	post.Title = "baz"
	post.Article = "bing"
	_, err := client.UpdatePost(post)

	// then
	if err == nil {
		t.Error(err)
	}

}
