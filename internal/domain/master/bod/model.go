package bod

type Bod struct {
	BodId       int    `json:"bod_id"`
	BodCallSign string `json:"bod_call_sign"`
	BodNamaId   string `json:"bod_nama_id"`
	BodNamaEn   string `json:"bod_nama_en"`
}
