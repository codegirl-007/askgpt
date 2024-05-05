package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ayush6624/go-chatgpt"
	"github.com/spf13/cobra"
)

var askgpt = &cobra.Command{
	Use:   "ask",
	Short: "Ask GPT a question",
	Long:  "Ask a question to GPT",
	Args:  cobra.ExactArgs(1),
	Run:   ask,
}

func ask(cmd *cobra.Command, args []string) {
	question := args[0]
	key := os.Getenv("OPENAI_KEY")

	c, err := chatgpt.NewClient(key)
	if err != nil {
		log.Fatal("ERROR", err)
	}

	ctx := context.Background()

	res, error := c.Send(ctx, &chatgpt.ChatCompletionRequest{
		Model: chatgpt.GPT4,
		Messages: []chatgpt.ChatMessage{
			{
				Role:    chatgpt.ChatGPTModelRoleSystem,
				Content: question,
			},
		},
	})
	if error != nil {
		log.Fatal("ERROR", error)
	}

	fmt.Println(res.Choices[0].Message.Content)
}
