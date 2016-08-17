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
	app.Usage = "The command like application that allows you to interact with amityd."
	app.Version = "0.0.2"

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "host", Value: "http://localhost:3000", Usage: "amityd server host"},
	}

	app.Commands = []cli.Command{
		{
			Name:        "add",
			Usage:       "Create a new post.",
			Description: "Adds new article to the database.\n\nEXAMPLE:\n   $ amity add \"Test Title\" \"Test article body...\"",
			ArgsUsage:   "[\"post title\"] [\"post body\"]",
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
			Name:        "ls",
			Usage:       "List all posts.",
			Description: "Displays the IDs and titles of posts on the server.\n\nEXAMPLE:\n   $ amity ls",
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
			Name:        "read",
			Usage:       "Display the article of the supplied ID.",
			Description: "Retrieves the article of the post, and displays it.\n\nEXAMPLE:\n   $ amity read 2",
			ArgsUsage:   "[ID]",
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
			Name:        "rm",
			Usage:       "Delete a post.",
			Description: "Remove the post with the supplied ID from the server.\n\nEXAMPLE:\n   $ amity rm 2",
			ArgsUsage:   "[ID]",
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
