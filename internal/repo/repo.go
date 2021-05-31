package repo

import (
	"github.com/ozoncp/ocp-role-api/internal/model"
)

type Repo interface {
	AddRoles(r []*model.Role) error
	RemoveRole(r *model.Role) error
	ListRoles(limit, offset uint) ([]*model.Role, error)
}

type roleStorageMem struct {
	roles []*model.Role
}

type RepoErr uint

const (
	AlreadyExistsErr RepoErr = iota
	NotFoundErr
)

func NewRoleStorageMem() Repo {
	return &roleStorageMem{}
}

func (err RepoErr) Error() string {
	return "RepoAlreadyExistsError"
}

func (r *roleStorageMem) AddRoles(roles []*model.Role) error {
	var i int
	for _, role := range roles {
		if i = r.search(role); i == -1 {
			return AlreadyExistsErr
		}
	}
	r.roles = append(r.roles, roles...)
	return nil
}

func (r *roleStorageMem) RemoveRole(role *model.Role) error {
	var i int
	if i = r.search(role); i == -1 {
		return NotFoundErr
	}
	r.roles[i] = r.roles[len(r.roles)]
	r.roles = r.roles[:len(r.roles)-1]
	return nil
}

func (r *roleStorageMem) ListRoles(limit, offset uint) ([]*model.Role, error) {
	min := func(i, j int) int {
		if i <= j {
			return i
		}
		return j
	}
	i := min(int(offset), len(r.roles))
	j := min(i+int(offset), len(r.roles))
	return r.roles[i:j], nil
}

func (r *roleStorageMem) search(role *model.Role) int {
	for i := 0; i < len(r.roles); i++ {
		if r.roles[i] == role {
			return i
		}
	}
	return -1
}
