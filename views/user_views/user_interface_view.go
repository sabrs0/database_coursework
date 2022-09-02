package user_views

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

type UserActor struct {
	User ents.User
	UC   ctrls.UserController
	FC   ctrls.FoundationController
	FgC  ctrls.FoundrisingController
	win  *gtk.Window
	b    *gtk.Builder
}

func NewUserActor(db *gorm.DB) UserActor {
	uc, fc, fgc := init_controllers(db)
	return UserActor{UC: uc, FC: fc, FgC: fgc}
}

func (UA *UserActor) User_auth(login string, password string) error {
	if UA.UC.US.ExistsByLogin(login) {
		U, err := UA.UC.GetByLogin(login)
		if err == nil {
			if U.Password == password {
				UA.User = U
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

	TR := repos.NewTransactionRepository(db)
	TS := servs.NewTransactionService(*TR)
	UR := repos.NewUserRepository(db)
	US := servs.NewUserService(*UR)
	var UC ctrls.UserController
	UC.US = US
	UC.FS = FS
	UC.FgS = FgS
	UC.TS = TS
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
func (UA *UserActor) bindUserActions() {
	obj, _ := UA.b.GetObject("getAllFoundations_button")
	getAllFoundations_button := obj.(*gtk.Button)
	getAllFoundations_button.Connect("clicked", func() {
		UA.User_getAllFoundations_window()
	})

	obj, _ = UA.b.GetObject("getFoundationByID_button")
	getFoundationByID_button := obj.(*gtk.Button)
	getFoundationByID_button.Connect("clicked", func() {
		UA.User_getFoundationByID_window()
	})
	obj, _ = UA.b.GetObject("getAllFoundrisings_button")
	getAllFoundrisings_button := obj.(*gtk.Button)
	getAllFoundrisings_button.Connect("clicked", func() {
		UA.User_getAllFoundrisings_window()
	})
	obj, _ = UA.b.GetObject("changeLogin_button")
	changeLogin_button := obj.(*gtk.Button)
	changeLogin_button.Connect("clicked", func() {
		User_changeLogin_window()
	})
	obj, _ = UA.b.GetObject("changePassword_button")
	changePassword_button := obj.(*gtk.Button)
	changePassword_button.Connect("clicked", func() {
		User_changePassword_window()
	})
	obj, _ = UA.b.GetObject("fillBalance_button")
	fillBalance_button := obj.(*gtk.Button)
	fillBalance_button.Connect("clicked", func() {
		User_fillBalance_window()
	})

	obj, _ = UA.b.GetObject("donate_button")
	donate_button := obj.(*gtk.Button)
	donate_button.Connect("clicked", func() {
		User_donate_window()
	})

}

func (UA *UserActor) init_labels() {

	obj, _ := UA.b.GetObject("ID_label")
	ID_label := obj.(*gtk.Label)
	str_id := "ID :" + strconv.FormatUint(UA.User.Id, 10)
	ID_label.SetText(str_id)

	obj, _ = UA.b.GetObject("closedFingAmount_label")
	closedFingAmount_label := obj.(*gtk.Label)
	closedFingAmount_label.SetText("Кол-во закрытых сборов :" + strconv.FormatUint(UA.User.ClosedFingAmount, 10))

	obj, _ = UA.b.GetObject("login_label")
	login_label := obj.(*gtk.Label)
	login_label.SetText("Логин :" + UA.User.Login)

	obj, _ = UA.b.GetObject("balance_label")
	balance_label := obj.(*gtk.Label)
	balance_label.SetText("Баланс :" + strconv.FormatFloat(UA.User.Balance, 'f', 2, 64))

	obj, _ = UA.b.GetObject("charitySum_label")
	charitySum_label := obj.(*gtk.Label)
	charitySum_label.SetText("Всего пожертвовано :" + strconv.FormatFloat(UA.User.CharitySum, 'f', 2, 64))
}
func (UA *UserActor) Update_User_interface_window() {
	fmt.Println("UPDATING USER PAGE WITH VALUES ", UA.User)
	UA.init_labels()
	//не знаю, нужно ли это или нет
	UA.win.ShowAll()
}
func (UA *UserActor) User_interface_window() {

	// Создаём билдер
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	// Загружаем в билдер окно из файла Glade
	err = b.AddFromFile("glade/user/user_role_interface.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Получаем объект главного окна по ID
	obj, err := b.GetObject("user_interface_window")
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
	UA.bindUserActions()

}
