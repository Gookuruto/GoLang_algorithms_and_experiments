package routers

import (
	"net/http"
	"workspace/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/", handlers.HandleRoot)
	http.HandleFunc("/hello", handlers.HandleHello)
	http.HandleFunc("/ws", handlers.HandleWSConnection)
}
