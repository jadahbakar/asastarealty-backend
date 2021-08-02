package bod_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jadahbakar/asastarealty-backend/internal/domain/master/bod"
	"github.com/stretchr/testify/assert"
)

func TestSearchAll(t *testing.T) {
	// setup db mocking
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func() { db.Close() }() // must close db after done

	// setup data mocking
	rows := sqlmock.NewRows([]string{"bod_id", "bod_call_sign", "bod_nama_id", "bod_nama_en"}).
		AddRow(1, "Call Sign 1", "Nama 1", "Name 1").
		AddRow(2, "Call Sign 2", "Nama 2", "Name 2")

	query := "SELECT bod_id, bod_call_sign, bod_nama_id, bod_nama_en FROM mst.bod ORDER BY bod_id"
	mock.ExpectQuery(query).WillReturnRows(rows)

	// testing
	repo := bod.NewBodRepository(db)
	list, err := repo.SearchAll()

	// assertion testing
	assert.NoError(t, err)
	assert.NotEmpty(t, list)
	assert.Len(t, list, 2)

	// assertion mocking
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSearchById(t *testing.T) {
	// setup db mocking
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func() { db.Close() }() // must close db after done

	// setup data mocking
	rows := sqlmock.NewRows([]string{"bod_id", "bod_call_sign", "bod_nama_id", "bod_nama_en"}).
		AddRow(1, "Call Sign 1", "Nama 1", "Name 1")

	query := "SELECT bod_id, bod_call_sign, bod_nama_id, bod_nama_en FROM mst.bod WHERE bod_id = $1"
	query = regexp.QuoteMeta(query)
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	// testing
	repo := bod.NewBodRepository(db)
	aBod, err := repo.SearchById(1)

	// assertion testing
	assert.NoError(t, err)
	assert.NotNil(t, aBod)

	// assertion mocking
	// assert.NoError(t, mock.ExpectationsWereMet())
}
