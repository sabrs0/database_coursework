package admin_views

import (
	"log"

	"fmt"

	"github.com/gotk3/gotk3/gtk"
	//"github.com/gotk3/gotk3/gtk"
	ents "db_course/business/entities"
	ctrls "db_course/controllers"
	"strings"
)

func GetTextFromTextView(descr_entry *gtk.TextView) (string, error) {
	var buf *gtk.TextBuffer
	var err error
	buf, err = descr_entry.GetBuffer()
	var text string
	if err == nil {

		start, end := buf.GetBounds()
		text, err = buf.GetText(start, end, false)
	}
	return text, err
}
func SetTextFromTextView(descr_entry *gtk.TextView, text string) {
	descr_entry.SetEditable(true)
	//var buf *gtk.TextBuffer
	buf, err := descr_entry.GetBuffer()
	if err == nil {
		buf.SetText(text)
		//descr_entry.SetBuffer(buf)
		descr_entry.SetEditable(false)
	}
}

func Admin_foundrising_Page(foundrising ents.Foundrising) {
	win, b := Get_window("glade/admin/admin_actions/foundrising/foundrising_getAll.glade", "foundrising_getAll_window")
	obj, err := b.GetObject("liststore1")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	store := obj.(*gtk.ListStore)
	if err == nil {
		addRow_foundrising(store, foundrising)
	} else {
		fmt.Println(err)
	}
	// Отображаем все виджеты в окне
	win.ShowAll()
}

func Admin_foundrising_MultiPage(foundrisings []ents.Foundrising) {
	win, b := Get_window("glade/admin/admin_actions/foundrising/foundrising_getAll.glade", "foundrising_getAll_window")
	obj, err := b.GetObject("liststore1")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	store := obj.(*gtk.ListStore)
	if err == nil {
		for _, U := range foundrisings {
			addRow_foundrising(store, U)
		}
	} else {
		fmt.Println(err)
	}
	// Отображаем все виджеты в окне
	win.ShowAll()
}
func Admin_foundrising_getAll_window(UC ctrls.FoundrisingController) {

	// Создаём билдер
	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	// Загружаем в билдер окно из файла Glade
	err = b.AddFromFile("glade/admin/admin_actions/foundrising/foundrising_getAll.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	// Получаем объект главного окна по ID
	obj, err := b.GetObject("foundrising_getAll_window")
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
	foundrisings, err := UC.GetAll()
	if err == nil {
		for _, U := range foundrisings {
			addRow_foundrising(store, U)
		}
	} else {
		Error_window(err.Error())
	}
	// Отображаем все виджеты в окне
	win.ShowAll()

}
func addRow_foundrising(listStore *gtk.ListStore, F ents.Foundrising) {
	// Get an iterator for a new row at the end of the list store
	iter := listStore.Append()
	err := listStore.Set(iter, []int{0, 1, 2, 3, 4, 5, 6, 7}, []interface{}{F.Id, F.Found_id, F.Description,
		fmt.Sprintf("%.2f", F.Required_sum), fmt.Sprintf("%.2f", F.Current_sum),
		F.Philantrops_amount,
		F.Creation_date[:strings.Index(F.Creation_date, "T")], F.Closing_date.String[:strings.Index(F.Closing_date.String, "T")]})
	if err != nil {
		log.Fatal("Unable to add row:", err)
		fmt.Println("Unable to add row:", err)
	}
}

func Admin_foundrising_getByID_window(UC ctrls.FoundrisingController) {

	win, b := Get_window("glade/admin/admin_actions/foundrising/foundrising_getByID.glade", "foundrising_getByID_window")
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
			foundrising, err := UC.GetByID(id)
			if err == nil {
				Admin_foundrising_Page(foundrising)
			} else {
				Error_window(err.Error())
			}
		}
	})

}

func Admin_foundrising_getByFoundID_window(UC ctrls.FoundrisingController) {

	win, b := Get_window("glade/admin/admin_actions/foundrising/foundrising_getByFoundID.glade", "foundrising_getByFoundID_window")
	// Отображаем все виджеты в окне
	// Получаем поле ввода
	obj, _ := b.GetObject("getByFoundID_button")
	getByFound_button := obj.(*gtk.Button)
	win.ShowAll()
	obj, _ = b.GetObject("ID_entry")
	id_entry := obj.(*gtk.Entry)
	getByFound_button.Connect("clicked", func() {
		id, err := id_entry.GetText()
		if err == nil {
			foundrising, err := UC.GetByFoundId(id)
			if err == nil {
				Admin_foundrising_MultiPage(foundrising)
			} else {
				Error_window(err.Error())
			}
		}
	})
}

