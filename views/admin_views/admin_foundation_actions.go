package admin_views

import (
	"log"

	"fmt"

	"db_course/my_errors"

	"github.com/gotk3/gotk3/gtk"
	//"github.com/gotk3/gotk3/gtk"
	ents "db_course/business/entities"
	ctrls "db_course/controllers"
)

func Admin_foundation_getAll_window(UC ctrls.FoundationController) {

	// Создаём билдер
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	// Загружаем в билдер окно из файла Glade
	err = b.AddFromFile("glade/admin/admin_actions/foundation/foundation_getAll.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Получаем объект главного окна по ID
	obj, err := b.GetObject("foundation_getAll_window")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Преобразуем из объекта именно окно типа gtk.Window
	// и соединяем с сигналом "destroy" чтобы можно было закрыть
	// приложение при закрытии окна
	win := obj.(*gtk.Window)
	/*win.Connect("destroy", func() {
		gtk.MainQuit()
	})*/

	obj, err = b.GetObject("liststore1")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	store := obj.(*gtk.ListStore)
	foundations, _ := UC.GetAll()
	for _, U := range foundations {
		addRow_foundation(store, U)
	}
	// Отображаем все виджеты в окне
	win.ShowAll()

}
func addRow_foundation(listStore *gtk.ListStore, F ents.Foundation) {
	// Get an iterator for a new row at the end of the list store
	iter := listStore.Append()
	err := listStore.Set(iter, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []interface{}{F.Id, F.Login, F.Password, F.Name,
		F.CurFoudrisingAmount, F.ClosedFoundrisingAmount, fmt.Sprintf("%.2f", F.Fund_balance), fmt.Sprintf("%.2f", F.Income_history),
		fmt.Sprintf("%.2f", F.Outcome_history), F.Volunteer_amount, F.Country})

	if err != nil {
		log.Fatal("Unable to add row:", err)
		fmt.Println("Unable to add row:", err)
	}
}

func Admin_foundation_getByID_window(UC ctrls.FoundationController) {

	win, b := Get_window("glade/admin/admin_actions/foundation/foundation_getById.glade", "foundation_getByID_window")
	// Отображаем все виджеты в окне
	// Получаем поле ввода
	obj, _ := b.GetObject("getById_button")
	getById_button := obj.(*gtk.Button)
	win.ShowAll()
	obj, _ = b.GetObject("ID_entry")
	id_entry := obj.(*gtk.Entry)
	getById_button.Connect("clicked", func() {
		id, err := id_entry.GetText()
		if err == nil {
			foundation, err := UC.GetByID(id)
			if err == nil {
				Admin_foundation_Page(foundation)
			} else {
				Error_window(err.Error())
			}
		}
	})

}
func Admin_foundation_Page(foundation ents.Foundation) {
	win, b := Get_window("glade/admin/admin_actions/foundation/foundation_getAll.glade", "foundation_getAll_window")
	obj, err := b.GetObject("liststore1")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	store := obj.(*gtk.ListStore)
	if err == nil {
		addRow_foundation(store, foundation)
	} else {
		fmt.Println(err)
	}
	// Отображаем все виджеты в окне
	win.ShowAll()
}

func Admin_foundation_MultiPage(foundations []ents.Foundation) {
	win, b := Get_window("glade/admin/admin_actions/foundation/foundation_getAll.glade", "foundation_getAll_window")
	obj, err := b.GetObject("liststore1")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	store := obj.(*gtk.ListStore)
	if err == nil {
		for _, U := range foundations {
			addRow_foundation(store, U)
		}
	} else {
		fmt.Println(err)
	}
	// Отображаем все виджеты в окне
	win.ShowAll()
}

func Admin_foundation_getByName_window(UC ctrls.FoundationController) {

	win, b := Get_window("glade/admin/admin_actions/foundation/foundation_getByName.glade", "foundation_getByName_window")

	obj, _ := b.GetObject("getByName_button")
	getByLogin_button := obj.(*gtk.Button)
	win.ShowAll()
	obj, _ = b.GetObject("name_entry")
	name_entry := obj.(*gtk.Entry)
	getByLogin_button.Connect("clicked", func() {
		name, err := name_entry.GetText()
		if err == nil {
			foundation, err := UC.FS.GetByName(name)
			if err == nil {
				Admin_foundation_Page(foundation)
			} else {
				Error_window(err.Error())
			}
		}
	})
}

func Admin_foundation_getByCountry_window(UC ctrls.FoundationController) {

	win, b := Get_window("glade/admin/admin_actions/foundation/foundation_getByCountry.glade", "foundation_getByCountry_window")

	obj, _ := b.GetObject("getByCountry_button")
	getByCountry_button := obj.(*gtk.Button)
	win.ShowAll()
	obj, _ = b.GetObject("country_entry")
	country_entry := obj.(*gtk.Entry)
	getByCountry_button.Connect("clicked", func() {
		country, err := country_entry.GetText()
		if err == nil {
			foundations, err := UC.GetByCountry(country)
			if err == nil {
				Admin_foundation_MultiPage(foundations)
			} else {
				Error_window(err.Error())
			}
		}
	})
}

func Admin_foundation_create_window(UC ctrls.FoundationController) {

	win, b := Get_window("glade/admin/admin_actions/foundation/foundation_create.glade", "foundation_create_window")

	obj, _ := b.GetObject("create_button")
	create_button := obj.(*gtk.Button)
	win.ShowAll()
	obj, _ = b.GetObject("login_entry")
	login_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("password_entry")
	password_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("country_entry")
	country_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("name_entry")
	name_entry := obj.(*gtk.Entry)

	var login, password, country, name string
	var err error
	create_button.Connect("clicked", func() {
		login, err = login_entry.GetText()
		if err == nil {
			password, err = password_entry.GetText()
			if err == nil {
				country, err = country_entry.GetText()
				if err == nil {
					name, err = name_entry.GetText()
					if err == nil {
						err := UC.Add(login, password, name, country)
						if err == nil {
							Success_window()
						} else if err.Error() == my_errors.ErrCountry {
							ErrCountry_window()
						} else {
							Error_window(err.Error())
						}
					}
				}
			}
		}
	})
}
func Admin_foundation_delete_window(UC ctrls.FoundationController) {

	win, b := Get_window("glade/admin/admin_actions/foundation/foundation_delete.glade", "foundation_delete_window")
	obj, _ := b.GetObject("delete_button")
	delete_button := obj.(*gtk.Button)
	win.ShowAll()
	obj, _ = b.GetObject("ID_entry")
	id_entry := obj.(*gtk.Entry)
	delete_button.Connect("clicked", func() {
		id, err := id_entry.GetText()
		if err == nil {
			err := UC.Delete(id)
			if err == nil {
				Success_window()
			} else {
				Error_window(err.Error())
			}
		}
	})
}
func Admin_foundation_update_window(UC ctrls.FoundationController) {

	win, b := Get_window("glade/admin/admin_actions/foundation/foundation_update.glade", "foundation_update_window")

	obj, _ := b.GetObject("update_button")
	update_button := obj.(*gtk.Button)
	win.ShowAll()
	obj, _ = b.GetObject("login_entry")
	login_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("password_entry")
	password_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("country_entry")
	country_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("name_entry")
	name_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("ID_entry")
	id_entry := obj.(*gtk.Entry)

	var login, password, country, name, id string
	var err error
	update_button.Connect("clicked", func() {
		id, err = id_entry.GetText()
		if err == nil {
			login, err = login_entry.GetText()
			if err == nil {
				password, err = password_entry.GetText()
				if err == nil {
					country, err = country_entry.GetText()
					if err == nil {
						name, err = name_entry.GetText()
						if err == nil {
							err := UC.Update(id, login, password, name, country)
							if err == nil {
								Success_window()
							} else if err.Error() == my_errors.ErrCountry {
								ErrCountry_window()
							} else {
								Error_window(err.Error())
							}
						}
					}
				}
			}
		}
	})
}
