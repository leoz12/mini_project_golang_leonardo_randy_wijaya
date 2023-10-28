package recommendationHandler

import "mini_project/features/recommendation"

type CreateRequest struct {
	Name string `json:"name" form:"name"`
}

type UpdateRequest struct {
	Name string `json:"name" form:"name"`
}

type RecommendationResponse struct {
	Id   string
	Name string
}

func CoreToResponse(data recommendation.Core) RecommendationResponse {
	return RecommendationResponse{
		Id:   data.Id,
		Name: data.Name,
	}
}
