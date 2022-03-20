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

func TestSqlstore_FindById(t *testing.T) {
	db, teardown := repository.TestDB(t, databaseURL)
	defer teardown("author")

	s := repository.NewSqlstore(db)
	u := model.Testuser(t)

	item, err := s.FindById(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, u.ID, item.ID)
}

func TestSqlstore_Delete(t *testing.T) {
	db, teardown := repository.TestDB(t, databaseURL)
	defer teardown("author")

	s := repository.NewSqlstore(db)
	u := model.Testuser(t)

	assert.NoError(t, s.Delete(u.ID))
	assert.NotNil(t, u)
}
