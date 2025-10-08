package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const (
	tgEndpointAPI = "https://api.telegram.org/bot"
)

type Client struct {
	token string
}

type Update struct {
	UpdateID int `json:"update_id"`
	Message  *struct {
		MessageID int    `json:"message_id"`
		From      *User  `json:"from"`
		Chat      Chat   `json:"chat"`
		Text      string `json:"text"`
	} `json:"message"`
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
}

type Chat struct {
	ID int64 `json:"ID"`
}

type sendMessageRequest struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func NewClient(token string) *Client {
	return &Client{token: token}
}

func (c *Client) GetUpdates(offset int) ([]Update, error) {
	url := fmt.Sprintf("%s%s/getUpdates", tgEndpointAPI, c.token)

	reqURL := url + "?offset=" + strconv.Itoa(offset)
	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, fmt.Errorf("http get: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read reasponse body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("telegram API returned status: %d, desc: %s", resp.StatusCode, string(body))
	}

	var result struct {
		OK     bool     `json:"ok"`
		Result []Update `json:"result"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response :%w", err)
	}

	return result.Result, nil
}

func (c *Client) SendMessage(chatID, text string) error {
	url := fmt.Sprintf("%s%s/sendMessage", tgEndpointAPI, c.token)

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
		return fmt.Errorf("telegram API returned status: %d", resp.StatusCode)
	}

	return nil
}
