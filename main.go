package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/wsxiaoys/terminal/color"
)

const (
	URL = "http://www.tuling123.com/openapi/api"
	KEY = "531352803455eaf17317cd5b315a347d"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		color.Print("@b>> ")
		words, _ := reader.ReadString('\n')
		text, err := request(strings.TrimSpace(words))
		if err != nil {
			color.Println("@rERROR : ", err)
		}

		color.Println("@r>> ", text)
	}

}

func request(words string) (string, error) {
	requestUrl := fmt.Sprintf("%s?key=%s&info=%s", URL, KEY, words)

	resp, err := http.Get(requestUrl)
	if err != nil {
		return "", err
	}

	json, err := simplejson.NewFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	result, err := json.Map()
	if err != nil {
		return "", err
	}

	textValue := result["text"]
	text := reflect.ValueOf(textValue).Interface().(string)

	return text, nil

}
