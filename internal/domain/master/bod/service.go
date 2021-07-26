package bod

type bodService struct {
	bodRepo BodRepository
}

func NewBodService(b BodRepository) BodService {
	return &bodService{bodRepo: b}
}

func (bs *bodService) FindAll() ([]Bod, error) {
	res, err := bs.bodRepo.SearchAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (bs *bodService) FindById(Id int) (res Bod, err error) {
	res, err = bs.bodRepo.SearchById(Id)
	if err != nil {
		return
	}
	return
}
