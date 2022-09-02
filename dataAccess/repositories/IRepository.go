package repositories

type IRepository interface {
	Delete(Entity interface{}) error
	SelectById(id uint64) (interface{}, bool)
}
