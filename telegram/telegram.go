package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

var (
	Token  string
	ChatId string
)

func getUrl() string {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	Token := os.Getenv("TOKEN")
	return fmt.Sprintf("https://api.telegram.org/bot%s", Token)
}

func SendMessage(text string) (bool, error) {
  envErr := godotenv.Load()
  if envErr != nil {
    log.Fatal("Error loading .env file")
  }

	ChatId := os.Getenv("CHAT_ID")

	// Global variables
	var err error
	var response *http.Response

	// Send the message
	url := fmt.Sprintf("%s/sendMessage", getUrl())
	body, _ := json.Marshal(map[string]string{
		"chat_id": ChatId,
		"text":    text,
	})
	response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return false, err
	}

	// Close the request at the end
	defer response.Body.Close()

	// Body
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	// Log
	log.Infof("Message '%s' was sent", text)
	log.Infof("Response JSON: %s", string(body))

	// Return
	return true, nil
}

// Find the chat id
// curl -s https://api.telegram.org/bot${TOKEN}/getUpdates
