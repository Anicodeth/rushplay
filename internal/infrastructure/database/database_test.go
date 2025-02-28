package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestDatabaseConnection(t *testing.T) {
	db, err := NewDatabase()

	assert.NoError(t, err, "Error connecting to database")

	assert.NotNil(t, db, "Database should not be nil")
}

func TestMigrate(t *testing.T) {
	db, _ := NewDatabase()

	err := Migrate(db)

	assert.NoError(t, err, "Error migrating database")
}