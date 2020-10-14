package reporter

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// SlackReporter is struct for create Slakc Report Information and send it by WebHook
type SlackReporter struct {
	title    string
	imageURL string
	content  string
}

// SetTitle will set report title with slack format
func (report *SlackReporter) SetTitle(title string) {
	report.title = fmt.Sprintf("*%v*", title)
}

// SetImageURL will set image url as attachment with slack format
func (report *SlackReporter) SetImageURL(imageURL string) {
	report.imageURL = imageURL
}

// SetContent will set content as text with slack format
func (report *SlackReporter) SetContent(content string) {
	report.content = strings.Replace(content, "\"", "\\\"", -1)
}

// Send will create post request to slack webhook
func (report SlackReporter) Send(url string) {
	content := fmt.Sprintf("{\"text\":\"%v\n%v\"", report.title, report.content)
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
