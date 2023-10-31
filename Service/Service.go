package Service

import (
	"errors"
	"fmt"
)

var ErrServiceNotFound = errors.New("Service not found")

type Service struct {
	Name             string
	FailureThreshold int
	FailCounter      int
}

func NewService(name string, failureThreshold int) *Service {
	return &Service{
		Name:             name,
		FailureThreshold: failureThreshold,
	}
}

func (s *Service) Alive() error {
	if s.FailCounter < s.FailureThreshold {
		s.FailCounter++
		return errors.New("service is not alive")
	}
	s.FailCounter = 0
	return nil
}

func (s *Service) Restart() {
	fmt.Printf("Restarting service: %s\n", s.Name)
}
