package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type Language struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func languageSeeder() {
	data := []string{
		"AGTA",
		"AGUTAYNON",
		"AKLANON",
		"ALANGAN",
		"AMBALA",
		"ATA",
		"ATI",
		"ATTA",
		"AYTA",
		"BAGOBO",
		"BALANGAO",
		"BALOGA",
		"BANTUANON",
		"BATAGNON",
		"BATAK",
		"BIKOLANO",
		"BINUKID",
		"BOHOLANO",
		"BOLINAO",
		"BONTOC",
		"BUHID",
		"CALUYANUN",
		"CAPIZNON",
		"CEBUANO",
		"CHAVACANO",
		"CUYONON",
		"DAVAWENO",
		"DAVAWENO ZAMBOANGENO",
		"DUMAGAT",
		"GA'DANG",
		"HANONOO",
		"HILIGAYNON",
		"IBALOI",
		"IBANAG",
		"IBATAAN",
		"IFUGAO",
		"ILOCANO",
		"ILONGGO",
		"ILONGOT",
		"IRAYA",
		"ISINAI",
		"ISNAG",
		"ITAWES",
		"ITNEG",
		"IVATAN",
		"IWANK",
		"KAGAYANEN",
		"KALAGAN",
		"KALINGA",
		"KALLAHAN",
		"KAMAYO",
		"KANKANAEY",
		"KAPAMPANGAN",
		"KARAO",
		"KAROLANOS",
		"KASIGURAN",
		"K'BLAAN",
		"KINARAY - A",
		"LOOCNON",
		"MAGAHAT",
		"MAGINDANAON",
		"MALAYNON",
		"MAMANWA",
		"MANDAYA",
		"MANOBO",
		"MARANAO",
		"MASBATEÑO",
		"MOLBOG",
		"PALAWANO",
		"PANGALATOK",
		"PANGASINENSE",
		"PARANAN",
		"POROHANON",
		"ROMBLOMANON",
		"SAMA",
		"SANGIHE",
		"SORSOGON",
		"SUBANON",
		"SULOD",
		"SURIGAONON",
		"TADYAWAN",
		"TAGBANWA",
		"TAUSUG",
		"TAWBUID",
		"T'BOLI",
		"TIRURAY",
		"WARAY - WARAY",
		"YAKAN",
		"YOGAD",
		"ZAMBAL",
		"AFRIKAANS",
		"AINU",
		"AKKADIAN",
		"ALBANIAN",
		"AMHARIC",
		"ANCIENT GREEK",
		"ANGLOROMANI",
		"ARABIC",
		"ARMENIAN",
		"ASTURIAN",
		"AYMARA",
		"BAHASA INDONESIA",
		"BASQUE",
		"BAVARIAN",
		"BELORUSSIAN",
		"BENGALI",
		"BRETON",
		"BULGARIAN",
		"BURMESE",
		"CAMBODIAN",
		"CANTONESE",
		"CATALAN",
		"CHICHEWA",
		"CHINESE",
		"CHURCH SLAVONIC",
		"CIMBRIAN",
		"CORNISH",
		"CORSICAN",
		"CROATIAN",
		"CZECH",
		"DAKOTA",
		"DANISH",
		"DOMARI",
		"DONGXIANG",
		"DRAVIDIAN",
		"DUTCH",
		"ENGLISH",
		"ESTONIAN",
		"FAROESE",
		"FARSI",
		"FINNISH",
		"FRANKISH",
		"FRENCH",
		"FRISIAN",
		"FRIULIAN",
		"GAELIC",
		"GALICIAN",
		"GEORGIAN",
		"GERMAN",
		"GOTHIC",
		"GREEK",
		"GREENLANDIC",
		"GUARANI",
		"GUJARATI",
		"HAUSA",
		"HAWAIIAN",
		"HEBREW",
		"HINDI",
		"HITTITE",
		"HMONG",
		"HUNGARIAN",
		"ICELANDIC",
		"INDIC LANGUAGE",
		"INDONESIAN",
		"INGUSH",
		"INUIT",
		"IRISH",
		"ITALIAN",
		"JAPANESE",
		"JUDEO",
		"JUTISH",
		"KAYORT",
		"KOLI",
		"KONKANI",
		"KOREAN",
		"KURDISH",
		"LADIN",
		"LADINO",
		"LAKHOTA",
		"LATIN",
		"LATVIAN",
		"LITHUANIAN",
		"MACEDONIAN",
		"MALAY",
		"MANDARIN",
		"MANX GAELIC",
		"MAORI",
		"MARATHI",
		"MONGOLIAN",
		"NEPALI",
		"NORWEGIAN",
		"OJIBWE",
		"OSETIN",
		"PALI",
		"PASHTO",
		"PERSIAN",
		"PIDGIN",
		"POLISH",
		"PONTIC",
		"PORTUGUESE",
		"PUNJABI",
		"QUECHUA",
		"RHAETO - ROMANCE",
		"ROMANIAN",
		"ROMANY",
		"RUSSIAN",
		"SANSKRIT",
		"SARDINIAN",
		"SAXON",
		"SCOTS",
		"SERBIAN",
		"SHELTA",
		"SHIYEYI",
		"SICILIAN",
		"SINDHI",
		"SINHALA",
		"SINHALESE",
		"SLOVAK",
		"SLOVENE",
		"SLOVENIAN",
		"SOMALI",
		"SPANISH",
		"SUMERIAN",
		"SWABIAN",
		"SWAHILI",
		"SWEDISH",
		"TAGALOG",
		"TAGNOBI",
		"TAMAZIGHT",
		"TAMIL",
		"THAI",
		"THARU",
		"TIBETAN",
		"TOK PISIN",
		"TURKISH",
		"UKRAINIAN",
		"URDU",
		"VIETNAMESE",
		"WELSH",
		"YENICHE",
		"YEVANIC",
		"YIDDISH",
		"YORUBA",
		"ZULU",
	}

	for _, name := range data {
		language := Language{Name: strings.ToUpper(name)}
		language.Create()
	}
}

func (language *Language) Create() *Language {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&language).Error; err != nil {
		panic(err)
	}
	return language
}

func (language Language) Index(q string) []Language {
	db := database.Conn()
	defer db.Close()

	languages := []Language{}
	results := make(chan []Language)

	go func() {
		db.Find(&languages, "name LIKE ?", database.WrapLike(q))
		results <- languages
	}()
	return <-results
}
