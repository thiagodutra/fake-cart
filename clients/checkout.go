package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Checkout struct {
	ID         string  `json:"id"`
	CustomerID string  `json:"customer_id"`
	Total      float64 `json:"total"`
	Items      []Item  `json:"items"`
}

type Item struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func SendCheckout(url string, checkout Checkout, maxRetries int, initialBackoff time.Duration) (*http.Response, error) {
	data, err := json.Marshal(checkout)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal checkout data: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	var resp *http.Response

	for attempts := 0; attempts <= maxRetries; attempts++ {
		resp, err = client.Do(req)
		if err == nil && (resp.StatusCode >= 200 || resp.StatusCode <= 299) {
			return resp, nil // Successful response
		}

		if resp != nil {
			log.Printf("failed to send checkout data, retrying in %s", initialBackoff)
			resp.Body.Close()
		}

		waitTime := initialBackoff * (1 << attempts)
		time.Sleep(waitTime)
	}

	return nil, fmt.Errorf("failed to send HTTP request after %d attempts: %w", maxRetries+1, err)

}
