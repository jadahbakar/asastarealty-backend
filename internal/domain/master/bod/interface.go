package bod

type BodRepository interface {
	SearchAll() ([]Bod, error)
	SearchById(Id int) (Bod, error)
}

type BodService interface {
	FindAll() ([]Bod, error)
	FindById(Id int) (Bod, error)
}
