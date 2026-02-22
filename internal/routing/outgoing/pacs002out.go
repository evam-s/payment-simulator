package routing

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func SendPacs002Out(pacs002XmlWithHeaders, url string) error {
	xml := []byte(pacs002XmlWithHeaders)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(xml))
	if err != nil {
		log.Println("Failed to create a New Http request:", err)
		return fmt.Errorf("Failed to create a New Http request: %w", err)
	}
	request.Header.Set("Content-Type", "application/xml")
	client := &http.Client{}
	if res, err := client.Do(request); err != nil {
		log.Println("PACS002 posting failed:", err)
		return fmt.Errorf("PACS002 posting failed: %w", err)
	} else {
		if res.StatusCode == 200 || res.StatusCode == 202 || res.StatusCode == 204 {
			bodyBytes, _ := io.ReadAll(res.Body)
			log.Println("PACS002 posted Successfully; Response:", string(bodyBytes), "Status Code:", res.StatusCode)
			res.Body.Close()
			return nil
		} else {
			log.Println("Invalid Status Code recieved during PACS002 Posting:", res.StatusCode)
			return fmt.Errorf("Invalid Status Code recieved during PACS002 Posting: %v", res.StatusCode)
		}
	}
}
