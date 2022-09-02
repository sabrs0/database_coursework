package services

type IService interface {
	Delete(id_ string) error
	GetById(id_ string) (interface{}, error)
}
