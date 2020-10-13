package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}
type StoryF struct {
	Intros    Chapter `json:"intro"`
	NewYork   Chapter `json:"new-york"`
	Debate    Chapter `json:"debate"`
	SeanKelly Chapter `json:"sean-kelly"`
	MarkBates Chapter `json:"mark-bates"`
	Denver    Chapter `json:"denver"`
	Home      Chapter `json:"home"`
}

func readFile(file string) string {
	dat, _ := ioutil.ReadFile(file)
	return string(dat)

}
func main() {
	var storyF StoryF
	data := readFile("gopher.json")
	err := json.Unmarshal([]byte(data), &storyF)
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(storyF)
	if err != nil {
		log.Fatal(err)
	}
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	out.WriteTo(os.Stdout)

}
