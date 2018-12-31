package main 

import ( 
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/xml"
	"gopkg.in/yaml.v2"
) 

func main() {
	
	file, err := ioutil.ReadFile("feeds.yml")
	
	if err != nil {
		fmt.Print("error reading file:")
	}

	var feeds YamlFeeds
	err = yaml.Unmarshal(file, &feeds)

	if err != nil {
		fmt.Print("error unmarshalling yaml")
	}

	fmt.Print("\n\n\n")

	words := map[string]*[]Item{}

	for i := 0; i < len(feeds.Feeds.News); i++ {
		resp, err := http.Get(feeds.Feeds.News[i])

		if err != nil {
			fmt.Print("error getting news\n")
		}

		defer resp.Body.Close()
	
		body, err := ioutil.ReadAll(resp.Body)

		var news Rss

		xml.Unmarshal(body, &news)

		fmt.Print("\n----------------\n")
		fmt.Print(news.Channel.Title + "\n")
		fmt.Print(news.Channel.Description + "\n")
		fmt.Print(feeds.Feeds.News[i] + "\n")
		fmt.Print("items: " + string(len(news.Channel.Items)) + "...\n")
		for j := 0; j< len(news.Channel.Items); j++ {
			fmt.Printf("%s\n", news.Channel.Items[j].Title)
		//	itemArr = [1]Item{ news.Channel.Items[j] }
		//	words[news.Channel.Items[j].Title] = &itemArr
		}
	}
	for key, value := range words {
		fmt.Println("Key:", key, "Value:", value)
	}

}

type YamlFeeds struct {
	Feeds struct {
		Reddit []string `yml:"reddit"`
		News []string `yml:"news"`
	}
}

type Rss struct {
	XmlName xml.Name `xml:"rss"`
	Channel Channel `xml:"channel"`
}

type Channel struct {
	XmlName xml.Name `xml:"channel"`
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`
	Image Image `xml:"image"`
	Items []Item `xml:"item"`
}

type Image struct {
	XmlName xml.Name `xml:"image"`
	Url string `xml:"url"`
}

type Item struct {
	XmlName xml.Name `xml:"item"`
	Title string `xml:"title"`
 	Link string `xml:"link"`
	Description string `xml:"description"`
}
