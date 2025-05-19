package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Client struct {
	WebhookURL string
}

func NewDiscordClient(webhookURL string) *Client {
	return &Client{WebhookURL: webhookURL}
}

func SendAlert(id int, price float64, diff float64, up bool, client *Client) {
	direction := "📉 DECREASE"
	color := 16711680
	
	if up {
		direction = "📈 INCREASE"
		color = 65280
	}

	description := fmt.Sprintf("%.2f%% SINCE THE LAST CALL\nNew price: %.4f $TAO", diff, price)
	
	payload := map[string]interface{}{
		"embeds": []map[string]interface{}{
			{
				"title":       fmt.Sprintf("%s - SN%d", direction, id),
				"description": description,
				"color":       color,
			},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Erreur encodage JSON : %v", err)
		return
	}

	resp, err := http.Post(client.WebhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Erreur envoi notification Discord : %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		log.Printf("Notification Discord refusée : %s", resp.Status)
	} else {
		log.Printf("Notification Discord envoyée pour SN%d (%.2f%%)", id, diff)
	}
}