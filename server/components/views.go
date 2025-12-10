package components

import "github.com/a-h/templ"

func IndexView() templ.Component {
	return Assemble(&PageComponents{
		Header: header("Pomůžeme vám uchovat webové zdroje k vaší práci"),
		Main:   indexView(),
	})
}

func GroupView(data *GroupViewData) templ.Component {
	return Assemble(&PageComponents{
		Title:  "Linkra - Stav archivace",
		Header: header("Zde můžete sledovat průběh archivace"),
		Main:   groupView(data),
	})
}

func SeedView(data *SeedViewData) templ.Component {
	return Assemble(&PageComponents{
		Title:  data.Title,
		Header: header("Přehled informací o zdroji"),
		Main:   seedView(data),
	})
}

func ErrorView(data *ErrorViewData) templ.Component {
	return Assemble(&PageComponents{
		Title:  data.Title,
		Header: header(""),
		Main:   errorView(data),
	})
}

func GeneratorView(data *GeneratorViewData) templ.Component {
	return Assemble(&PageComponents{
		Title:  "Linkra - Generátor citací",
		Header: header("Pomůžeme vám vytvořit citace k vašim webovým zdrojům"),
		Main:   generatorView(data),
	})
}

func RedirectErrorView(data *RedirectErrorViewData) templ.Component {
	return Assemble(&PageComponents{
		Title:  "Linkra - Chyba při přesměrování",
		Header: header(""),
		Main:   redirectErrorView(data),
	})
}
