package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct{}

func Start(host string, port int, startMessage ...string) {

	if len(startMessage) > 0 {
		message := startMessage[0]
		message = strings.ReplaceAll(message, "{host}", host)
		message = strings.ReplaceAll(message, "{port}", strconv.Itoa(port))
		fmt.Printf(message)
	}

	err := http.ListenAndServe(host+":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		return
	}
}
