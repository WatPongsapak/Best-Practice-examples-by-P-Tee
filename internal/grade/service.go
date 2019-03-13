package grade

import (
	"context"
	"strconv"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s Service) Calculate(ctx context.Context, score string) (CalculateResponse, error) {
	pi, err := strconv.Atoi(score)
	if err != nil {
		return CalculateResponse{}, err
	}

	grade := "F"
	switch {
	case pi > 60 && pi < 70:
		grade = "D"
	case pi > 71 && pi < 80:
		grade = "C"
	case pi > 81 && pi < 90:
		grade = "B"
	case pi > 91:
		grade = "A"
	}
	return CalculateResponse{Score: pi, Grade: grade}, nil
}
