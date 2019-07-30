package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	fmt.Printf("books  : %#v\n", books)

	// var (
	// 	wBooks [winter]string
	// 	sBooks [summer]string
	// )

	// wBooks[0] = books[0]

	// for i, b := range books[1:] {
	// 	sBooks[i] = b
	// }

	// fmt.Printf("\nwinter: %#v\n", wBooks)
	// fmt.Printf("\nsummer: %#v\n", sBooks)

	// var published [len(books)]bool

	// published[0] = true
	// published[len(books)-1] = true

	// fmt.Println("\nPublished books: ")
	// for i, ok := range published {
	// 	if ok {
	// 		fmt.Printf("+ %s\n", books[i])
	// 	}
	// }
}
