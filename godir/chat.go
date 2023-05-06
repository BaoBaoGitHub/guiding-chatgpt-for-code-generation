package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/xyhelper/chatgpt-go"
)

func main() {
	token := uuid.New().String()
	cli := chatgpt.NewClient(
		chatgpt.WithTimeout(120*time.Second),
		chatgpt.WithAccessToken(token),
		chatgpt.WithBaseURI("https://freechat2.xyhelper.cn"),
		chatgpt.WithModel("gpt-3.5-turbo"),
	)
	conversationID := ""
	parentMessage := ""

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Please input your query: ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		stream, err := cli.GetChatStream(message, conversationID, parentMessage)
		if err != nil {
			log.Fatalf("get chat stream failed: %v\n", err)
		}

		var answer string
		for text := range stream.Stream {
			print(strings.Replace(text.Content, answer, "", 1))

			answer = text.Content
			conversationID = text.ConversationID
			parentMessage = text.MessageID
		}

		if stream.Err != nil {
			log.Fatalf("stream closed with error: %v\n", stream.Err)
		}
		println()
	}
}
