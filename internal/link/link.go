package link

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"time"
)

type Link struct {
	ID        int64     `json:"id" db:"id"`
	URL       string    `json:"url" db:"url"`
	Code      string    `json:"code" db:"code"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CreateLinkReq struct {
	URL string `json:"url"`
}

func (req *CreateLinkReq) Validate() error {
	if req.URL == "" {
		return fmt.Errorf("param: %s (type: %s) is required", "url", "string")
	}

	_, err := url.Parse(req.URL)

	if err != nil {
		return fmt.Errorf("invalid url format: %s", req.URL)
	}

	return nil
}

type CreateLinkRes struct {
	URL string `json:"url"`
}

type GetLinkByCodeReq struct {
	Code string `json:"code"`
}

func (req *GetLinkByCodeReq) Validate() error {
	if req.Code == "" {
		return fmt.Errorf("param: %s (type: %s) is required", "code", "string")
	}

	rgx := regexp.MustCompile(`^\w{6}$`)

	isValidCode := rgx.MatchString(req.Code)
	if !isValidCode {
		return fmt.Errorf("invalid code format: %s. The code must be exactly 6 alphanumeric characters", req.Code)
	}

	return nil
}

type GetLinkByCodeRes struct {
	URL string `json:"url"`
}

type Repository interface {
	CreateLink(ctx context.Context, user *Link) (*Link, error)
	GetLinkByCode(ctx context.Context, code string) (*Link, error)
}

type Service interface {
	CreateLink(ctx context.Context, req *CreateLinkReq) (*CreateLinkRes, error)
	GetLinkByCode(ctx context.Context, code string) (*GetLinkByCodeRes, error)
}

type Code interface {
	GenerateUniqueCode() string
}
