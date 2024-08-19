package user

import (
	"context"

	"github.com/google/uuid"
)

func (r *libSqlRepo) Get() (*user, error) {
	u, err := r.db.GetUser(context.Background(), uuid.New())
	if err != nil {
		return nil, err
	}

	user := New(u)

	return &user, nil
}
