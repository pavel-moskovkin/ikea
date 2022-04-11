package storage

import (
	"context"
	"testing"
	"time"

	"ikea/models"
	"ikea/testutils/postgres"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	defer postgres.CleanupAndRecover(t, db)
	store := NewStorageFromDB(db)

	createUser(t, store)
}

func createUser(t *testing.T, db *Storage) uuid.NullUUID {
	ctx := context.Background()
	user := models.User{
		FirstName:  "FirstName",
		LastName:   "LastName",
		MiddleName: "MiddleName",
		FullName:   "FullName",
		Sex:        models.Female,
		BirthDate:  time.Date(2022, 4, 10, 23, 44, 00, 00, time.UTC),
	}
	uid, err := db.UserCreate(ctx, user)
	require.NoError(t, err)
	// require.NotEmpty(t, uid) TODO fix
	return uid
}
