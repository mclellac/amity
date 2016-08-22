package client

import (
	"strconv"

	"github.com/mclellac/amity/lib/api"
)

type Client struct {
	Host string
}

// CreatePost submits a new post entry to the datastore.
func (client *Client) CreatePost(title string, article string) (api.Post, error) {
	var resp api.Post

	post := api.Post{Title: title, Article: article}
	url := client.Host + "/post/new"
	r, err := makeRequest("POST", url, post)
	if err != nil {
		return resp, err
	}

	err = processResponseEntity(r, &resp, 201)
	return resp, err
}

// GetAllPosts retrieves all posts from the datastore.
func (client *Client) GetAllPosts() ([]api.Post, error) {
	var resp []api.Post

	url := client.Host + "/posts"
	r, err := makeRequest("GET", url, nil)
	if err != nil {
		return resp, err
	}

	err = processResponseEntity(r, &resp, 200)
	return resp, err
}

// GetPost takes an integer ID as an argument, and retrieves a post
// from the datastore that corresponds with that ID.
func (client *Client) GetPost(id int32) (api.Post, error) {
	var resp api.Post

	url := client.Host + "/post/" + strconv.FormatInt(int64(id), 10)
	r, err := makeRequest("GET", url, nil)
	if err != nil {
		return resp, err
	}
	err = processResponseEntity(r, &resp, 200)

	return resp, err
}

// DeletePost takes an ID and removes the corresponding post from the datastore
func (client *Client) DeletePost(id int32) error {
	url := client.Host + "/post/" + strconv.FormatInt(int64(id), 10) + "/"
	r, err := makeRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return processResponse(r, 204)
}

func (client *Client) UpdatePost(post api.Post) (api.Post, error) {
	var resp api.Post

	url := client.Host + "/post/" + strconv.FormatInt(int64(post.Id), 10)
	r, err := makeRequest("PUT", url, post)
	if err != nil {
		return resp, err
	}
	err = processResponseEntity(r, &resp, 200)

	return resp, err
}
