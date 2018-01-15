package session

import "fmt"

type serviceRegistry map[string]*Service

// DuplicateServiceError service error
type DuplicateServiceError struct {
	Existing *Service
	New      *Service
	Name     string
}

func (e DuplicateServiceError) Error() string {
	return fmt.Sprintf("service: name '%s' already exists", e.Name)
}
