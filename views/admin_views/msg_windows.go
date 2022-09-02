package admin_views

import "github.com/gotk3/gotk3/gtk"

func Success_window() {
	win, _ := Get_window("glade/success.glade", "success_window")
	win.ShowAll()

}
func Error_window(errmsg string) {
	win, b := Get_window("glade/error.glade", "error_window")
	obj, _ := b.GetObject("err_label")
	err_label := obj.(*gtk.Label)
	msg := "ОШИБКА\n" + errmsg
	err_label.SetText(msg)
	win.ShowAll()

}
func NotFound_window() {
	win, _ := Get_window("glade/NotFound.glade", "NotFound_window")
	win.ShowAll()

}
func AlreadyExists_window() {
	win, _ := Get_window("glade/AlreadyExists.glade", "AlreadyExists_window")
	win.ShowAll()

}

func ErrCountry_window() {
	win, _ := Get_window("glade/errCountry.glade", "errCountry_window")
	win.ShowAll()

}
