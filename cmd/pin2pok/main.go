package main

import (
	"encoding/json"
	"fmt"
	"github.com/johnharris85/pokcli/pkg/client"
	"io/ioutil"
	"os"
	"strings"
	"time"

	_ "github.com/johnharris85/pokcli/pkg/client"
	log "github.com/sirupsen/logrus"
)

type pinboardBookmarkList []pinboardBookmark

type pinboardBookmark struct {
	URL         string    `json:"href"`
	Description string    `json:"description"`
	Extended    string    `json:"extended"`
	Meta        string    `json:"meta"`
	Hash        string    `json:"hash"`
	Time        time.Time `json:"time"`
	Shared      string    `json:"shared"`
	ToRead      string    `json:"toread"`
	Tags        string    `json:"tags"`
}

type pocketBookmarkList []pocketBookmark

type pocketBookmark struct {
	ItemID int       `json:"item_id,omitempty"`
	RefID  int       `json:"ref_id,omitempty"`
	Tags   string    `json:"tags,omitempty"`
	Time   time.Time `json:"time,omitempty"`
	Title  string    `json:"title,omitempty"`
	URL    string    `json:"url,omitempty"`
}

func main() {
	// l := log.New()
	pinboardFile, err := os.Open("/home/john/Downloads/pinboard.json")
	if err != nil {
		log.Fatal(err)
	}
	defer pinboardFile.Close()

	content, err := ioutil.ReadAll(pinboardFile)
	if err != nil {
		log.Fatal(err)
	}

	list := new(pinboardBookmarkList)

	json.Unmarshal([]byte(content), &list)

	c, err := client.NewClientWithOpts(client.WithCredsFile("/home/john/.pokcli/credentials"))
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(c.GetArticles("oauth")))
	/*
		pocketList := &pocketBookmarkList{}

		for _, b := range *list {
			*pocketList = append(*pocketList, *pinboardToPocket(b))
		}

		pl, _ := json.Marshal(pocketList)
		fmt.Println(string(pl))
	*/
}

func pinboardToPocket(pinBookmark pinboardBookmark) *pocketBookmark {
	pocketTags := strings.Replace(pinBookmark.Tags, " ", ",", -1)
	if pocketTags == "" {
		pocketTags = "no-tag"
	}
	return &pocketBookmark{
		Tags:  pocketTags,
		Time:  pinBookmark.Time,
		Title: pinBookmark.Description,
		URL:   pinBookmark.URL,
	}
}