func Admin_foundrising_getByCrDate_window(UC ctrls.FoundrisingController) {

	win, b := Get_window("glade/admin/admin_actions/foundrising/foundrising_getBycrDate.glade", "foundrising_getByCrDate_window")

	obj, _ := b.GetObject("getByCrDate_button")
	getByCrDate_button := obj.(*gtk.Button)
	win.ShowAll()
	obj, _ = b.GetObject("date_entry")
	date_entry := obj.(*gtk.Entry)
	getByCrDate_button.Connect("clicked", func() {
		date, err := date_entry.GetText()
		if err == nil {
			foundrisings, err := UC.GetByCreDate(date)
			if err == nil {
				Admin_foundrising_MultiPage(foundrisings)
			} else {
				Error_window(err.Error())
			}
		}
	})
}

func Admin_foundrising_getByClDate_window(UC ctrls.FoundrisingController) {

	win, b := Get_window("glade/admin/admin_actions/foundrising/foundrising_getByclDate.glade", "foundrising_getByClDate_window")

	obj, _ := b.GetObject("getByClDate_button")
	getByClDate_button := obj.(*gtk.Button)
	win.ShowAll()
	obj, _ = b.GetObject("date_entry")
	date_entry := obj.(*gtk.Entry)
	getByClDate_button.Connect("clicked", func() {
		date, err := date_entry.GetText()
		if err == nil {
			foundrisings, err := UC.GetByCloDate(date)
			if err == nil {
				Admin_foundrising_MultiPage(foundrisings)
			} else {
				Error_window(err.Error())
			}
		}
	})
}

func Admin_foundrising_create_window(UC ctrls.FoundrisingController) {

	win, b := Get_window("glade/admin/admin_actions/foundrising/foundrising_create.glade", "foundrising_create_window")

	obj, _ := b.GetObject("create_button")
	create_button := obj.(*gtk.Button)
	win.ShowAll()
	obj, _ = b.GetObject("ID_entry")
	ID_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("sum_entry")
	sum_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("descr_entry")
	descr_entry := obj.(*gtk.TextView)

	var id, sum, descr string
	var err error
	create_button.Connect("clicked", func() {
		id, err = ID_entry.GetText()
		if err == nil {
			sum, err = sum_entry.GetText()
			if err == nil {
				descr, err = GetTextFromTextView(descr_entry)
				if err == nil {
					err := UC.Add(id, descr, sum)
					if err == nil {
						found, _ := UC.FndS.GetById(id)
						found.CurFoudrisingAmount += 1
						UC.FndS.FR.Update(found)
						Success_window()
					} else {
						Error_window(err.Error())
					}
				} else {
					Error_window(err.Error())
				}
			} else {
				Error_window(err.Error())
			}
		} else {
			Error_window(err.Error())
		}
	})
}
func Admin_foundrising_delete_window(UC ctrls.FoundrisingController) {

	win, b := Get_window("glade/admin/admin_actions/foundrising/foundrising_delete.glade", "foundrising_delete_window")
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
func Admin_foundrising_update_window(UC ctrls.FoundrisingController) {

	win, b := Get_window("glade/admin/admin_actions/foundrising/foundrising_update.glade", "foundrising_update_window")

	obj, _ := b.GetObject("update_button")
	update_button := obj.(*gtk.Button)
	win.ShowAll()

	obj, _ = b.GetObject("ID_entry")
	id_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("sum_entry")
	sum_entry := obj.(*gtk.Entry)

	obj, _ = b.GetObject("descr_entry")
	descr_entry := obj.(*gtk.TextView)

	var id, descr, sum string
	var err error
	update_button.Connect("clicked", func() {
		id, err = id_entry.GetText()
		if err == nil {
			descr, err = GetTextFromTextView(descr_entry)
			if err == nil {
				sum, err = sum_entry.GetText()
				if err == nil {
					err := UC.Update(id, descr, sum)
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
