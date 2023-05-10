package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	openaiURL = "https://api.openai.com/v1/engines/davinci/completions"
	apiKey    = "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"
)

type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type OpenaiRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type OpenaiResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Choices []Choice `json:"choices"`
	Usages  Usage    `json:"usage"`
}

type Choice struct {
	Index        int     `json:"index"`
	Messages     Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// チャットGPTとちゃっとする
func main() {

	reader := bufio.NewReader(os.Stdin)

	var message []Message

	fmt.Println("チャットを始めます。")
	fmt.Println("終了するには「exit」と入力してください。")

	for {
		fmt.Print("あなた：")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			break
		}
		message = append(message, Message{
			Role:    "User",
			Content: text,
		})

	}
	response := getOpenAIResponse("", message)
	fmt.Println(response)

	fmt.Println("チャットを終了します。")

}

func getOpenAIResponse(apiKey string, messages []Message) OpenaiResponse {
	requestBody := OpenaiRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
	}

	requestJSON, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(requestJSON))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response OpenaiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		println("Error: ", err.Error())
		return OpenaiResponse{}
	}

	fmt.Println(response)
	// messages = append(messages, Message{
	// 	Role:    "assistant",
	// 	Content: response.Choices[0].Messages.Content,
	// })

	return response
}
