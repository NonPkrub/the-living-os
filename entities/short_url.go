package entities

import "time"

type URLShorten struct {
	ID          string     `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	OriginalUrl string     `json:"originalUrl"`
	ShortCode   string     `json:"shortCode" gorm:"uniqueIndex"`
	ShortUrl    string     `json:"shortUrl"`
	CreatedAt   time.Time  `json:"createdAt"`
	ExpiresAt   *time.Time `json:"expiresAt"`
	Clicks      int        `json:"clicks"`
}

type URLShortenCreateResponse struct {
	ShortCode   string `json:"shortCode"`
	ShortUrl    string `json:"shortUrl"`
	OriginalUrl string `json:"originalUrl"`
}

type URLShortenGetResponse struct {
	ShortCode   string    `json:"shortCode"`
	OriginalUrl string    `json:"originalUrl"`
	Clicks      int       `json:"clicks"`
	CreatedAt   time.Time `json:"createdAt"`
}

type URLShortenCreateRequest struct {
	OriginalUrl string `json:"originalUrl" validate:"required"`
	ShortCode   string `json:"shortCode" default:"abc123"`
}
