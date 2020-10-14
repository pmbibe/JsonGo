package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var storyF StoryF

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}
type StoryF map[string]Chapter

func readFile(file string) string {
	dat, _ := ioutil.ReadFile(file)
	return string(dat)

}
func rePrintJSON(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	out.WriteTo(os.Stdout)
}

//use Decode
func useDecode(file string) {
	f, _ := os.Open(file)
	d := json.NewDecoder(f)
	d.Decode(&storyF)
	rePrintJSON(storyF)
}

//use UnMarshal
func useMarshal(file string) {
	data := readFile(file)
	err := json.Unmarshal([]byte(data), &storyF)
	if err != nil {
		log.Fatal(err)
	}
	rePrintJSON(storyF)
}
func main() {
	file := "gopher.json"
	useMarshal(file)
}
