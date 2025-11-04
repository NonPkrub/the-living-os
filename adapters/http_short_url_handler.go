package adapters

import (
	"net/http"
	"short-url/entities"
	"short-url/usecases"

	"github.com/gofiber/fiber/v2"
)

type HttpShortUrlHandler struct {
	urlUseCase usecases.URLUseCase
}

func NewHttpShortUrlHandler(urlUseCase usecases.URLUseCase) *HttpShortUrlHandler {
	return &HttpShortUrlHandler{urlUseCase: urlUseCase}
}

// @Summary Create a short url
// @Description Create a short url
// @Tags ShortURL
// @Accept json
// @Produce json
// @Param request body entities.URLShortenCreateRequest true "request"
// @Success 201 {object} entities.URLShortenCreateResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/shorten [post]
func (h *HttpShortUrlHandler) CreateShortUrl(c *fiber.Ctx) error {
	var request entities.URLShortenCreateRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.urlUseCase.CreatedShortUrl(&request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"short_url": res.ShortCode, "short_code": res.ShortUrl, "original_url": res.OriginalUrl})
}

// @Summary Get statistics for a short URL
// @Description Returns information such as clicks and createdAt
// @Tags ShortURL
// @Param shortCode path string true "Short Code"
// @Produce json
// @Success 200 {object} entities.URLShortenGetResponse
// @Failure 404 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/stats/{shortCode} [get]
func (h *HttpShortUrlHandler) GetShortUrl(c *fiber.Ctx) error {
	code := c.Params("shortCode")

	if code == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "shortCode is required"})
	}

	urlShorter, err := h.urlUseCase.GetShortUrl(code)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"short_code": urlShorter.ShortCode, "original_url": urlShorter.OriginalUrl, "clicks": urlShorter.Clicks, "created_at": urlShorter.CreatedAt})
}

// @Summary Redirect to the original URL
// @Description Redirect user using short code
// @Tags ShortURL
// @Param shortCode path string true "Short Code"
// @Success 301
// @Failure 404 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/{shortCode} [get]
func (h *HttpShortUrlHandler) Redirect(c *fiber.Ctx) error {
	code := c.Params("shortCode")

	if code == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "shortCode is required"})
	}

	originalUrl, err := h.urlUseCase.Redirect(code)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	if originalUrl == "" {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "original url not found"})
	}

	return c.Status(http.StatusMovedPermanently).Redirect(originalUrl)
}
