package components

import "github.com/a-h/templ"

func IndexView() templ.Component {
	return Assemble(&PageComponents{
		Header: indexHeader(),
		Main:   indexView(),
	})
}

func GroupView(data *GroupViewData) templ.Component {
	return Assemble(&PageComponents{
		Title:  "Linkra - Stav archivace",
		Header: groupHeader(),
		Main:   groupView(data),
	})
}

func SeedView(data *SeedViewData) templ.Component {
	return Assemble(&PageComponents{
		Title:  data.Title,
		Header: seedHeader(data.Seed.URL),
		Main:   seedView(data),
	})
}

func ErrorView(data *ErrorViewData) templ.Component {
	return Assemble(&PageComponents{
		Title:  data.Title,
		Header: errorHeader(data),
		Main:   errorView(data),
	})
}

func GeneratorView(data *GeneratorViewData) templ.Component {
	return Assemble(&PageComponents{
		Title:  "Linkra - Generátor citací",
		Header: generatorHeader(),
		Main:   generatorView(data),
	})
}

func RedirectErrorView(data *RedirectErrorViewData) templ.Component {
	return Assemble(&PageComponents{
		Title:  "Linkra - Chyba při přesměrování",
		Header: redirectErrorHeader(),
		Main:   redirectErrorView(data),
	})
}
