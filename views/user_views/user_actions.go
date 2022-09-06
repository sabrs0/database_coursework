package user_views

import (
	"log"
	"strconv"

	"fmt"

	"github.com/gotk3/gotk3/gtk"
	//"github.com/gotk3/gotk3/gtk"
	ents "db_course/business/entities"
	av "db_course/views/admin_views"
	"strings"
)

func (UA *UserActor) User_getAllFoundations_window() {

	// Создаём билдер
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	// Загружаем в билдер окно из файла Glade
	err = b.AddFromFile("glade/user/user_actions/getAllFoundations.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Получаем объект главного окна по ID
	obj, err := b.GetObject("getAllFoundations_window")
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
	foundations, _ := UA.FC.GetAll()
	for _, U := range foundations {
		addRow_foundation(store, U)
	}
	// Отображаем все виджеты в окне
	win.ShowAll()

}
func addRow_foundation(listStore *gtk.ListStore, F ents.Foundation) {
	// Get an iterator for a new row at the end of the list store
	iter := listStore.Append()
	err := listStore.Set(iter, []int{0, 1, 2, 3, 4}, []interface{}{F.Id, F.Name,
		F.CurFoudrisingAmount, F.Volunteer_amount, F.Country})

	if err != nil {
		log.Fatal("Unable to add row:", err)
		fmt.Println("Unable to add row:", err)
	}
}
func (UA *UserActor) User_getFoundationByID_window() {
	win, b := Get_window("glade/user/user_actions/getFoundationByID.glade", "getFoundationByID1_window")
	obj, _ := b.GetObject("getById_button")
	getById_button := obj.(*gtk.Button)
	win.ShowAll()
	obj, _ = b.GetObject("ID_entry")
	id_entry := obj.(*gtk.Entry)
	getById_button.Connect("clicked", func() {
		id, err := id_entry.GetText()
		if err == nil {
			foundation, err := UA.FC.GetByID(id)
			if err == nil {
				UA.User_foundation_Page(foundation)
			} else {
				av.Error_window(err.Error())
			}
		}
	})
}
func (UA *UserActor) User_foundation_Page(foundation ents.Foundation) {
	win, b := Get_window("glade/user/user_actions/getFoundationByID2.glade", "getFoundationByID2_window")

	obj, _ := b.GetObject("ID_label")
	ID_label := obj.(*gtk.Label)
	str_id := strconv.FormatUint(foundation.Id, 10)
	ID_label.SetText("ID :" + str_id)

	obj, _ = b.GetObject("curFingAmount_label")
	curFingAmount_label := obj.(*gtk.Label)
	str_curFingAmount := strconv.FormatUint(uint64(foundation.CurFoudrisingAmount), 10)
	curFingAmount_label.SetText("Кол-во действующих сборов сборов :" + str_curFingAmount)

	obj, _ = b.GetObject("vAmount_label")
	vAmount_label := obj.(*gtk.Label)
	str_vAmount := strconv.FormatUint(uint64(foundation.Volunteer_amount), 10)
	vAmount_label.SetText("Кол-во волонтеров :" + str_vAmount)

	obj, _ = b.GetObject("country_label")
	country_label := obj.(*gtk.Label)
	country_label.SetText("Страна :" + foundation.Country)
	win.ShowAll()

	obj, _ = b.GetObject("name_label")
	name_label := obj.(*gtk.Label)
	name_label.SetText("Название :" + foundation.Name)

	obj, _ = b.GetObject("donate_button")
	donate_button := obj.(*gtk.Button)
	donate_button.Connect("clicked", func() {
		UA.User_donate_lite_window(str_id, ents.TO_FOUNDATION)
		win.Close()
	})
}

func (UA *UserActor) User_donate_lite_window(str_id string, to_type bool) {

	win, b := Get_window("glade/user/user_actions/donate_lite.glade", "donate_lite_window")
	win.ShowAll()
	obj, _ := b.GetObject("sum_entry")
	sum_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("comm_entry")
	comm_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("donate_button")
	donate_button := obj.(*gtk.Button)
	var sum, comm string
	var err error
	donate_button.Connect("clicked", func() {
		sum, err = sum_entry.GetText()
		if err == nil {
			comm, err = comm_entry.GetText()
			if err == nil {
				if to_type == ents.TO_FOUNDATION {
					err = UA.UC.DonateToFoundation(sum, comm, str_id, &UA.User)
				} else {
					err = UA.UC.DonateToFoundrising(sum, comm, str_id, &UA.User)
				}
				if err == nil {
					av.Success_window()
					UA.Update_User_interface_window()
					win.Close()
				} else {
					av.Error_window(err.Error())
				}
			} else {
				av.Error_window(err.Error())
			}

		} else {
			av.Error_window(err.Error())
		}
	})
}

func (UA *UserActor) User_getAllFoundrisings_window() {
	win, b := Get_window("glade/user/user_actions/getAllFoundrisings1.glade", "getAllFoundrisings1_window")
	// Отображаем все виджеты в окне
	// Получаем поле ввода
	obj, _ := b.GetObject("getAllFoundrisings1_button")
	getAllFoundrisings1_button := obj.(*gtk.Button)
	win.ShowAll()

	obj, _ = b.GetObject("ID_entry")
	id_entry := obj.(*gtk.Entry)

	getAllFoundrisings1_button.Connect("clicked", func() {
		id, err := id_entry.GetText()
		if err == nil {
			foundrisings, err := UA.FgC.GetByFoundIdActive(id)
			if err == nil {
				UA.User_foundrising_Pages(foundrisings)
			} else {
				av.Error_window(err.Error())
			}
		}
	})

}
func (UA *UserActor) SetFoundrisingLabels(foundrising ents.Foundrising, b *gtk.Builder) {
	obj, _ := b.GetObject("ID_label")
	ID_label := obj.(*gtk.Label)
	str_id := strconv.FormatUint(foundrising.Id, 10)
	ID_label.SetText("ID :" + str_id)

	obj, _ = b.GetObject("currentSum_label")
	currentSum_label := obj.(*gtk.Label)
	str_currentSum := strconv.FormatFloat(foundrising.Current_sum, 'f', 2, 64)
	currentSum_label.SetText("Собрано :" + str_currentSum)

	obj, _ = b.GetObject("foundID_label")
	foundID_label := obj.(*gtk.Label)
	str_reqSum := strconv.FormatFloat(foundrising.Required_sum, 'f', 2, 64)
	foundID_label.SetText("Необходимая сумма :" + str_reqSum)

	obj, _ = b.GetObject("descr_view")
	descr_view := obj.(*gtk.TextView)
	av.SetTextFromTextView(descr_view, foundrising.Description)

	obj, _ = b.GetObject("creationDate_label")
	creationDate_label := obj.(*gtk.Label)
	creationDate_label.SetText("Дата создания :" + foundrising.Creation_date[:strings.Index(foundrising.Creation_date, "T")])

	obj, _ = b.GetObject("closingDate_label")
	closingDate_label := obj.(*gtk.Label)
	var str_closing string
	if !foundrising.Closing_date.Valid {
		str_closing = "Сбор еще не закрыт"
	} else {
		str_closing = foundrising.Closing_date.String[:strings.Index(foundrising.Closing_date.String, "T")]
	}
	closingDate_label.SetText("Дата закрытия :" + str_closing)

	obj, _ = b.GetObject("philantropsAmount_label")
	philantropsAmount_label := obj.(*gtk.Label)
	str_phAmount := strconv.FormatUint(foundrising.Philantrops_amount, 10)
	philantropsAmount_label.SetText("Кол-во филантропов :" + str_phAmount)

}
func (UA *UserActor) User_foundrising_Pages(foundrisings []ents.Foundrising) {
	win, b := Get_window("glade/user/user_actions/getAllFoundrisings2.glade", "getAllFoundrisings2_window")

	cur_foundrising_index := 0
	UA.SetFoundrisingLabels(foundrisings[cur_foundrising_index], b)
	win.ShowAll()

	obj, _ := b.GetObject("next_button")
	next_button := obj.(*gtk.Button)
	next_button.Connect("clicked", func() {
		cur_foundrising_index = (cur_foundrising_index + 1) % len(foundrisings)
		UA.SetFoundrisingLabels(foundrisings[cur_foundrising_index], b)
		win.ShowAll()
		//win.Close()
	})

	obj, _ = b.GetObject("prev_button")
	prev_button := obj.(*gtk.Button)
	prev_button.Connect("clicked", func() {
		cur_foundrising_index = (cur_foundrising_index - 1 + len(foundrisings)) % len(foundrisings)
		UA.SetFoundrisingLabels(foundrisings[cur_foundrising_index], b)
		win.ShowAll()
		//win.Close()
	})

	//С эти разберемся чутка позже
	obj, _ = b.GetObject("donate_button")
	donate_button := obj.(*gtk.Button)
	donate_button.Connect("clicked", func() {
		UA.User_donate_lite_window(strconv.FormatUint(foundrisings[cur_foundrising_index].Id, 10), ents.TO_FOUNDRISING)
		foundrisings[cur_foundrising_index], _ = UA.FgC.GetByID(strconv.FormatUint(foundrisings[cur_foundrising_index].Id, 10))
		UA.SetFoundrisingLabels(foundrisings[cur_foundrising_index], b)
		win.Close()
	})
}
func (UA *UserActor) User_changeLogin_window() {
	win, b := Get_window("glade/user/user_actions/change_login.glade", "changeLogin_window")

	win.ShowAll()

	obj, _ := b.GetObject("change_button")
	change_button := obj.(*gtk.Button)

	obj, _ = b.GetObject("login_entry")
	login_entry := obj.(*gtk.Entry)
	change_button.Connect("clicked", func() {
		login, err := login_entry.GetText()
		if err == nil {
			if len(login) == 0 {
				av.Error_window("логин не был введен")
			} else {
				err := UA.UC.Update(strconv.FormatUint(UA.User.Id, 10), login, "")
				if err == nil {
					UA.User.SetLogin(login)
					av.Success_window()
					UA.Update_User_interface_window()
					win.Close()
				} else {
					av.Error_window(err.Error())
				}
			}
		} else {
			av.Error_window(err.Error())
		}

	})
}

func (UA *UserActor) User_changePassword_window() {
	win, b := Get_window("glade/user/user_actions/change_password.glade", "changePassword_window")

	win.ShowAll()

	obj, _ := b.GetObject("changePassword_button")
	change_button := obj.(*gtk.Button)

	obj, _ = b.GetObject("password_entry")
	password_entry := obj.(*gtk.Entry)
	change_button.Connect("clicked", func() {
		password, err := password_entry.GetText()
		if err == nil {
			if len(password) == 0 {
				av.Error_window("пароль не был введен")
			} else {
				err := UA.UC.Update(strconv.FormatUint(UA.User.Id, 10), "", password)
				if err == nil {
					UA.User.SetPassword(password)
					av.Success_window()
					UA.Update_User_interface_window()
					win.Close()
				} else {
					av.Error_window(err.Error())
				}
			}
		} else {
			av.Error_window(err.Error())
		}

	})
}

func (UA *UserActor) User_fillBalance_window() {

	win, b := Get_window("glade/user/user_actions/fill_balance.glade", "fillBalance_window")

	win.ShowAll()

	obj, _ := b.GetObject("fillBalance_button")
	fillBalance_button := obj.(*gtk.Button)

	obj, _ = b.GetObject("sum_entry")
	sum_entry := obj.(*gtk.Entry)
	fillBalance_button.Connect("clicked", func() {
		sum, err := sum_entry.GetText()
		if err == nil {
			if len(sum) == 0 {
				av.Error_window("сумма не была введена")
			} else {
				err := UA.UC.ReplenishBalance(sum, &UA.User)
				if err == nil {
					av.Success_window()
					UA.Update_User_interface_window()
					win.Close()
				} else {
					av.Error_window(err.Error())
				}
			}
		} else {
			av.Error_window(err.Error())
		}

	})
}

func (UA *UserActor) User_donate_full_window() {
	win, b := Get_window("glade/user/user_actions/donate.glade", "donate_window")
	win.ShowAll()
	var (
		is_foundation  bool
		is_foundrising bool
	)
	obj, _ := b.GetObject("kostil_radio_button")
	kostil_radio_button := obj.(*gtk.RadioButton)
	kostil_radio_button.SetVisible(false)

	obj, _ = b.GetObject("foundation_radio_button")
	foundation_radio_button := obj.(*gtk.RadioButton)

	foundation_radio_button.Connect("toggled", func() {
		is_foundation = true
		is_foundrising = false
	})

	obj, _ = b.GetObject("foundrising_radio_button")
	foundrising_radio_button := obj.(*gtk.RadioButton)

	foundrising_radio_button.Connect("toggled", func() {
		is_foundation = false
		is_foundrising = true
	})

	obj, _ = b.GetObject("ID_entry")
	ID_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("sum_entry")
	sum_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("сomm_entry")
	сomm_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("donate_button")
	donate_button := obj.(*gtk.Button)
	var sum, comm, str_id string
	var err error
	donate_button.Connect("clicked", func() {
		str_id, err = ID_entry.GetText()
		if err == nil {
			sum, err = sum_entry.GetText()
			if err == nil {
				comm, err = сomm_entry.GetText()
				if err == nil {
					if !is_foundation && !is_foundrising {
						av.Error_window("необходимо выбрать получателя")
					} else {
						if is_foundation {
							err = UA.UC.DonateToFoundation(sum, comm, str_id, &UA.User)
						} else if is_foundrising {
							err = UA.UC.DonateToFoundrising(sum, comm, str_id, &UA.User)
						}
						if err == nil {
							av.Success_window()
							UA.Update_User_interface_window()
							win.Close()
						} else {
							av.Error_window(err.Error())
						}
					}
				} else {
					av.Error_window(err.Error())
				}

			} else {
				av.Error_window(err.Error())
			}
		} else {
			av.Error_window(err.Error())
		}
	})
}
