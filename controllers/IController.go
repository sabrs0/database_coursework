package controllers

type IController interface {
	Delete()
	GetAll()
	GetByID()
}
