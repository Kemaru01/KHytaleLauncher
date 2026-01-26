package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

func GenerateDummyJwt(uuid, name string) string {
	base64Url := func(input string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(input))
	}

	headerJSON, _ := json.Marshal(map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	})
	header := base64Url(string(headerJSON))

	iat := time.Now().Unix()
	exp := iat + 60*60*24*30

	payloadData := map[string]interface{}{
		"sub":   uuid,
		"name":  name,
		"scope": "hytale:client",
		"iat":   iat,
		"exp":   exp,
	}

	payloadJSON, err := json.Marshal(payloadData)
	if err != nil {
		return ""
	}
	payload := base64Url(string(payloadJSON))

	signature := base64Url("fake_signature_for_insecure_mode")

	return fmt.Sprintf("%s.%s.%s", header, payload, signature)
}
