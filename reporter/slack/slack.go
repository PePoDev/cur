// Package slack use for send post request to slack webhook
package slack

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Reporter struct use as builder design pattern
type Reporter struct {
	title    string
	imageURL string
	content  string
}

// SetTitle will set report title with slack format
func (report *Reporter) SetTitle(title string) {
	report.title = title
}

// SetImageURL will set image url as attachment with slack format
func (report *Reporter) SetImageURL(imageURL string) {
	report.imageURL = imageURL
}

// SetContent will set content as text with slack format
func (report *Reporter) SetContent(content string) {
	report.content = content
}

// Send will create post request to slack webhook
func (report Reporter) Send() {
	url := os.Getenv("SLACK_WEBHOOK_URL")
	content := fmt.Sprintf("{\"text\":\"%v\n%v\"", fmt.Sprintf("*%v*", report.title), report.content)
	if report.imageURL != "" {
		content += fmt.Sprintf(",\"attachments\":[{\"image_url\": \"%v\"}]", report.imageURL)
	}
	content += "}"
	log.Println(content)
	var jsonStr = []byte(content)
	log.Printf("Send Message: %s -> URL: %s\n", jsonStr, url)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
}
