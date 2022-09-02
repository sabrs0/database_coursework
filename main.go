package main

import (
	"db_course/views"

	//ents "db_course/business/entities"
	conf_pk "db_course/config_pkg"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
func create_controller(db *gorm.DB, repo *repos.IRepository, service *servs.IService, controller *ctrls.IController) {
	repo = repos.NewUserRepository(db)
	var US *servs.UserService
	US = servs.NewUserService(*UR)
	var UC ctrls.UserController
	UC.US = US
}*/
func main() {
	dsn := conf_pk.Connect_string
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err == nil {
		views.Gtk_init(db)
	}
}
