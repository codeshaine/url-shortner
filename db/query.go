package db

import (
	"errors"
	"log"

	"github.com/codeshaine/url-shortner/internal/utils"
)

type UrlData struct {
	Id         int    `json:"id,omitempty"`
	LongUrl    string `json:"long-url"`
	ShortUrl   string `json:"short-url"`
	ClickCount int    `json:"click-count,omitempty"`
}

func InsertUrl(longUrl, shortUrl string) (UrlData, error) {
	valid := utils.IsValidUrl(longUrl)
	if !valid {
		log.Println("url is not valid")
		return UrlData{}, errors.New("url is not valid")
	}
	_, exErr := Db.Exec("INSERT INTO urls (long_url,short_url) VALUES($1,$2)", longUrl, shortUrl)
	if exErr != nil {
		log.Println("Error occured while checking url exist:", exErr)
		return UrlData{}, exErr
	}
	var urlData UrlData
	res := Db.QueryRow("SELECT long_url,short_url,click FROM urls WHERE short_url=$1", shortUrl)
	err := res.Scan(&urlData.LongUrl, &urlData.ShortUrl, &urlData.ClickCount)
	if err != nil {
		log.Println("error while retrieving the data:", err)
		return UrlData{}, err
	}
	return urlData, nil // ok  as long as you dont exceed the int32 limit
}

func GetLongUrl(shortUrl string) (UrlData, error) { //retrieving the data and updating the counter

	valid := utils.IsValidUrl(shortUrl)
	if !valid {
		log.Println("url is not valid")
		return UrlData{}, errors.New("url is not valid")
	}

	tx, txErr := Db.Begin()
	if txErr != nil {
		log.Println("Error occured while starting transaction")
		tx.Rollback()
		return UrlData{}, txErr
	}

	_, upErr := tx.Exec("UPDATE urls SET click=click+1 WHERE short_url=$1", shortUrl)
	if upErr != nil {
		log.Println("Error occured while updating table in transaction")
		tx.Rollback()
		return UrlData{}, upErr
	}

	var urlData UrlData
	res := tx.QueryRow(`SELECT long_url ,short_url ,click FROM urls WHERE short_url=$1 `, shortUrl)
	scanErr := res.Scan(&urlData.LongUrl, &urlData.ShortUrl, &urlData.ClickCount)
	if scanErr != nil {
		log.Println("Error occured while retrieving data from scanner")
		tx.Rollback()
		return UrlData{}, scanErr
	}
	comErr := tx.Commit()
	if comErr != nil {
		log.Println("Error occured while commtting transaction")
		tx.Rollback()
		return UrlData{}, comErr

	}
	return urlData, nil
}
