package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
	"unicode"
)

type data struct {
	ResponseData struct {
		TranslatedText string `json:"TranslatedText"`
	} `json:"responseData"`
}

func Translate(text string, flag bool) string {
	langpair := "en|zh-cn"
	if flag == false {
		langpair = "zh-cn|en"
	}
	url := fmt.Sprintf("%s%s%s%s", "http://mymemory.translated.net/api/get?q=", url.QueryEscape(text), "&langpair=", langpair)

	httpClient := http.Client{
		Timeout: time.Second * 5, // Maximum of 5 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/535.19(KHTML, like Gecko) Chrome/18.0.1025.168 Safari/535.19")

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonData := data{}
	jsonErr := json.Unmarshal(body, &jsonData)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	result := jsonData.ResponseData.TranslatedText
	return result
}

func IsAsciiPrintable(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII || !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gt <word>")
		os.Exit(2)
	}

	user_input := os.Args[1]
	is_english := IsAsciiPrintable(user_input)
	fmt.Println(Translate(user_input, is_english))
}
