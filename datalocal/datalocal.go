package datalocal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var storyF StoryF

//Option Oki
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

//Chapter Oki
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

//StoryF Oki
type StoryF map[string]Chapter

//ReadFile Oki
func ReadFile(file string) string {
	dat, _ := ioutil.ReadFile(file)
	return string(dat)

}

//RePrintJSON Oki
func RePrintJSON(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	out.WriteTo(os.Stdout)

}

//UseDecode Oki
func UseDecode(file string) StoryF {
	f, _ := os.Open(file)
	d := json.NewDecoder(f)
	d.Decode(&storyF)
	RePrintJSON(storyF)
	return storyF
}

//UseMarshal Oki
func UseMarshal(file string) StoryF {
	data := ReadFile(file)
	err := json.Unmarshal([]byte(data), &storyF)
	if err != nil {
		log.Fatal(err)
	}
	RePrintJSON(storyF)
	return storyF
}

//Data Oki
func Data(file string) {
	UseMarshal(file)
}
