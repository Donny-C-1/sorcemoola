package services

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"os"
)

type PaystackService struct {
	SecretKey string
	ApiUrl    string
}

func NewPaystackService() *PaystackService {
	return &PaystackService{
		SecretKey: os.Getenv("PAYSTACK_SECRET"),
		ApiUrl:    "https://api.paystack.co/transaction/initialize",
	}
}

func (ps *PaystackService) Initialize(email string, amount int64, metadata map[string]any) (map[string]any, error) {
	body := map[string]any{
		"amount":   amount,
		"email":    email,
		"metadata": metadata,
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", ps.ApiUrl, bytes.NewBuffer((jsonBody)))
	req.Header.Set("Authorization", "Bearer "+ps.SecretKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]any
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

func (ps *PaystackService) VerifySignature(body []byte, signature string) bool {
	hash := hmac.New(sha512.New, []byte(ps.SecretKey))
	hash.Write(body)
	expected := hex.EncodeToString(hash.Sum(nil))
	return expected == signature
}
