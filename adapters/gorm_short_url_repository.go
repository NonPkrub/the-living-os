package adapters

import (
	"short-url/entities"
	"short-url/usecases"

	"gorm.io/gorm"
)

type GormShortUrlRepository struct {
	DB *gorm.DB
}

func NewGormShortUrlRepository(db *gorm.DB) usecases.URLRepository {
	return &GormShortUrlRepository{DB: db}
}

func (r *GormShortUrlRepository) Save(url *entities.URLShorten) (*entities.URLShortenCreateResponse, error) {
	return &entities.URLShortenCreateResponse{
		ShortCode:   url.ShortCode,
		ShortUrl:    url.ShortUrl,
		OriginalUrl: url.OriginalUrl}, r.DB.Create(url).Error
}

func (r *GormShortUrlRepository) Get(shortCode string) (*entities.URLShortenGetResponse, error) {
	var urlShorten entities.URLShorten
	return &entities.URLShortenGetResponse{
		ShortCode:   urlShorten.ShortCode,
		OriginalUrl: urlShorten.OriginalUrl,
		Clicks:      urlShorten.Clicks,
		CreatedAt:   urlShorten.CreatedAt}, r.DB.Where("short_code = ?", shortCode).First(&urlShorten).Error
}

func (r *GormShortUrlRepository) UpdateCount(shortCode string) (string, error) {
	var urlShorter entities.URLShorten
	return shortCode, r.DB.Model(&urlShorter).Where("short_code = ?", shortCode).Update("clicks", gorm.Expr("clicks + 1")).Error
}

func (r *GormShortUrlRepository) CheckDuplicate(url string, shortCode string) error {
	var urlShorter entities.URLShorten
	err := r.DB.Where("original_url = ? AND short_code = ?", url, shortCode).First(&urlShorter).Error
	if err != nil && err.Error() == "record not found" {
		return nil
	}
	return err
}
