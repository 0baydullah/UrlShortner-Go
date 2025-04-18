package database

import (
	"UrlShortner/config"
)

func SaveUrl(shortKey, originalUrl string) error {
	_, err := config.DB.Exec
}
