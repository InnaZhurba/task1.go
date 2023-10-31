package Service

import (
	"errors"
	"fmt"
)

var ErrServiceNotFound = errors.New("Service not found")

type Service struct {
	Name             string
	FailureThreshold int
	consecutiveFail  int
}

func NewService(name string, failureThreshold int) *Service {
	return &Service{
		Name:             name,
		FailureThreshold: failureThreshold,
	}
}

func (s *Service) Alive() error {
	if s.consecutiveFail < s.FailureThreshold {
		s.consecutiveFail++
		return errors.New("Service is not operational")
	}
	s.consecutiveFail = 0
	return nil
}

func (s *Service) Restart() {
	fmt.Printf("Restarting service: %s\n", s.Name)
}
