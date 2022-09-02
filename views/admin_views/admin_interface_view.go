package admin_views

import (
	"fmt"
	"log"

	servs "db_course/business/services"
	ctrls "db_course/controllers"
	repos "db_course/dataAccess/repositories"

	"github.com/gotk3/gotk3/gtk"
	"gorm.io/gorm"
)

func init_controllers(db *gorm.DB) (ctrls.UserController, ctrls.FoundationController, ctrls.FoundrisingController) {

	FR := repos.NewFoundationRepository(db)
	FS := servs.NewFoundationService(*FR)
	var FC ctrls.FoundationController
	FC.FS = FS

	FgR := repos.NewFoundrisingRepository(db)
	FgS := servs.NewFoundrisingService(*FgR)
	var FgC ctrls.FoundrisingController
	FgC.FS = FgS
	FgC.FndS = FS

	UR := repos.NewUserRepository(db)
	US := servs.NewUserService(*UR)
	var UC ctrls.UserController
	UC.US = US
	UC.FS = FS
	UC.FgS = FgS
	return UC, FC, FgC
}
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
func bindUserActions(b *gtk.Builder, UC ctrls.UserController) {
	obj, _ := b.GetObject("getAll_user_button")
	getAll_user_button := obj.(*gtk.Button)
	getAll_user_button.Connect("clicked", func() {
		fmt.Println("get all clicked")
		Admin_user_getAll_window(UC)
	})

	obj, _ = b.GetObject("getByID_user_button")
	getByID_user_button := obj.(*gtk.Button)
	getByID_user_button.Connect("clicked", func() {
		fmt.Println("get all clicked")
		Admin_user_getByID_window(UC)
	})
	obj, _ = b.GetObject("getByLog_user_button")
	getByLog_user_button := obj.(*gtk.Button)
	getByLog_user_button.Connect("clicked", func() {
		Admin_user_getByLogin_window(UC)
	})
	obj, _ = b.GetObject("add_user_button")
	add_user_button := obj.(*gtk.Button)
	add_user_button.Connect("clicked", func() {
		Admin_user_create_window(UC)
	})
	obj, _ = b.GetObject("delete_user_button")
	delete_user_button := obj.(*gtk.Button)
	delete_user_button.Connect("clicked", func() {
		Admin_user_delete_window(UC)
	})
	obj, _ = b.GetObject("update_user_button")
	update_user_button := obj.(*gtk.Button)
	update_user_button.Connect("clicked", func() {
		Admin_user_update_window(UC)
	})
}

func bindFoundationActions(b *gtk.Builder, UC ctrls.FoundationController) {
	obj, _ := b.GetObject("getAll_foundation_button")
	getAll_foundation_button := obj.(*gtk.Button)
	getAll_foundation_button.Connect("clicked", func() {
		fmt.Println("get all clicked")
		Admin_foundation_getAll_window(UC)
	})

	obj, _ = b.GetObject("getByID_foundation_button")
	getByID_foundation_button := obj.(*gtk.Button)
	getByID_foundation_button.Connect("clicked", func() {
		fmt.Println("get all clicked")
		Admin_foundation_getByID_window(UC)
	})
	obj, _ = b.GetObject("getByName_foundation_button")
	getByLog_user_button := obj.(*gtk.Button)
	getByLog_user_button.Connect("clicked", func() {
		Admin_foundation_getByName_window(UC)
	})
	obj, _ = b.GetObject("getByCountry_foundation_button")
	getByCountry_foundation_button := obj.(*gtk.Button)
	getByCountry_foundation_button.Connect("clicked", func() {
		Admin_foundation_getByCountry_window(UC)
	})
	obj, _ = b.GetObject("add_foundation_button")
	add_foundation_button := obj.(*gtk.Button)
	add_foundation_button.Connect("clicked", func() {
		Admin_foundation_create_window(UC)
	})
	obj, _ = b.GetObject("delete_foundation_button")
	delete_foundation_button := obj.(*gtk.Button)
	delete_foundation_button.Connect("clicked", func() {
		Admin_foundation_delete_window(UC)
	})
	obj, _ = b.GetObject("update_foundation_button")
	update_foundation_button := obj.(*gtk.Button)
	update_foundation_button.Connect("clicked", func() {
		Admin_foundation_update_window(UC)
	})
}

