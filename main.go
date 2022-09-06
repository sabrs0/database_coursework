package main

import (
	"db_course/views"

	conf_pk "db_course/config_pkg"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := conf_pk.Connect_string
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err == nil {
		views.Gtk_init(db)
	}

}
