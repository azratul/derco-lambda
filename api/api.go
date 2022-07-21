package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"

	"derco-backend/config"
	"derco-backend/jsonplaceholder"

	"log"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	log.Println("***GetPosts***")
	url := fmt.Sprintf("%s/posts", config.Conf.JsonPlaceHolder.Url)
	var posts jsonplaceholder.Posts
	posts.Stage = os.Getenv("GIN_MODE")

	if err := request(url, &posts.Post); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		log.Fatal(err)
		return
	}

	wg := sync.WaitGroup{}

	// TODO: Optimize this. Too many request. Save the data in a slice
	for i, post := range posts.Post {
		if os.Getenv("GIN_MODE") != "release" {
			log.Println("Post ID: ", post.ID)
		}
		wg.Add(1)
		go func(i int, post jsonplaceholder.Post) {
			posts.Post[i].User = jsonplaceholder.User{}
			if err := request(fmt.Sprintf("%s/users/%d", config.Conf.JsonPlaceHolder.Url, post.UserID), &posts.Post[i].User); err != nil {
				c.JSON(http.StatusInternalServerError, err)
				log.Fatal(err)
				return
			}
			wg.Done()
		}(i, post)
	}

	wg.Wait()

	c.JSON(http.StatusOK, posts)
}

func request(url string, obj interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(&obj)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	req.Close = true

	return nil
}
