package main

import "encoding/hex"
import "fmt"
import "os"

func main () {
	src := []byte( os.Args[1]  )
	fmt.Println( hex.EncodeToString( src))
}
