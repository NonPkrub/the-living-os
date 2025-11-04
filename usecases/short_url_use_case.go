package usecases

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"short-url/entities"
)

type URLUseCase interface {
	CreatedShortUrl(url *entities.URLShortenCreateRequest) (*entities.URLShortenCreateResponse, error)
	GetShortUrl(url string) (*entities.URLShortenGetResponse, error)
	Redirect(url string) (string, error)
}

type ShortURLService struct {
	repo URLRepository
}

func NewShortURLService(repo URLRepository) URLUseCase {
	return &ShortURLService{
		repo: repo,
	}
}

func (s *ShortURLService) CreatedShortUrl(url *entities.URLShortenCreateRequest) (*entities.URLShortenCreateResponse, error) {
	err := s.repo.CheckDuplicate(url.OriginalUrl, url.ShortCode)
	if err != nil {
		return nil, err
	}

	linkGenerator, _ := GenerateShortURL()

	req := &entities.URLShorten{
		OriginalUrl: url.OriginalUrl,
		ShortCode:   url.ShortCode,
		ShortUrl:    fmt.Sprintf("%s/%s", url.ShortCode, linkGenerator),
		ExpiresAt:   nil,
		Clicks:      0,
	}
	return s.repo.Save(req)
}

func (s *ShortURLService) GetShortUrl(url string) (*entities.URLShortenGetResponse, error) {
	return s.repo.Get(url)
}

func (s *ShortURLService) Redirect(url string) (string, error) {
	return s.repo.UpdateCount(url)
}

func GenerateShortURL() (string, error) {
	randomBytes := make([]byte, 6)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	linkGenerator := base64.URLEncoding.EncodeToString(randomBytes)[:8]
	return linkGenerator, nil
}
