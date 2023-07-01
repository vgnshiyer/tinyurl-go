package routes

import(
	"time"
)

type request struct {
	URL             string			`json."url"`
	CustomShort     string			`json:"short"`
	Expiry 			time.Duration   `json:"expiry"`
}

type response struct {
	URL			     	string				`json."url"`
	CustomShort			string				`json:"short"`
	Expiry				time.Duration		`json:"expiry"`
	XRateRemaining		int					`json:"rate_limit"`
	XRateLimitReset		time.Duration		`json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map("error":"cannot parse JSON"))
	}

	// rate limiting

	// check url
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map("error":"URL not found"))
	}

	// check domain error
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map("error":"Unkown error"))
	}

	// enforce http, SSL
	body.URL = helpers.EnforceHTTP(body.URL)
}