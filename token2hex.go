package main

import "encoding/hex"
//import "encoding/base64"
import "fmt"
import "os"

func main () {
//	data, err := base64.StdEncoding.DecodeString ( os.Args[1] )
//	if err !=nil {
//		fmt.Println ( "Errr string" )
//	}
//	fmt.Println ( data )
	src := []byte( os.Args[1]  )
	fmt.Println( hex.EncodeToString( src))
}
