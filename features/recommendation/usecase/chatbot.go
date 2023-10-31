package recommendationUseCase

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

func GetCompletionFromMessages(
	ctx context.Context,
	client *openai.Client,
	messages []openai.ChatCompletionMessage,
) (openai.ChatCompletionResponse, error) {

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)
	return resp, err
}
