package model

import (
	"fmt"
	"strings"
)

type Role struct {
	Id        uint64
	Service   string
	Operation string
}

type ServiceRoles struct {
	Service    string
	Operations []string
}

func (r *Role) String() string {
	return fmt.Sprintf("Role{Service: %s, Operation: %s}", r.Service, r.Operation)
}

func (sr *ServiceRoles) String() string {
	var b strings.Builder

	for i, op := range sr.Operations {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(op)
	}

	return fmt.Sprintf("ServiceRoles{Service: %s, Operations: [%s]}", sr.Service, b.String())
}
