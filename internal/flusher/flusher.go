package flusher

import (
	"github.com/ozoncp/ocp-role-api/internal/model"
	"github.com/ozoncp/ocp-role-api/internal/repo"
)

type Flusher interface {
	Flush([]*model.Role) []*model.Role
}

type flusher struct {
	chunkSize int
	roleRepo  repo.Repo
}

func New(chunkSize int, repo repo.Repo) Flusher {
	return &flusher{chunkSize, repo}
}

func (f *flusher) Flush(roles []*model.Role) []*model.Role {
	for i := 0; i < len(roles); i += f.chunkSize {
		j := i + f.chunkSize
		if j > len(roles) {
			j = len(roles)
		}
		chunk := roles[i:j]

		if err := f.roleRepo.AddRoles(chunk); err != nil {
			return roles[i:]
		}
	}
	return nil
}
