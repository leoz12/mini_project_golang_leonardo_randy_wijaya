package recommendationUseCase

import (
	"context"
	"errors"
	"mini_project/app/configs"
	"mini_project/features/recommendation"

	openai "github.com/sashabaranov/go-openai"
)

type recommendationUsecase struct {
	recommendationRepository recommendation.DataInterface
}

// RecommendGame implements recommendation.UseCaseInterface.
func (uc *recommendationUsecase) RecommendGame(id string) (string, error) {
	resp, err := uc.recommendationRepository.SelectById(id)

	if id == "" {
		return "", errors.New("id is required")
	}

	if err != nil {
		return "", err

	}

	client := openai.NewClient(configs.OPEN_AI_KEY)
	ctx := context.Background()
	message := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: resp.Name,
		},
	}
	resultBot, errBot := GetCompletionFromMessages(ctx, client, message)

	if errBot != nil {
		return "", errBot
	}

	return resultBot.Choices[0].Message.Content, nil
}

// GetAll implements recommendation.UseCaseInterface.
func (uc *recommendationUsecase) GetAll() ([]recommendation.Core, error) {
	resp, err := uc.recommendationRepository.SelectAll()

	return resp, err
}

// GetById implements recommendation.UseCaseInterface.
func (uc *recommendationUsecase) GetById(id string) (recommendation.Core, error) {
	resp, err := uc.recommendationRepository.SelectById(id)

	if id == "" {
		return recommendation.Core{}, errors.New("id is required")
	}

	return resp, err
}

// Insert implements recommendation.UseCaseInterface.
func (uc *recommendationUsecase) Insert(data recommendation.Core) (recommendation.Core, error) {
	if data.Name == "" {
		return recommendation.Core{}, errors.New("name is required")
	}
	response, err := uc.recommendationRepository.Insert(data)

	return response, err
}

// Update implements recommendation.UseCaseInterface.
func (uc *recommendationUsecase) Update(id string, data recommendation.Core) error {
	if data.Name == "" {
		return errors.New("name is required")
	}
	err := uc.recommendationRepository.Update(id, data)

	return err
}

// Delete implements recommendation.UseCaseInterface.
func (uc *recommendationUsecase) Delete(id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	err := uc.recommendationRepository.Delete(id)
	return err
}

func New(recommendationRepo recommendation.DataInterface) recommendation.UseCaseInterface {
	return &recommendationUsecase{
		recommendationRepository: recommendationRepo,
	}
}
