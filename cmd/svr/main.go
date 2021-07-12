package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"sort"
)

type Posts []Post

type Post struct {
	UserId int    `json:"userId,omitempty"`
	Id     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
}

const URL = "https://jsonplaceholder.typicode.com/posts"

func main() {
	//Fetch JSON data
	res, err := http.Get(URL)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	//Convert ioReadCloser to []Byte
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	//Unmarshal the response into data structure
	posts := Posts{}
	err = json.Unmarshal(body, &posts)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	sortById(posts)
	jsonPrint(posts)
}

//sortById sorts the slice by Id value in greatest to least
func sortById(posts Posts) Posts {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Id > posts[j].Id
	})

	return posts
}

//jsonPrint marshals a go data struct to JSON to be pretty printed in the output
func jsonPrint(posts Posts) {
	postJSON, err := json.MarshalIndent(posts, "", " ")
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	fmt.Printf("JSON data: \n %s\n", string(postJSON))
}