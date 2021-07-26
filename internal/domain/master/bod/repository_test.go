package bod_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jadahbakar/asastarealty-backend/internal/master/bod"
	"github.com/stretchr/testify/assert"
)

var u = &bod.Bod{
	BodId:       1,
	BodCallSign: "Call Sign 1",
	BodNamaId:   "Nama 1",
	BodNamaEn:   "Name 1",
}

func TestSearchAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	rows := sqlmock.NewRows([]string{"bod_id", "bod_call_sign", "bod_nama_id", "bod_nama_en"}).
		AddRow(1, "Call Sign 1", "Nama 1", "Name 1").
		AddRow(2, "Call Sign 2", "Nama 2", "Name 2")
	// mockBods := []bod.Bod{
	// 	{BodId: 1, BodCallSign: "Call Sign 1", BodNamaId: "Nama 1", BodNamaEn: "Name 1"},
	// 	{BodId: 2, BodCallSign: "Call Sign 2", BodNamaId: "Nama 2", BodNamaEn: "Name 2"},
	// }
	// rows := sqlmock.NewRows([]string{"bod_id", "bod_call_sign", "bod_nama_id", "bod_nama_en"}).
	// 	AddRow(mockBods[0].BodId, mockBods[0].BodCallSign, mockBods[0].BodNamaId, mockBods[0].BodNamaEn).
	// 	AddRow(mockBods[1].BodId, mockBods[1].BodCallSign, mockBods[1].BodNamaId, mockBods[1].BodNamaEn)
	query := `SELECT bod_id, bod_call_sign, bod_nama_id, bod_nama_en FROM mst\.bod ORDER BY bod_id`
	mock.ExpectQuery(query).WillReturnRows(rows)
	a := bod.NewBodRepository(db)
	list, err := a.SearchAll()
	assert.NotEmpty(t, list)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestSearchById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	rows := sqlmock.NewRows([]string{"bod_id", "bod_call_sign", "bod_nama_id", "bod_nama_en"}).
		AddRow(1, "Call Sign 1", "Nama 1", "Name 1")
	query := `SELECT bod_id, bod_call_sign, bod_nama_id, bod_nama_en FROM mst\.bod WHERE BY bod_id\=\?`
	BodId := int(1)
	mock.ExpectQuery(query).WillReturnRows(rows)
	a := bod.NewBodRepository(db)
	aBod, err := a.SearchById(BodId)
	assert.NoError(t, err)
	assert.NotNil(t, aBod)
}
