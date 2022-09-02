package admin_views

import (
	"log"

	ents "db_course/business/entities"
	ctrls "db_course/controllers"

	"fmt"

	"github.com/gotk3/gotk3/gtk"
	//"github.com/gotk3/gotk3/gtk"
)

func Admin_user_getAll_window(UC ctrls.UserController) {

	// Создаём билдер
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	// Загружаем в билдер окно из файла Glade
	err = b.AddFromFile("glade/admin/admin_actions/user/user_getAll.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Получаем объект главного окна по ID
	obj, err := b.GetObject("user_getAll_window")
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
	Users, _ := UC.GetAll()
	fmt.Println("Users are ", Users)
	for _, U := range Users {
		addRow_User(store, U)
	}
	// Отображаем все виджеты в окне
	win.ShowAll()

}
func addRow_User(listStore *gtk.ListStore, U ents.User) {
	// Get an iterator for a new row at the end of the list store
	iter := listStore.Append()
	err := listStore.Set(iter, []int{0, 1, 2, 3, 4, 5}, []interface{}{U.Id, U.Login, U.Password,
		fmt.Sprintf("%.2f", U.Balance), fmt.Sprintf("%.2f", U.CharitySum),
		U.ClosedFingAmount})

	if err != nil {
		log.Fatal("Unable to add row:", err)
		fmt.Println("Unable to add row:", err)
	}
}

func Admin_user_getByID_window(UC ctrls.UserController) {

	win, b := Get_window("glade/admin/admin_actions/user/user_getById.glade", "user_getByID_window")
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
			User, err := UC.GetByID(id)
			if err == nil {
				Admin_User_Page(User)
			} else {
				Error_window(err.Error())
			}
		}
	})

}
func Admin_User_Page(User ents.User) {
	win, b := Get_window("glade/admin/admin_actions/user/user_getAll.glade", "user_getAll_window")
	obj, err := b.GetObject("liststore1")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	store := obj.(*gtk.ListStore)
	if err == nil {
		addRow_User(store, User)
	} else {
		fmt.Println(err)
	}
	// Отображаем все виджеты в окне
	win.ShowAll()
}

func Admin_user_getByLogin_window(UC ctrls.UserController) {

	win, b := Get_window("glade/admin/admin_actions/user/user_getByLogin.glade", "user_getByLogin_window")

	obj, _ := b.GetObject("getByLogin_button")
	getByLogin_button := obj.(*gtk.Button)
	win.ShowAll()
	obj, _ = b.GetObject("Login_entry")
	login_entry := obj.(*gtk.Entry)
	getByLogin_button.Connect("clicked", func() {
		login, err := login_entry.GetText()
		if err == nil {
			User, err := UC.GetByLogin(login)
			if err == nil {
				Admin_User_Page(User)
			} else {
				Error_window(err.Error())
			}
		}
	})
}
func Admin_user_create_window(UC ctrls.UserController) {

	win, b := Get_window("glade/admin/admin_actions/user/user_create.glade", "user_create_window")

	obj, _ := b.GetObject("create_button")
	create_button := obj.(*gtk.Button)
	win.ShowAll()
	obj, _ = b.GetObject("login_entry")
	login_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("password_entry")
	password_entry := obj.(*gtk.Entry)
	var login, password string
	var err error
	create_button.Connect("clicked", func() {
		login, err = login_entry.GetText()
		if err == nil {
			password, err = password_entry.GetText()
			if err == nil {
				err := UC.Add(login, password)
				if err == nil {
					Success_window()
				} else {
					Error_window(err.Error())
				}
			}
		}
	})
}
func Admin_user_delete_window(UC ctrls.UserController) {

	win, b := Get_window("glade/admin/admin_actions/user/user_delete.glade", "user_delete_window")
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
func Admin_user_update_window(UC ctrls.UserController) {

	win, b := Get_window("glade/admin/admin_actions/user/user_update.glade", "user_update_window")

	obj, _ := b.GetObject("update_button")
	create_button := obj.(*gtk.Button)
	win.ShowAll()

	obj, _ = b.GetObject("ID_entry")
	id_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("login_entry")
	login_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("password_entry")
	password_entry := obj.(*gtk.Entry)
	var login, password, id string
	var err error
	create_button.Connect("clicked", func() {
		id, err = id_entry.GetText()
		if err == nil {
			login, err = login_entry.GetText()
			if err == nil {
				password, err = password_entry.GetText()
				if err == nil {
					err := UC.Update(id, login, password)
					if err == nil {
						Success_window()
					} else {
						Error_window(err.Error())
					}
				}
			}
		}
	})
}
