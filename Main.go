package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrServiceNotFound = errors.New("Service not found")

// Service - структура - веб-сервіс
type Service struct {
	Name             string
	FailureThreshold int
	consecutiveFail  int
}

// Alive - чи сервіс живий
func (s *Service) Alive() error {

	// тут випадковий генератор помилок - імітує стани сервісу
	if s.consecutiveFail < s.FailureThreshold {
		s.consecutiveFail++
		return errors.New("Service is not alive")
	}
	s.consecutiveFail = 0
	return nil
}

// Restart - перезапуск сервісів
func (s *Service) Restart() {
	fmt.Printf("Restarting service: %s\n", s.Name)
}

// HealthChecker - структура-оркестратор для перевірки сервісів
type HealthChecker struct {
	Services []*Service
}

// Check - перевіряє стан сервісів
func (hc *HealthChecker) Check() {
	for _, service := range hc.Services {
		err := service.Alive()
		if err != nil {
			if err == ErrServiceNotFound {
				// відсутній - пропускаємо
				continue
			}

			// Якщо сервіс не живий, перевіряємо FailureThreshold. + рестартуємо
			if service.consecutiveFail >= service.FailureThreshold {
				service.Restart()
			}
		} else {
			// Якщо живий = онулюємо лічильник помилок
			service.consecutiveFail = 0
		}
	}
}

func main() {
	// Створили сервіси
	service1 := &Service{Name: "Service1", FailureThreshold: 3}
	service2 := &Service{Name: "Service2", FailureThreshold: 2}
	service3 := &Service{Name: "Service3", FailureThreshold: 3}
	service4 := &Service{Name: "Service4", FailureThreshold: 2}

	healthChecker := &HealthChecker{
		Services: []*Service{service1, service2, service3, service4},
	}

	// Періодично перевіряємо стан сервісів (кожна 1 хвилина)
	for {
		healthChecker.Check()
		time.Sleep(1 * time.Minute)
	}
}
