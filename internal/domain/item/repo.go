package item

import "github.com/google/uuid"

type RepoWriter interface {
	Create(name, description string) (*Item, error)
	Delete(id uuid.UUID) error
	Update(id uuid.UUID, name, description string) (*Item, error)
}
