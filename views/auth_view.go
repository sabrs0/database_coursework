package views

import (
	"log"

	"db_course/views/admin_views"

	"db_course/views/user_views"

	"db_course/views/foundation_views"

	"github.com/gotk3/gotk3/gtk"
	"gorm.io/gorm"
)

const (
	admin_login    = "admin"
	admin_password = "admin"
)

func Get_window(filename string, winName string) (*gtk.Window, *gtk.Builder) {
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	// Загружаем в билдер окно из файла Glade
	err = b.AddFromFile(filename)
	if err != nil {
		log.Fatal("Ошибка:", err)
		return nil, nil
	}

	// Получаем объект главного окна по ID
	obj, err := b.GetObject(winName)
	if err != nil {
		log.Fatal("Ошибка:", err)
		return nil, nil
	}

	// Преобразуем из объекта именно окно типа gtk.Window
	// и соединяем с сигналом "destroy" чтобы можно было закрыть
	// приложение при закрытии окна
	win := obj.(*gtk.Window)
	/*win.Connect("destroy", func() {
		gtk.MainQuit()
	})*/
	return win, b
}
func Gtk_init(db *gorm.DB) {
	// Инициализируем GTK.
	gtk.Init(nil)

	Auth_window(db)

	// Выполняем главный цикл GTK (для отрисовки). Он остановится когда
	// выполнится gtk.MainQuit()
	gtk.Main()
}
func Auth_window(db *gorm.DB) {

	win, b := Get_window("glade/auth.glade", "window")
	// Отображаем все виджеты в окне
	// Получаем поле ввода
	obj, _ := b.GetObject("login entry")
	login_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("password entry")
	password_entry := obj.(*gtk.Entry)

	var (
		is_admin bool
		is_user  bool
		is_found bool
	)
	/*obj, _ = b.GetObject("radio_found")
	radio_found_button := obj.(*gtk.RadioButton)

	radio_found_button.Connect("toggled", func() {
		is_found = true
	})*/

	win.ShowAll()
	obj, _ = b.GetObject("err_label")
	err_label := obj.(*gtk.Label)
	err_label.SetVisible(false)

	obj, _ = b.GetObject("kostil_button")
	radio_kostil_button := obj.(*gtk.RadioButton)
	radio_kostil_button.SetVisible(false)

	obj, _ = b.GetObject("radio_admin")
	radio_admin_button := obj.(*gtk.RadioButton)

	radio_admin_button.Connect("toggled", func() {
		is_admin = true
		is_user = false
		is_found = false
	})

	obj, _ = b.GetObject("radio_user")
	radio_user_button := obj.(*gtk.RadioButton)

	radio_user_button.Connect("toggled", func() {
		is_user = true
		is_admin = false
		is_found = false
	})

	obj, _ = b.GetObject("radio_found")
	radio_foundation_button := obj.(*gtk.RadioButton)

	radio_foundation_button.Connect("toggled", func() {
		is_user = false
		is_admin = false
		is_found = true
	})

	var login, password string
	var err error
	// Получаем кнопку
	obj, _ = b.GetObject("login_button")
	login_button := obj.(*gtk.Button)
	login_button.Connect("clicked", func() {
		login, err = login_entry.GetText()
		if err == nil {
			password, err = password_entry.GetText()
			if err == nil {
				if is_admin {
					if (login == admin_login) && (password == admin_password) {
						win.Close()
						admin_views.Admin_interface_window(db)
					} else {
						err_label.SetVisible(true)
					}
				} else if is_user {
					UA := user_views.NewUserActor(db)
					err := UA.User_auth(login, password)
					if err == nil {
						win.Close()
						UA.User_interface_window()
					} else {
						err_label.SetVisible(true)
					}
				} else if is_found {
					FA := foundation_views.NewFoundationActor(db)
					err := FA.Foundation_auth(login, password)
					if err == nil {
						win.Close()
						FA.Foundation_interface_window()
					} else {
						err_label.SetVisible(true)
					}
				}
			}
		}
	})

}
