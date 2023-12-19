package main

import (
	"fmt"

	"shortLink/shortenLink"
	showoriginal "shortLink/showOriginal"
)

var myMap = make(map[string]string)

func main() {
	counter := 0

	for counter < 1 {
		fmt.Println("1.shorten a link\n2.give the short url for the full url\n3.exit")
		var inputUser int
		fmt.Scan(&inputUser)
		if inputUser == 1 {
			fmt.Println("enter url with scheme")
			var i string
			fmt.Scan(&i)
			shortenLink.ShortenLink(i, myMap)
		} else if inputUser == 2 {
			fmt.Println("enter url with scheme")
			var i string
			fmt.Scan(&i)
			showoriginal.ShowOriginalLink(i, myMap)
		} else if inputUser == 3 {
			fmt.Println("bye")
			break
		} else {
			fmt.Println("the number your enteres is out of options")
		}
	}
}
