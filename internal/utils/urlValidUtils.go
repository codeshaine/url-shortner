package utils

import (
	"log"
	"regexp"
)

func IsValidUrl(url string) bool {
	res, err := regexp.MatchString(`^[a-zA-Z0-9/:_+\?=&.-]+$`, url)
	if err != nil {
		log.Println("Error occured during url validation:", err)
		return false
	}
	return res

}
