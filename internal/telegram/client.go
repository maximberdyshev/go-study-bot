package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	tgEndpointAPI = "https://api.telegram.org/bot"
	httpTimeout   = 10 * time.Second
)

type Client struct {
	token string
	http  *http.Client
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
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
}

type Chat struct {
	ID int64 `json:"id"`
}

type sendMessageRequest struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func NewClient(token string) *Client {
	return &Client{
		token: token,
		http: &http.Client{
			Timeout: httpTimeout,
		},
	}
}

func (c *Client) GetUpdates(ctx context.Context, offset int) ([]Update, error) {
	url := fmt.Sprintf("%s%s/getUpdates", tgEndpointAPI, c.token)

	reqURL := url + "?offset=" + strconv.Itoa(offset)
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	resp, err := c.http.Do(req)
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

func (c *Client) SendMessage(ctx context.Context, chatID, text string) error {
	url := fmt.Sprintf("%s%s/sendMessage", tgEndpointAPI, c.token)

	payload := sendMessageRequest{
		ChatID: chatID,
		Text:   text,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("http do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram API error: %d %s", resp.StatusCode, string(body))
	}
	return nil
}
