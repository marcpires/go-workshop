package main

// Este código é um exemplo, durante os encontros, iremos criar o pacote `client`
import (
	"log"
	"net/http"
)

func main() {
	// Mascaramento de variável (variable shadowing)
	var client *http.Client

	if tracing {
		client, err := NewHTTPClientWithTracing()
		if err != nil {
			return err
		}
		log.Println(client)
	} else {
		client, err := NewHTTPClient()
		if err != nil {
			return err
		}
		log.Println(client)
	}

	// Usar o clients

}
