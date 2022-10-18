package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	pp "github.com/k0kubun/pp/v3"
)

// Discord color values
const (
	ColorRed   = 0x992D22
	ColorGreen = 0x2ECC71
	ColorGrey  = 0x95A5A6
)

type (
	DiscordRequest struct {
		Content string         `json:"content"`
		Embeds  []DiscordEmbeded `json:"embeds"`
	}

	DiscordEmbeded struct {
		Title       string                `json:"title"`
		Description string                `json:"description"`
		Color       int                   `json:"color"`
		Fields      []DiscordEmbededField `json:"fields"`
	}

	DiscordEmbededField struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}

 	DiscordClient struct {
		WebhookUrl string
	}
)

func ValidateDiscordWebhookUrl(webhookUrl string) {
	re := regexp.MustCompile(`https://discord(?:app)?.com/api/webhooks/[0-9]{18,19}/[a-zA-Z0-9_-]+`)
	if ok := re.Match([]byte(webhookUrl)); !ok {
		log.Printf("The Discord WebHook URL doesn't seem to be valid.")
	}
}

func getDiscordClient(webhookUrl string) *DiscordClient {
	ValidateDiscordWebhookUrl(webhookUrl)
	return &DiscordClient{WebhookUrl: webhookUrl}
}


func (client *DiscordClient) alertsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		client.postHandler(w, r)
	default:
		http.Error(w, "unsupported HTTP method", 400)
	}
}

func (client *DiscordClient) postHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var m HookMessage
	if err := dec.Decode(&m); err != nil {
		log.Printf("error decoding message: %v", err)
		http.Error(w, "invalid request body", 400)
		return
	}
	log.Printf("Recieved alerts: \n")
	pp.Println(m)

	go client.postDiscordMessage(m)
}


func (client *DiscordClient) postDiscordMessage(hookMessage HookMessage) {
	for _, alert := range hookMessage.Alerts {
		request := client.BuildDiscordMessageRequest(hookMessage, alert)
		go client.MakePostRequest(request)
	}
}

func (client *DiscordClient) BuildDiscordMessageRequest(message HookMessage, alert Alert) DiscordRequest {
	embdedFields := []DiscordEmbededField{
		{Name: "Description", Value: alert.Annotations.Description},
		{Name: "GeneratorURL", Value: alert.GeneratorURL},
		{Name: "Labels", Value: getLabelString(alert.Labels)},
	}

	RichEmbed := DiscordEmbeded{
		Title:       fmt.Sprintf("[%s:%d] %s", strings.ToUpper(alert.Status), len(message.Alerts), message.CommonLabels.Alertname),
		Description: message.CommonAnnotations.Summary,
		Color:       ColorGrey,
		Fields:      embdedFields,
	}

	if alert.Status == "firing" {
		RichEmbed.Color = ColorRed
	} else if alert.Status == "resolved" {
		RichEmbed.Color = ColorGreen
	}

	return DiscordRequest{
		Content: message.CommonAnnotations.Summary,
		Embeds: []DiscordEmbeded{RichEmbed},
	}
}

func (client *DiscordClient) MakePostRequest(discordRequest DiscordRequest) {
	requestJson, _ := json.Marshal(discordRequest)
	log.Printf("Sending Post Request to Discord Webhook")
	_, err := http.Post(client.WebhookUrl, "application/json", bytes.NewReader(requestJson))

	if err != nil {
		log.Printf("ERROR: Recieved error response %s", err)
	}
}

func getLabelString(labels map[string]string) string {
	labelString := ""
	for key, value := range labels {
		labelString += fmt.Sprintf("\n%s %s", key, value)
	}
	return labelString
}
