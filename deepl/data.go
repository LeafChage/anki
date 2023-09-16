package deepl

type SourceLang string

const (
	SourceLangBulgarian  SourceLang = "BG"
	SourceLangCzech      SourceLang = "CS"
	SourceLangDanish     SourceLang = "DA"
	SourceLangGerman     SourceLang = "DE"
	SourceLangGreek      SourceLang = "EL"
	SourceLangEnglish    SourceLang = "EN"
	SourceLangSpanish    SourceLang = "ES"
	SourceLangEstonian   SourceLang = "ET"
	SourceLangFinnish    SourceLang = "FI"
	SourceLangFrench     SourceLang = "FR"
	SourceLangHungarian  SourceLang = "HU"
	SourceLangIndonesian SourceLang = "ID"
	SourceLangItalian    SourceLang = "IT"
	SourceLangJapanese   SourceLang = "JA"
	SourceLangKorean     SourceLang = "KO"
	SourceLangLithuanian SourceLang = "LT"
	SourceLangLatvian    SourceLang = "LV"
	SourceLangNorwegian  SourceLang = "NB"
	SourceLangDutch      SourceLang = "NL"
	SourceLangPolish     SourceLang = "PL"
	SourceLangPortuguese SourceLang = "PT"
	SourceLangRomanian   SourceLang = "RO"
	SourceLangRussian    SourceLang = "RU"
	SourceLangSlovak     SourceLang = "SK"
	SourceLangSlovenian  SourceLang = "SL"
	SourceLangSwedish    SourceLang = "SV"
	SourceLangTurkish    SourceLang = "TR"
	SourceLangUkrainian  SourceLang = "UK"
	SourceLangChinese    SourceLang = "ZH"
)

type TargetLang string

const (
	TargetLangBulgarian = "BG"
	TargetLangCzech     = "CS"
	TargetLangDanish    = "DA"
	TargetLangGerman    = "DE"
	TargetLangGreek     = "EL"
	// unspecified variant for backward compatibility; please select EN-GB or EN-US instead
	TargetLangEnglish         = "EN"
	TargetLangEnglishBritish  = "EN-GB"
	TargetLangEnglishAmerican = "EN-US"
	TargetLangSpanish         = "ES"
	TargetLangEstonian        = "ET"
	TargetLangFinnish         = "FI"
	TargetLangFrench          = "FR"
	TargetLangHungarian       = "HU"
	TargetLangIndonesian      = "ID"
	TargetLangItalian         = "IT"
	TargetLangJapanese        = "JA"
	TargetLangKorean          = "KO"
	TargetLangLithuanian      = "LT"
	TargetLangLatvian         = "LV"
	TargetLangNorwegian       = "NB"
	TargetLangDutch           = "NL"
	TargetLangPolish          = "PL"
	// unspecified variant for backward compatibility; please select PT-BR or PT-PT instead
	TargetLangPortuguese = "PT"
	// Brazilian
	TargetLangPortugueseBR = "PT-BR"
	// all Portuguese varieties excluding Brazilian Portuguese
	TargetLangPortuguesePT      = "PT-PT"
	TargetLangRomanian          = "RO"
	TargetLangRussian           = "RU"
	TargetLangSlovak            = "SK"
	TargetLangSlovenian         = "SL"
	TargetLangSwedish           = "SV"
	TargetLangTurkish           = "TR"
	TargetLangUkrainian         = "UK"
	TargetLangChineseSimplified = "ZH"
)
