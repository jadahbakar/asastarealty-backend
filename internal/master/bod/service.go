package bod

import (
	"log"
)

type bodService struct {
	bodRepo BodRepository
}

func NewBodService(b BodRepository) BodService {
	return &bodService{bodRepo: b}
}

func (bs *bodService) FindAll() ([]Bod, error) {
	// return nil, nil
	res, err := bs.bodRepo.SearchAll()
	log.Println("Entering BOD Service FindAll Error....")
	log.Printf("rows Service -> %v", res)
	log.Printf("err Service  -> %v", err)
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
