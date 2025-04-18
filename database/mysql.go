package database

import (
	"UrlShortner/config"
)

func SaveUrl(shortKey, originalUrl string) error {
	_, err := config.DB.Exec("INSERT INTO urls(shortKey, originalUrl) VALUES(?, ?)", shortKey, originalUrl)
	return err
}

func GetUrl(shortKey string) (string, error) {
	var url string
	err := config.DB.QueryRow("SELECT url FROM urls WHERE shortKey=?", shortKey).Scan(&url)
	return url, err
}
