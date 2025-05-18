package twitter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Client struct {
	URL string
}

func NewTwitterClient(apiURL string) *Client {
	return &Client{URL: apiURL}
}

func TweetAlert(id int, price float64, diff float64, up bool, client *Client) {
	direction := "📉 DECREASE"
	if up {
		direction = "📈 INCREASE"
	}

	text := fmt.Sprintf("%s > %.2f%% SINCE THE LAST CALL \nNew price : %.4f $TAO\n#TAO #SN%d", direction, diff, price, id)

	payload := map[string]string{"text": text}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Erreur encodage JSON : %v", err)
		return
	}

	resp, err := http.Post(client.URL+"/tweet", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Erreur envoi tweet : %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Tweet refusé par le service Flask : %s", resp.Status)
	} else {
		log.Printf("Tweet envoyé : %s", text)
	}
}