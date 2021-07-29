package bod

type bodService struct {
	bodRepo BodRepository
}

func NewBodService(b BodRepository) BodService {
	return &bodService{bodRepo: b}
}

func (bs *bodService) FindAll() ([]Bod, error) {
	return bs.bodRepo.SearchAll()
	// why one linener,
	// karena tidak ada handle error yang di modif,
	// hanya menerukan saja, dari repo ke handler
	// kecuali pada saat ada error ada sebuat custom error atau logic tambahan
}

func (bs *bodService) FindById(Id int) (res Bod, err error) {
	res, err = bs.bodRepo.SearchById(Id)
	return
	// why one linener,
	// same as above
}
