package main

func endApp() {
	println("Aplikasi selesai")
	massage := recover()
	println("Pesan error:", massage)
}
func runApp(err bool) {
	defer endApp()
	if err {
		panic("APLIKASI ERROR")
	}
}
func main() {
	runApp(true)
}
