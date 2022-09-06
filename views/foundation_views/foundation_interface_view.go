package foundation_views

import (
	"fmt"
	"log"
	"strconv"

	ents "db_course/business/entities"
	servs "db_course/business/services"
	ctrls "db_course/controllers"
	repos "db_course/dataAccess/repositories"

	"github.com/gotk3/gotk3/gtk"
	"gorm.io/gorm"
)

type FoundationActor struct {
	Foundation ents.Foundation
	FC         ctrls.FoundationController
	FgC        ctrls.FoundrisingController
	win        *gtk.Window
	b          *gtk.Builder
}

func NewFoundationActor(db *gorm.DB) FoundationActor {
	fc, fgc := init_controllers(db)
	return FoundationActor{FC: fc, FgC: fgc}
}

func (UA *FoundationActor) Foundation_auth(login string, password string) error {
	if UA.FC.FS.ExistsByLogin(login) {
		U, err := UA.FC.GetByLogin(login)
		if err == nil {
			if U.Password == password {
				UA.Foundation = U
				return nil
			} else {
				return fmt.Errorf("неверный пароль для %s", U.Login)
			}
		} else {
			return err
		}
	} else {
		return fmt.Errorf("пользователя с таким логином не существует")
	}
}

func init_controllers(db *gorm.DB) (ctrls.FoundationController, ctrls.FoundrisingController) {

	FR := repos.NewFoundationRepository(db)
	FS := servs.NewFoundationService(*FR)
	var FC ctrls.FoundationController
	FC.FS = FS

	FgR := repos.NewFoundrisingRepository(db)
	FgS := servs.NewFoundrisingService(*FgR)
	var FgC ctrls.FoundrisingController
	FgC.FS = FgS
	FgC.FndS = FS

	TR := repos.NewTransactionRepository(db)
	TS := servs.NewTransactionService(*TR)

	FC.TS = TS
	FC.FgS = FgS
	return FC, FgC
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
func (UA *FoundationActor) bindFoundationActions() {
	obj, _ := UA.b.GetObject("getAllMyFoundrisings_button")
	getAllMyFoundrisings_button := obj.(*gtk.Button)
	getAllMyFoundrisings_button.Connect("clicked", func() {

		UA.Foundation_getAllMyFoundrisings_window()
	})

	obj, _ = UA.b.GetObject("getMyFoundrisingByID_button")
	getMyFoundrisingByID_button := obj.(*gtk.Button)
	getMyFoundrisingByID_button.Connect("clicked", func() {
		UA.Foundation_getMyFoundrisingByID_window()
	})
	obj, _ = UA.b.GetObject("createFoundrising_button")
	createFoundrising_button := obj.(*gtk.Button)
	createFoundrising_button.Connect("clicked", func() {
		UA.Foundation_createFoundrising_window()
	})
	obj, _ = UA.b.GetObject("changeLogin_button")
	changeLogin_button := obj.(*gtk.Button)
	changeLogin_button.Connect("clicked", func() {
		UA.Foundation_changeLogin_window()
	})
	obj, _ = UA.b.GetObject("changePassword_button")
	changePassword_button := obj.(*gtk.Button)
	changePassword_button.Connect("clicked", func() {
		UA.Foundation_changePassword_window()
	})
	obj, _ = UA.b.GetObject("fillBalance_button")
	fillBalance_button := obj.(*gtk.Button)
	fillBalance_button.Connect("clicked", func() {
		UA.Foundation_fillBalance_window()
	})

	obj, _ = UA.b.GetObject("donate_main_button")
	donate_main_button := obj.(*gtk.Button)
	donate_main_button.Connect("clicked", func() {

		UA.Foundation_donate_full_window()
	})

}

func (UA *FoundationActor) init_labels() {

	obj, _ := UA.b.GetObject("ID_label")
	ID_label := obj.(*gtk.Label)
	str_id := "ID :" + strconv.FormatUint(UA.Foundation.Id, 10)
	ID_label.SetText(str_id)

	obj, _ = UA.b.GetObject("closedFingAmount_label")
	closedFingAmount_label := obj.(*gtk.Label)
	closedFingAmount_label.SetText("Кол-во закрытых сборов :" + strconv.FormatUint(uint64(UA.Foundation.ClosedFoundrisingAmount), 10))

	obj, _ = UA.b.GetObject("login_label")
	login_label := obj.(*gtk.Label)
	login_label.SetText("Логин :" + UA.Foundation.Login)

	obj, _ = UA.b.GetObject("balance_label")
	balance_label := obj.(*gtk.Label)
	balance_label.SetText("Баланс :" + strconv.FormatFloat(UA.Foundation.Fund_balance, 'f', 2, 64))

	obj, _ = UA.b.GetObject("name_label")
	name_label := obj.(*gtk.Label)
	name_label.SetText("Название :" + UA.Foundation.Name)

	obj, _ = UA.b.GetObject("curFingAmount_label")
	curFingAmount_label := obj.(*gtk.Label)
	curFingAmount_label.SetText("Кол-во активных сборов :" + strconv.FormatUint(uint64(UA.Foundation.CurFoudrisingAmount), 10))

	obj, _ = UA.b.GetObject("country_label")
	country_label := obj.(*gtk.Label)
	country_label.SetText("Страна :" + UA.Foundation.Country)

	obj, _ = UA.b.GetObject("vAmount_label")
	vAmount_label := obj.(*gtk.Label)
	vAmount_label.SetText("Кол-во волонтеров :" + strconv.FormatUint(uint64(UA.Foundation.Volunteer_amount), 10))

	obj, _ = UA.b.GetObject("income_label")
	income_label := obj.(*gtk.Label)
	income_label.SetText("Получено средств :" + strconv.FormatFloat(UA.Foundation.Income_history, 'f', 2, 64))

	obj, _ = UA.b.GetObject("outcome_label")
	outcome_label := obj.(*gtk.Label)
	outcome_label.SetText("Пожертвовано средств :" + strconv.FormatFloat(UA.Foundation.Outcome_history, 'f', 2, 64))
}
func (UA *FoundationActor) Update_Foundation_interface_window() {
	fmt.Println("UPDATING Foundation PAGE WITH VALUES ", UA.Foundation)
	UA.init_labels()
	//не знаю, нужно ли это или нет
	UA.win.ShowAll()
}
func (UA *FoundationActor) Foundation_interface_window() {

	// Создаём билдер
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	// Загружаем в билдер окно из файла Glade
	err = b.AddFromFile("glade/foundation/foundation_role_interface.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Получаем объект главного окна по ID
	obj, err := b.GetObject("foundation_interface_window")
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
	UA.b = b
	UA.win = win
	UA.init_labels()
	win.ShowAll()
	UA.bindFoundationActions()

}
