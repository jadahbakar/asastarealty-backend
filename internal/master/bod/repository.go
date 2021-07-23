package bod

import (
	"database/sql"
	"log"
)

type bodRepository struct {
	Conn *sql.DB
}

func NewBodRepository(Conn *sql.DB) BodRepository {
	return &bodRepository{Conn: Conn}
}

func (br *bodRepository) SearchAll() ([]Bod, error) {
	query := `SELECT bod_id, bod_call_sign, bod_nama_id, bod_nama_en
				FROM mst.bod ORDER BY bod_id`
	rows, err := br.Conn.Query(query)
	log.Println("Entering BOD Repo SearchAll Error....")
	log.Printf("rows -> %v", rows)
	log.Printf("err  -> %v", err)
	if err != nil {
		return nil, err
	}

	result := make([]Bod, 0)
	t := Bod{}
	for rows.Next() {
		err = rows.Scan(
			&t.BodId,
			&t.BodCallSign,
			&t.BodNamaId,
			&t.BodNamaEn,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, t)
	}
	if rows.Err() != nil {
		// if any error occurred while reading rows.
		log.Println("Error will reading user table: \n", err)
		return nil, rows.Err()
	}
	return result, nil
}

func (br *bodRepository) SearchById(Id int) (Bod, error) {
	var t Bod
	query := `SELECT bod_id, bod_call_sign, bod_nama_id, bod_nama_en
				FROM mst.bod WHERE BY bod_id=?`
	err := br.Conn.QueryRow(query, Id).Scan(
		&t.BodId,
		&t.BodCallSign,
		&t.BodNamaId,
		&t.BodNamaEn)
	if err != nil {
		log.Println("Error will reading user table: \n", err)
	}
	return t, nil
}
