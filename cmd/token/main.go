package main

import (
	"github.com/numbermess/token-extractor/pkg/token"
	"log"
)

func main() {

	value, err := token.ExtractToken("Basic YW5pdGFAbmV3ZHVjay5jb206UGFzc3dvcmQx")
	if err != nil {
		log.Panic(err)
	}

	log.Printf("%+v", value)

	value, err = token.ExtractToken("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImFuaXRhQG5ld2R1Y2suY29tIiwiVHlwZSI6IkFjY2VzcyIsImV4cCI6MTU2NjI0MTYwMH0.p61bNfzfQzjnSDje42kjSwViixncwQ_xPWS2RW5y7eo")
	if err != nil {
		log.Panic(err)
	}

	log.Printf("%+v", value)

}
