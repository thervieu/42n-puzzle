package server

import (
	"fmt"
	"log"
	"strconv"

	"net/http"

	"github.com/thervieu/42n-puzzle/utils"
)


func Start() {
	port := utils.ReadPort()

	mux := http.NewServeMux()

	// register the solve handler to the mux
	mux.HandleFunc("/solve", solveRouteHandler)

	// Wrapp mux in the CORS middleware
	wrapped_mux := AllowCORS(mux)
	addr := ":" + strconv.Itoa(port)

	fmt.Println("start api on port :" + strconv.Itoa(port))
	log.Fatal(http.ListenAndServe(addr, wrapped_mux))
}