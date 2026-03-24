package components

import "github.com/a-h/templ"

func IndexView() templ.Component {
	return Assemble(&PageComponents{
		Header: header(&Translations{
			Czech:   "Pomůžeme vám uchovat webové zdroje k vaší práci",
			English: "We will help preserve resources used in your work",
		}),
		Main: indexView(),
	})
}

func GroupView(data *GroupViewData) templ.Component {
	return Assemble(&PageComponents{
		Title: &Translations{
			Czech:   "Linkra - Stav archivace",
			English: "Linkra - Capture state",
		},
		Header: header(&Translations{
			Czech:   "Zde můžete sledovat průběh archivace",
			English: "Here you can see the archiving progress",
		}),
		Main: groupView(data),
	})
}

func SeedView(data *SeedViewData) templ.Component {
	return Assemble(&PageComponents{
		Title: data.Title,
		Header: header(&Translations{
			Czech:   "Přehled informací o zdroji",
			English: "Overview of resource",
		}),
		Main: seedView(data),
	})
}

func ErrorView(data *ErrorViewData) templ.Component {
	return Assemble(&PageComponents{
		Title:  data.Title,
		Header: header(&Translations{Czech: "", English: ""}),
		Main:   errorView(data),
	})
}

func GeneratorView(data *GeneratorViewData) templ.Component {
	return Assemble(&PageComponents{
		Title: &Translations{
			Czech:   "Linkra - Generátor citací",
			English: "Linkra - Citation generator",
		},
		Header: header(&Translations{
			Czech:   "Pomůžeme vám vytvořit citace k vašim webovým zdrojům",
			English: "We will help you cite your sources",
		}),
		Main: generatorView(data),
	})
}

func RedirectErrorView(data *RedirectErrorViewData) templ.Component {
	return Assemble(&PageComponents{
		Title: &Translations{
			Czech:   "Linkra - Chyba při přesměrování",
			English: "Linkra - Error during redirection",
		},
		Header: header(&Translations{Czech: "", English: ""}),
		Main:   redirectErrorView(data),
	})
}
