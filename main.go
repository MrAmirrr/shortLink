package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/google/uuid"
)

var myMap = make(map[string]string)

func main() {
	menu()
}

func shortenLink() {
	fmt.Println("Enter the url that you want to make shortlink of:")
	var link string
	fmt.Scan(&link)
	// fmt.Println(link)
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

func showOriginalLink() {
	fmt.Println("enter the short-link to view the original link if it exists: ")
	var link string
	fmt.Scan(&link)
	// fmt.Println(link)
	u, err := url.ParseRequestURI(link)
	if err != nil {
		log.Fatalln("can not parse the url")
	}
	// fmt.Println(u.Path)
	s := strings.Split(u.Path, "/")
	// fmt.Println(s[1])
	f := myMap[s[1]]
	// fmt.Println(myMap)
	fmt.Println(f)
}

func menu() {
	counter := 0

	for counter < 1 {
		fmt.Println("1.shorten a link\n2.give the short url for the full url\n3.exit")
		var inputUser int
		fmt.Scan(&inputUser)
		if inputUser == 1 {
			shortenLink()
		} else if inputUser == 2 {
			showOriginalLink()
		} else if inputUser == 3 {
			fmt.Println("bye")
			break
		} else {
			fmt.Println("the number your enteres is out of options")
		}
	}
}
