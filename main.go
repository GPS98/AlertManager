package main

import (
	"fmt"
	"log"
	//"alertmanager/storage"
)

func main() {
	store, err := InMemoryStorage()
	if err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":3000", store)
	server.Run()

	fmt.Println("Hello, Golang!")
}

// export GOROOT=/usr/local/go
// export GOPATH=/Users/prems/code/
// export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
