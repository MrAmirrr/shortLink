package shortenLink

import (
	"fmt"
	"log"
	"net/url"

	"github.com/google/uuid"
)

func ShortenLink(link string, myMap map[string]string) {
	_, err := url.ParseRequestURI(link)
	if err != nil {
		log.Fatalln("can not parse the url")
	}
	// fmt.Println(u.Host)
	newUUID := uuid.New()
	strUuid := newUUID.String()
	// fmt.Println(strUuid)2
	myMap[strUuid] = link
	fmt.Print("here is the shortlink:\n|||||||||||||||||||||| \nhttp://shorten.sh/")
	fmt.Println(newUUID)
}
