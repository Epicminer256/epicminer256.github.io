package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	UUID       string    `json:"UUID"`
	TimePosted int       `json:"timeposted"`
	Title      string    `json:"title"`
	Content    []Section `json:"content"`
}

type Section struct {
	Heading string `json:"heading"`
	Content string `json:"content"`
}

func currentTime() int {
	curTime := time.Now()
	return int(curTime.Unix())
}

func genUUID() string {
	id := uuid.New()
	return id.String()
}

func parsePosts() ([]Post, error) {
	var posts []Post
	byteArray, err := os.ReadFile("blog.json")
	if err == nil {
		err = json.Unmarshal(byteArray, &posts)
	}
	return posts, err
}
func savePosts(posts []Post) error {
	byteArray, err := json.MarshalIndent(posts, "", "\t")
	if err == nil {
		err = os.WriteFile("blog.json", byteArray, 0777)
	}

	return err
}

func main() {
	posts, _ := parsePosts()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("\nEnter a command (\"h\" for help)")
		scanner.Scan()

		if scanner.Text() == "" {
			fmt.Println("You left it empty, please enter a command")
			continue
		}

		if scanner.Text()[0] == 'h' {
			fmt.Println("[H]elp")
			fmt.Println("[N]ew post")
			fmt.Println("[O]pen post")
			fmt.Println("[Q]uit")
			continue
		}

		if scanner.Text()[0] == 'n' {
			post := Post{}
			fmt.Println("Enter a title (leave blank to exit):")
			scanner.Scan()

			if scanner.Text() == "" {
				continue
			}

			post.Title = scanner.Text()

			for {
				section := Section{}
				fmt.Println("Enter section name (leave blank to finalyze post):")

				scanner.Scan()

				if scanner.Text() == "" {
					fmt.Println("done")
					break
				}

				section.Heading = scanner.Text()

				fmt.Println("Enter content of section:")
				scanner.Scan()

				section.Content = scanner.Text()

				post.Content = append(post.Content, section)
				post.UUID = genUUID()
				post.TimePosted = currentTime()
			}
			for {
				byteArray, err := json.MarshalIndent(post, "", "\t")
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(string(byteArray) + "\nDoes the following look ok?")
				fmt.Println("[y] or [n]")

				scanner.Scan()

				if scanner.Text()[0] == 'y' {
					posts = append(posts, post)
					savePosts(posts)
					break
				}
				if scanner.Text()[0] == 'n' {
					break
				}
			}
			continue
		}

		if scanner.Text()[0] == 'o' {
			for {
				fmt.Println("Posts:")
				for i := 0; i < len(posts); i++ {
					fmt.Println("[" + strconv.Itoa(i+1) + "]" + posts[i].Title)
				}
				print("\nSelect a post to open (leave empty to exit)\n")
				scanner.Scan()

				if scanner.Text() == "" || scanner.Text()[0] == 'q' {
					break
				}

				postNum, err := strconv.Atoi(scanner.Text())
				postNum = postNum - 1

				if err != nil {
					fmt.Println("Invalid input")
				}

				if postNum >= 0 && postNum < len(posts) {
					fmt.Println("=== " + posts[postNum].Title + " ===")
					for i := 0; i < len(posts[postNum].Content); i++ {
						fmt.Println("\n--- " + posts[postNum].Content[i].Heading + "---")
						fmt.Println(posts[postNum].Content[i].Content)
					}
					fmt.Println("")
				} else {
					fmt.Println("Post not found")
				}
			}
			continue
		}

		if scanner.Text()[0] == 'q' {
			fmt.Println("Exiting...")
			break
		}

	}

}
