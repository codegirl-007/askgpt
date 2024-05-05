/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ayush6624/go-chatgpt"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "askgpt",
	Short: "A brief description of your application",
	Long:  `A longer description`,
	Args:  cobra.ExactArgs(1),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
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
					Content: question + " Limit it to 5 sentences.",
				},
			},
		})
		if error != nil {
			log.Fatal("ERROR", error)
		}

		fmt.Println("\n", res.Choices[0].Message.Content)
	},
}

func init() {
	rootCmd.AddCommand(askgpt)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.askgpt.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
