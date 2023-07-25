package link

import (
	"context"
	"fmt"
	"os"
	"time"
)

type service struct {
	Repository
	Code
	timeout time.Duration
}

func NewService(repository Repository, code Code) Service {
	return &service{repository, code, time.Duration(2) * time.Second}
}

func (s *service) CreateLink(ctx context.Context, req *CreateLinkReq) (*CreateLinkRes, error) {

	link := &Link{
		URL:  req.URL,
		Code: s.GenerateUniqueCode(),
	}

	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	link, err := s.Repository.CreateLink(ctx, link)

	if err != nil {
		return nil, err
	}

	baseURL := os.Getenv("BASE_URL")
	URL := fmt.Sprintf("%s/%s", baseURL, link.Code)

	return &CreateLinkRes{
		URL: URL,
	}, err
}

func (s *service) GetLinkByCode(ctx context.Context, code string) (*GetLinkByCodeRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	link, err := s.Repository.GetLinkByCode(ctx, code)

	if err != nil {
		return nil, err
	}

	res := &GetLinkByCodeRes{
		URL: link.URL,
	}

	return res, nil
}
