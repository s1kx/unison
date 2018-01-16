package client

import (
	"fmt"

	"github.com/s1kx/unison"
)

type serviceRegistry map[string]*unison.Service

// DuplicateServiceError service error
type DuplicateServiceError struct {
	Existing *unison.Service
	New      *unison.Service
	Name     string
}

func (e DuplicateServiceError) Error() string {
	return fmt.Sprintf("service: name '%s' already exists", e.Name)
}
