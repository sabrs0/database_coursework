package foundation_views

import (
	"strconv"
	"strings"

	"github.com/gotk3/gotk3/gtk"
	//"github.com/gotk3/gotk3/gtk"
	ents "db_course/business/entities"
	av "db_course/views/admin_views"
)

func (UA *FoundationActor) Foundation_donate_lite_window(str_id string) {

	win, b := Get_window("glade/foundation/foundation_actions/donate_lite.glade", "donate_lite_window")
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
				err = UA.FC.DonateToFoundrising(sum, comm, str_id, &UA.Foundation) //Вернемся после рекламы
				if err == nil {
					av.Success_window()
					UA.Update_Foundation_interface_window()
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

func (UA *FoundationActor) Foundation_getAllMyFoundrisings_window() {
	var foundrisings []ents.Foundrising
	var err error
	foundrisings, err = UA.FgC.GetByFoundId(strconv.FormatUint(UA.Foundation.Id, 10))
	if err == nil {
		win, b := Get_window("glade/foundation/foundation_actions/getAllFoundrisings.glade", "getAllFoundrisings_window")

		cur_foundrising_index := 0
		UA.SetFoundrisingLabels(foundrisings[cur_foundrising_index], b)
		win.ShowAll()

		obj, _ := b.GetObject("next_button")
		next_button := obj.(*gtk.Button)
		next_button.Connect("clicked", func() {
			foundrisings, err = UA.FgC.GetByFoundId(strconv.FormatUint(UA.Foundation.Id, 10))
			if err == nil {
				cur_foundrising_index = (cur_foundrising_index + 1) % len(foundrisings)
				UA.SetFoundrisingLabels(foundrisings[cur_foundrising_index], b)
				win.ShowAll()
				//win.Close()
			} else {
				av.Error_window(err.Error())
			}
		})

		obj, _ = b.GetObject("prev_button")
		prev_button := obj.(*gtk.Button)
		prev_button.Connect("clicked", func() {
			foundrisings, err = UA.FgC.GetByFoundId(strconv.FormatUint(UA.Foundation.Id, 10))
			if err == nil {
				cur_foundrising_index = (cur_foundrising_index - 1 + len(foundrisings)) % len(foundrisings)
				UA.SetFoundrisingLabels(foundrisings[cur_foundrising_index], b)
				win.ShowAll()
			} else {
				av.Error_window(err.Error())
			}
			//win.Close()
		})

		//С эти разберемся чутка позже
		obj, _ = b.GetObject("donate_button")
		donate_button := obj.(*gtk.Button)
		donate_button.Connect("clicked", func() {
			UA.Foundation_donate_lite_window(strconv.FormatUint(foundrisings[cur_foundrising_index].Id, 10))
			foundrisings[cur_foundrising_index], _ = UA.FgC.GetByID(strconv.FormatUint(foundrisings[cur_foundrising_index].Id, 10))
			UA.SetFoundrisingLabels(foundrisings[cur_foundrising_index], b)
			win.Close()
		})
	} else {
		av.Error_window(err.Error())
	}
}

func (UA *FoundationActor) Foundation_getMyFoundrisingByID_window() {
	win, b := Get_window("glade/foundation/foundation_actions/getFoundrisingByID.glade", "getMyFoundrisingByID_window")
	obj, _ := b.GetObject("getById_button")
	getById_button := obj.(*gtk.Button)
	win.ShowAll()
	obj, _ = b.GetObject("ID_entry")
	id_entry := obj.(*gtk.Entry)
	getById_button.Connect("clicked", func() {
		id, err := id_entry.GetText()
		if err == nil {
			foundrising, err := UA.FgC.GetByIdAndFoundId(id, strconv.FormatUint(UA.Foundation.Id, 10))
			if err == nil {
				UA.Foundation_foundrising_Page(foundrising)
			} else {
				av.Error_window(err.Error())
			}
		}
	})
}

func (UA *FoundationActor) SetFoundrisingLabels(foundrising ents.Foundrising, b *gtk.Builder) {
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
func (UA *FoundationActor) Foundation_foundrising_Page(foundrising ents.Foundrising) {
	win, b := Get_window("glade/foundation/foundation_actions/getFoundrisingByID2.glade", "getFoundrisingById2_window")

	UA.SetFoundrisingLabels(foundrising, b)
	win.ShowAll()
	obj, _ := b.GetObject("donate_button")
	donate_button := obj.(*gtk.Button)
	donate_button.Connect("clicked", func() {
		UA.Foundation_donate_lite_window(strconv.FormatUint(foundrising.Id, 10))
		win.Close()
	})
}

func (UA *FoundationActor) Foundation_createFoundrising_window() {
	win, b := Get_window("glade/foundation/foundation_actions/create_foundrising.glade", "create_foundrising_window")

	obj, _ := b.GetObject("create_button")
	create_button := obj.(*gtk.Button)
	win.ShowAll()

	obj, _ = b.GetObject("sum_entry")
	sum_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("descr_entry")
	descr_entry := obj.(*gtk.TextView)

	var sum, descr string
	var err error
	create_button.Connect("clicked", func() {

		sum, err = sum_entry.GetText()
		if err == nil {
			descr, err = av.GetTextFromTextView(descr_entry)
			if err == nil {
				err := UA.FgC.Add(strconv.FormatUint(UA.Foundation.Id, 10), descr, sum)
				if err == nil {
					UA.Foundation.CurFoudrisingAmount += 1
					UA.FC.FS.FR.Update(UA.Foundation)
					av.Success_window()
					UA.Update_Foundation_interface_window()
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

func (UA *FoundationActor) Foundation_changeLogin_window() {
	win, b := Get_window("glade/foundation/foundation_actions/change_login.glade", "changeLogin_window")

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
				err := UA.FC.Update(strconv.FormatUint(UA.Foundation.Id, 10), login, "", "", "")
				if err == nil {
					UA.Foundation.SetLogin(login)
					av.Success_window()
					UA.Update_Foundation_interface_window()
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

func (UA *FoundationActor) Foundation_changePassword_window() {
	win, b := Get_window("glade/foundation/foundation_actions/change_password.glade", "changePassword_window")

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
				err := UA.FC.Update(strconv.FormatUint(UA.Foundation.Id, 10), "", password, "", "")
				if err == nil {
					UA.Foundation.SetPassword(password)
					av.Success_window()
					UA.Update_Foundation_interface_window()
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

func (UA *FoundationActor) Foundation_fillBalance_window() {

	win, b := Get_window("glade/foundation/foundation_actions/fill_balance.glade", "fillBalance_window")

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
				err := UA.FC.ReplenishBalance(sum, &UA.Foundation)
				if err == nil {
					av.Success_window()
					UA.Update_Foundation_interface_window()
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

func (UA *FoundationActor) Foundation_donate_full_window() {
	win, b := Get_window("glade/foundation/foundation_actions/donate.glade", "donate_window")
	win.ShowAll()

	obj, _ := b.GetObject("ID_entry")
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
					err = UA.FC.DonateToFoundrising(sum, comm, str_id, &UA.Foundation)
					if err == nil {
						av.Success_window()
						UA.Update_Foundation_interface_window()
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
	})
}
