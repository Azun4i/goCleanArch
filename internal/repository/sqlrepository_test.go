package repository_test

import (
	"github.com/stretchr/testify/assert"
	"goCleanArch/internal/model"
	"goCleanArch/internal/repository"
	"testing"
)

func TestSqlstore_Create(t *testing.T) {
	db, teardown := repository.TestDB(t, databaseURL)
	defer teardown("author")

	s := repository.NewSqlstore(db)
	u := model.Testuser(t)

	assert.NoError(t, s.Create(u))
	assert.NotNil(t, u)
}

func TestSqlstore_Delete(t *testing.T) {
	db, teardown := repository.TestDB(t, databaseURL)
	defer teardown("author")

	s := repository.NewSqlstore(db)
	u := model.Testuser(t)

	assert.NoError(t, s.Delete(u.ID))
	assert.NotNil(t, u)
}
