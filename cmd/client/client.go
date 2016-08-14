package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mclellac/amity/lib/client"
	"github.com/mclellac/amity/lib/ui"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "amity"
	app.Usage = "Create & list posts on an Amity server."
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "host", Value: "http://localhost:3000", Usage: "amityd server host"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "(title body) - create a new post",
			Action: func(c *cli.Context) error {
				title := c.Args().Get(0)
				article := c.Args().Get(1)
				host := c.GlobalString("host")
				client := client.Client{Host: host}

				post, err := client.CreatePost(title, article)
				if err != nil {
					log.Fatal(err)
				}
				ui.DrawDivider()
				fmt.Printf("%s%+v%s\n", ui.Grey, post, ui.Reset)

				return nil
			},
		},
		{
			Name:  "ls",
			Usage: "list all posts",
			Action: func(c *cli.Context) error {
				host := c.GlobalString("host")
				client := client.Client{Host: host}

				posts, err := client.GetAllPosts()
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("%s[%sid:%s]\t%sTitle:%s\n",
					ui.LightBlue,
					ui.Grey,
					ui.LightBlue,
					ui.Grey,
					ui.Reset)

				ui.DrawDivider()

				for _, post := range posts {
					client.ListView(post.Id, post.Title)
				}

				return nil
			},
		},
		{
			Name:  "read",
			Usage: "(id) delete a post",
			Action: func(c *cli.Context) error {
				idStr := c.Args().Get(0)

				id, err := strconv.Atoi(idStr)
				if err != nil {
					log.Print(err)
					return nil
				}

				host := c.GlobalString("host")
				client := client.Client{Host: host}

				post, err := client.GetPost(int32(id))
				if err != nil {
					log.Fatal(err)
					return nil
				}

				client.ReadView(post.Id, post.Title, post.Article)

				return nil
			},
		},
		{
			Name:  "rm",
			Usage: "(id) delete a post",
			Action: func(c *cli.Context) error {
				idStr := c.Args().Get(0)

				id, err := strconv.Atoi(idStr)
				if err != nil {
					log.Print(err)
					return nil
				}

				host := c.GlobalString("host")
				client := client.Client{Host: host}

				post, err := client.GetPost(int32(id))
				if err != nil {
					log.Fatal(err)
					return nil
				}

				client.DeletePost(post.Id)
				fmt.Printf("%s%+v%s\n", ui.Grey, post, ui.Reset)

				return nil
			},
		},
	}

	app.Run(os.Args)
}
