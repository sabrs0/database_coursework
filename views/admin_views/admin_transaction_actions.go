package admin_views

import (
	"log"

	"fmt"

	"github.com/gotk3/gotk3/gtk"
	//"github.com/gotk3/gotk3/gtk"
	ents "db_course/business/entities"
	ctrls "db_course/controllers"
)

func Admin_transaction_Page(transaction ents.Transaction) {
	win, b := Get_window("glade/admin/admin_actions/transaction/transaction_getAll.glade", "transaction_getAll_window")
	obj, err := b.GetObject("liststore1")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	store := obj.(*gtk.ListStore)
	if err == nil {
		addRow_transaction(store, transaction)
	} else {
		fmt.Println(err)
	}
	// Отображаем все виджеты в окне
	win.ShowAll()
}

func Admin_transaction_MultiPage(transactions []ents.Transaction) {
	win, b := Get_window("glade/admin/admin_actions/transaction/transaction_getAll.glade", "transaction_getAll_window")
	obj, err := b.GetObject("liststore1")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	store := obj.(*gtk.ListStore)
	if err == nil {
		for _, U := range transactions {
			addRow_transaction(store, U)
		}
	} else {
		fmt.Println(err)
	}
	// Отображаем все виджеты в окне
	win.ShowAll()
}
func Admin_transaction_getAll_window(UC ctrls.TransactionController) {

	// Создаём билдер
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	// Загружаем в билдер окно из файла Glade
	err = b.AddFromFile("glade/admin/admin_actions/transaction/transaction_getAll.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Получаем объект главного окна по ID
	obj, err := b.GetObject("transaction_getAll_window")
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
	transactions, err := UC.GetAll()
	if err == nil {
		for _, U := range transactions {
			addRow_transaction(store, U)
		}
	} else {
		Error_window(err.Error())
	}
	// Отображаем все виджеты в окне
	win.ShowAll()

}
func addRow_transaction(listStore *gtk.ListStore, T ents.Transaction) {
	// Get an iterator for a new row at the end of the list store
	iter := listStore.Append()
	var from_essence string
	var to_essence string
	if T.From_essence_type == ents.FROM_USER {
		from_essence = "пользователь"
	} else {
		from_essence = "фонд"
	}
	if T.To_essence_type == ents.TO_FOUNDATION {
		to_essence = "фонд"
	} else {
		to_essence = "сбор"
	}
	err := listStore.Set(iter, []int{0, 1, 2, 3, 4, 5, 6}, []interface{}{T.Id, from_essence, T.From_id,
		to_essence, T.To_id, fmt.Sprintf("%.2f", T.Sum_of_money), T.Comment})
	if err != nil {
		log.Fatal("Unable to add row:", err)
		fmt.Println("Unable to add row:", err)
	}
}

func Admin_transaction_getByID_window(UC ctrls.TransactionController) {

	win, b := Get_window("glade/admin/admin_actions/transaction/transaction_getByID.glade", "transaction_getByID_window")
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
			transaction, err := UC.GetByID(id)
			if err == nil {
				Admin_transaction_Page(transaction)
			} else {
				Error_window(err.Error())
			}
		}
	})

}

func Admin_foundrising_getFromID_window(UC ctrls.TransactionController) {

	win, b := Get_window("glade/admin/admin_actions/transaction/transaction_getByFromID.glade", "transaction_getByFromID_window")
	var sender string
	/*var (
		is_user       bool
		is_foundation bool
	)*/
	win.ShowAll()
	obj, _ := b.GetObject("kostil_radio_button")
	radio_kostil_button := obj.(*gtk.RadioButton)
	radio_kostil_button.SetVisible(false)

	obj, _ = b.GetObject("user_radio_button")
	user_radio_button := obj.(*gtk.RadioButton)

	obj, _ = b.GetObject("foundation_radio_button")
	foundation_radio_button := obj.(*gtk.RadioButton)

	user_radio_button.Connect("toggled", func() {
		/*is_user = true
		is_foundation = false*/
		sender = "false"
	})

	foundation_radio_button.Connect("toggled", func() {
		/*is_user = false
		is_foundation = true*/
		sender = "true"
	})

	obj, _ = b.GetObject("getByFromId_button")
	getByFromId_button := obj.(*gtk.Button)
	obj, _ = b.GetObject("ID_entry")
	id_entry := obj.(*gtk.Entry)
	getByFromId_button.Connect("clicked", func() {
		if sender == "" {
			Error_window("необходимо выбрать отправителя")
		} else {

			id, err := id_entry.GetText()
			if err == nil {
				transaction, err := UC.GetFromId(sender, id)
				if err == nil {
					Admin_transaction_MultiPage(transaction)
				} else {
					Error_window(err.Error())
				}
			}
		}
	})
}

func Admin_foundrising_getToID_window(UC ctrls.TransactionController) {

	win, b := Get_window("glade/admin/admin_actions/transaction/transaction_getByToID.glade", "transaction_getByToID_window")
	var reciever string
	/*var (
		is_user       bool
		is_foundation bool
	)*/
	win.ShowAll()
	obj, _ := b.GetObject("kostil_radio_button")
	radio_kostil_button := obj.(*gtk.RadioButton)
	radio_kostil_button.SetVisible(false)

	obj, _ = b.GetObject("foundation_radio_button")
	foundation_radio_button := obj.(*gtk.RadioButton)

	obj, _ = b.GetObject("foundrising_radio_button")
	foundrising_radio_button := obj.(*gtk.RadioButton)

	foundation_radio_button.Connect("toggled", func() {
		/*is_user = true
		is_foundation = false*/
		reciever = "false"
	})

	foundrising_radio_button.Connect("toggled", func() {
		/*is_user = false
		is_foundation = true*/
		reciever = "true"
	})

	obj, _ = b.GetObject("getByToId_button")
	getByToId_button := obj.(*gtk.Button)

	obj, _ = b.GetObject("ID_entry")
	id_entry := obj.(*gtk.Entry)
	getByToId_button.Connect("clicked", func() {
		if reciever == "" {
			Error_window("необходимо выбрать отправителя")
		} else {

			id, err := id_entry.GetText()
			if err == nil {
				transaction, err := UC.GetToId(reciever, id)
				if err == nil {
					Admin_transaction_MultiPage(transaction)
				} else {
					Error_window(err.Error())
				}
			}
		}
	})
}

func Admin_transaction_delete_window(UC ctrls.TransactionController) {

	win, b := Get_window("glade/admin/admin_actions/transaction/transaction_delete.glade", "transaction_delete_window")
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
