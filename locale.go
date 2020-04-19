package thailandpost

type Lang string

const (
	LangEN Lang = "EN"
	LangTH Lang = "TH"
	LangCH Lang = "CH"
	LangZH Lang = "CH" // alias for Chinese language
)

func isLanguageNeedBEConversion(lang Lang) bool {
	return lang == LangTH
}
