package Healthchecker

import (
	"task1.go/Service"
)

type HealthChecker struct {
	Services []*Service.Service
}

func NewHealthChecker(services []*Service.Service) *HealthChecker {
	return &HealthChecker{Services: services}
}

func (hc *HealthChecker) Check() {
	for _, service := range hc.Services {
		err := service.Alive()
		if err != nil {
			if err == service.ErrServiceNotFound {
				continue
			}

			if service.ConsecutiveFail >= service.FailureThreshold {
				service.Restart()
			}
		} else {
			service.ConsecutiveFail = 0
		}
	}
}
