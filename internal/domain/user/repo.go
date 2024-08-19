package user

type RepoReader interface {
	Get() (*user, error)
}
