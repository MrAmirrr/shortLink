package showoriginal

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func ShowOriginalLink(link string, myMap map[string]string) {
	u, err := url.ParseRequestURI(link)
	if err != nil {
		log.Fatalln("can not parse the url")
	}
	s := strings.Split(u.Path, "/")
	f := myMap[s[1]]
	fmt.Println(f)
}
