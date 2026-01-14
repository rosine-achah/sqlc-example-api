// // package api

// // import (
// // 	"bytes"
// // 	"encoding/json"
// // 	"net/http"
// // )

// // type CampayClient struct {
// // 	Token string
// // }

// // type CampayRequest struct {
// // 	Amount      int    `json:"amount"`
// // 	Currency    string `json:"currency"`
// // 	From        string `json:"from"`
// // 	Description string `json:"description"`
// // }

// // func (c *CampayClient) CollectPayment(req CampayRequest) error {
// // 	body, err := json.Marshal(req)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	httpReq, err := http.NewRequest(
// // 		"POST",
// // 		"https://api.campay.net/api/collect/",
// // 		bytes.NewBuffer(body),
// // 	)
// // 	if err != nil {
// // 		return err
// // 	}

// // 	httpReq.Header.Set("Authorization", "Token "+c.Token)
// // 	httpReq.Header.Set("Content-Type", "application/json")

// // 	client := &http.Client{}
// // 	_, err = client.Do(httpReq)
// // 	return err
// // }

// package api

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// type CampayClient struct {
// 	Token string
// }

// type CampayRequest struct {
// 	Amount      int    `json:"amount"`
// 	Currency    string `json:"currency"`
// 	From        string `json:"from"`
// 	Description string `json:"description"`
// }

// func (c *CampayClient) CollectPayment(req CampayRequest) error {
// 	body, err := json.Marshal(req)
// 	if err != nil {
// 		return err
// 	}

// 	httpReq, err := http.NewRequest(
// 		"POST",
// 		"https://api.campay.net/api/collect/",
// 		bytes.NewBuffer(body),
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	httpReq.Header.Set("Authorization", "Token "+c.Token)
// 	httpReq.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(httpReq)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// Print response from Campay for debugging
// 	respBody, _ := io.ReadAll(resp.Body)
// 	fmt.Println("📨 Campay response:", string(respBody))

// 	return nil
// }

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type CampayClient struct {
	Token string
}

type CampayRequest struct {
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
	From        string `json:"from"`
	Description string `json:"description"`
}

func (c *CampayClient) CollectPayment(req CampayRequest) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest(
		"POST",
		"https://api.campay.net/api/collect/",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	httpReq.Header.Set("Authorization", "Token "+c.Token)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return err
	}

	// defer resp.Body.Close()

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("failed to close response body:", err)
		}
	}()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("📨 Campay response:", string(respBody))

	return nil
}
