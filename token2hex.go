package main

import "encoding/hex"
import "fmt"
import "os"

func main () {

	if len (os.Args) < 2 {
		fmt.Println ( "Usage: token2hex [DEVICE-TOKEN]" )
		return
	}
	deviceToken := os.Args[1]

	src := []byte( deviceToken ) 

	fmt.Println( hex.EncodeToString( src ))
}
