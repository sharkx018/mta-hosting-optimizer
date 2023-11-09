package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/mta-hosting-optimizer/internal/constant"
	"github.com/mta-hosting-optimizer/internal/handler"
	"github.com/mta-hosting-optimizer/internal/repo"
	"github.com/mta-hosting-optimizer/internal/usecase"
	"os"

	"net/http"
)

func main() {

	// Init the repo service
	ipConfigRepo := repo.New()

	// Init the use case
	inefficientServersUc := usecase.New(ipConfigRepo, getThreshold())

	// Init the handlerResource
	handlerResource := handler.New(inefficientServersUc)

	// Init the router
	router := chi.NewRouter()
	router.Get("/get-inefficient-mtas", handlerResource.GetInefficientMTAsHandler)

	// Init the server
	fmt.Printf("Application started at port %s\n", constant.ConfigPort)
	err := http.ListenAndServe(constant.ConfigPort, router)
	if err != nil {
		fmt.Println("Error while starting the server", err.Error())
	}

}

func getThreshold() int {
	threshold := constant.ThresholdNumber // Default threshold
	envThreshold := os.Getenv("THRESHOLD")
	if envThreshold != "" {
		fmt.Sscan(envThreshold, &threshold)
	}
	return threshold
}
