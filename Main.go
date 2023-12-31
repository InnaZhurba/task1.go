package main

import (
	"task1.go/Healthchecker"
	"task1.go/Service"
	"time"
)

func main() {
	service1 := Service.NewService("Service1", 3)
	service2 := Service.NewService("Service2", 2)
	service3 := Service.NewService("Service3", 3)
	service4 := Service.NewService("Service4", 2)
	service5 := Service.NewService("Service5", 1)
	service6 := Service.NewService("Service6", 1)
	service7 := Service.NewService("Service7", 5)
	service8 := Service.NewService("Service8", 3)

	healthChecker := Healthchecker.NewHealthChecker([]*Service.Service{service1, service2, service3, service4, service5, service6, service7, service8})

	for {
		healthChecker.Check()
		time.Sleep(1 * time.Minute)
	}
}
