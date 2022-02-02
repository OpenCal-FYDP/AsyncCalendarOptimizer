package main

import (
	"github.com/OpenCal-FYDP/AsyncCalendarOptimizer/internal/service"
	"github.com/OpenCal-FYDP/AsyncCalendarOptimizer/rpc"
	"log"
	"net/http"
)

func main() {
	service := service.New()
	server := rpc.NewAsyncCalendarOptimizerServiceServer(service)
	log.Fatal(http.ListenAndServe(":8080", server))
}
