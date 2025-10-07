package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	token string
}

type sendMessageRequest struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func NewClient(token string) *Client {
	return &Client{token: token}
}

func (c *Client) SendMessage(chatID, text string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", c.token)

	payload := sendMessageRequest{
		ChatID: chatID,
		Text:   text,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal payload: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("http post: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram API returned status %d", resp.StatusCode)
	}

	return nil
}
