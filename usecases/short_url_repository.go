package usecases

import "short-url/entities"

type URLRepository interface {
	Save(url *entities.URLShorten) (*entities.URLShortenCreateResponse, error)
	Get(url string) (*entities.URLShortenGetResponse, error)
	CheckDuplicate(url string, shortCode string) error
	UpdateCount(url string) (string, error)
}
