package item

type RepoWriter interface {
	Create(name, description string) (*Item, error)
	Delete(id int64) error
	Update(id int64, name, description string) (*Item, error)
}
