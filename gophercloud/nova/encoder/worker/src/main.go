// TODO

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Video URL and/or format is required!")
		return
	}

	video := args[0]
	format := args[1]

	fmt.Printf("Encoding %s video to %s format \n", video, format)

}
