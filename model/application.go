package model

type Application struct {
	id     string
	course *Course
}

func (app *Application) Course() *Course {
	return app.course
}

func (app *Application) Id() string {
	return app.id
}