func bindFoundrisingActions(b *gtk.Builder, UC ctrls.FoundrisingController) {
	obj, _ := b.GetObject("getAll_foundrising_button")
	getAll_foundrising_button := obj.(*gtk.Button)
	getAll_foundrising_button.Connect("clicked", func() {
		fmt.Println("get all clicked")
		Admin_foundrising_getAll_window(UC)
	})

	obj, _ = b.GetObject("getByID_foundrising_button")
	getByID_foundrising_button := obj.(*gtk.Button)
	getByID_foundrising_button.Connect("clicked", func() {
		fmt.Println("get all clicked")
		Admin_foundrising_getByID_window(UC)
	})

	obj, _ = b.GetObject("getByFoundID_foundrising_button")
	getByFoundID_foundrising_button := obj.(*gtk.Button)
	getByFoundID_foundrising_button.Connect("clicked", func() {
		fmt.Println("get all clicked")
		Admin_foundrising_getByFoundID_window(UC)
	})
	obj, _ = b.GetObject("getByCrDate_foundrising_button")
	getByCrDate_foundrising_button := obj.(*gtk.Button)
	getByCrDate_foundrising_button.Connect("clicked", func() {
		Admin_foundrising_getByCrDate_window(UC)
	})
	obj, _ = b.GetObject("getByClDate_foundrising_button")
	getByClDate_foundrising_button := obj.(*gtk.Button)
	getByClDate_foundrising_button.Connect("clicked", func() {
		Admin_foundrising_getByClDate_window(UC)
	})
	obj, _ = b.GetObject("add_foundrising_button")
	add_foundrising_button := obj.(*gtk.Button)
	add_foundrising_button.Connect("clicked", func() {
		Admin_foundrising_create_window(UC)
	})
	obj, _ = b.GetObject("delete_foundrising_button")
	delete_foundrising_button := obj.(*gtk.Button)
	delete_foundrising_button.Connect("clicked", func() {
		Admin_foundrising_delete_window(UC)
	})
	obj, _ = b.GetObject("update_foundrising_button")
	update_foundrising_button := obj.(*gtk.Button)
	update_foundrising_button.Connect("clicked", func() {
		Admin_foundrising_update_window(UC)
	})
}

func bindTransactionActions(b *gtk.Builder, UC ctrls.TransactionController) {
	obj, _ := b.GetObject("getAll_transaction_button")
	getAll_transaction_button := obj.(*gtk.Button)
	getAll_transaction_button.Connect("clicked", func() {
		fmt.Println("get all clicked")
		Admin_transaction_getAll_window(UC)
	})

	obj, _ = b.GetObject("getByID_transaction_button")
	getByID_transaction_button := obj.(*gtk.Button)
	getByID_transaction_button.Connect("clicked", func() {
		Admin_transaction_getByID_window(UC)
	})

	obj, _ = b.GetObject("getByFromID_transaction_button")
	getFromID_transaction_button := obj.(*gtk.Button)
	getFromID_transaction_button.Connect("clicked", func() {
		Admin_foundrising_getFromID_window(UC)
	})

	obj, _ = b.GetObject("getByToID_transaction_button")
	getToID_transaction_button := obj.(*gtk.Button)
	getToID_transaction_button.Connect("clicked", func() {
		Admin_foundrising_getToID_window(UC)
	})

	obj, _ = b.GetObject("delete_transaction_button")
	delete_transaction_button := obj.(*gtk.Button)
	delete_transaction_button.Connect("clicked", func() {
		Admin_transaction_delete_window(UC)
	})

}

func Admin_interface_window(db *gorm.DB) {

	// Создаём билдер
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	// Загружаем в билдер окно из файла Glade
	err = b.AddFromFile("glade/admin/admin_interface.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Получаем объект главного окна по ID
	obj, err := b.GetObject("admin_window")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Преобразуем из объекта именно окно типа gtk.Window
	// и соединяем с сигналом "destroy" чтобы можно было закрыть
	// приложение при закрытии окна
	win := obj.(*gtk.Window)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	// Отображаем все виджеты в окне
	win.ShowAll()
	UC, FC, FgC := init_controllers(db)
	//ТУТ КРИНЖ, опасно ставить уже используемые services ему в агрументы!!!
	TC := ctrls.NewTransactionController(db, FC.FS, FgC.FS, UC.US)
	UC.TS = TC.TS
	bindUserActions(b, UC)
	bindFoundationActions(b, FC)
	bindFoundrisingActions(b, FgC)
	bindTransactionActions(b, TC)

}
