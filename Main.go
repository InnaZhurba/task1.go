package main

import (
	"time"
)

func main() {
	service1 := service.NewService("Service1", 3)
	service2 := service.NewService("Service2", 2)
	service3 := service.NewService("Service3", 3)
	service4 := service.NewService("Service4", 2)

	healthChecker := healthchecker.NewHealthChecker([]*service.Service{service1, service2, service3, service4})

	for {
		healthChecker.Check()
		time.Sleep(1 * time.Minute)
	}
}
