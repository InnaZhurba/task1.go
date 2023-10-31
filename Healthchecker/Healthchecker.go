package Healthchecker

type HealthChecker struct {
	Services []*service.Service
}

func NewHealthChecker(services []*service.Service) *HealthChecker {
	return &HealthChecker{Services: services}
}

func (hc *HealthChecker) Check() {
	for _, service := range hc.Services {
		err := service.Alive()
		if err != nil {
			if err == service.ErrServiceNotFound {
				continue
			}

			if service.consecutiveFail >= service.FailureThreshold {
				service.Restart()
			}
		} else {
			service.consecutiveFail = 0
		}
	}
}
