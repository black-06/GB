package main

// PDFViewerApplication.save()
func main() {
	err := AutoGBT2260()
	if err != nil {
		panic(err)
	}
	err = UpdateSource("GB_T_38635.1-2020")
	if err != nil {
		panic(err)
	}
	err = UpdateReadme()
	if err != nil {
		panic(err)
	}
}
