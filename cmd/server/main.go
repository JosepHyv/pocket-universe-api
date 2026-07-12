package main

import (
	"fmt"
	"pocket-universe/internal/config"
	"net/http"
)


type Api struct {
	Server http.Server
}

func main(){
	_, err := config.LoadConfig()
	if err != nil {
		fmt.Print(err)
	}

}
