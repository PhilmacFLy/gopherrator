package feedtypes

import (
	"errors"
	"strings"
)

type language int

func (l language) MarshalJSON() ([]byte, error) {
	var Result string
	switch l {
	case langAA:
		Result = "aa"
	case langAB:
		Result = "ab"
	case langAE:
		Result = "ae"
	case langAF:
		Result = "af"
	case langAK:
		Result = "ak"
	case langAM:
		Result = "am"
	case langAN:
		Result = "an"
	case langAR:
		Result = "ar"
	case langAS:
		Result = "as"
	case langAV:
		Result = "av"
	case langAY:
		Result = "ay"
	case langAZ:
		Result = "az"
	case langBA:
		Result = "ba"
	case langBE:
		Result = "be"
	case langBG:
		Result = "bg"
	case langBH:
		Result = "bh"
	case langBI:
		Result = "bi"
	case langBM:
		Result = "bm"
	case langBN:
		Result = "bn"
	case langBO:
		Result = "bo"
	case langBR:
		Result = "br"
	case langBS:
		Result = "bs"
	case langCA:
		Result = "ca"
	case langCE:
		Result = "ce"
	case langCH:
		Result = "ch"
	case langCO:
		Result = "co"
	case langCR:
		Result = "cr"
	case langCS:
		Result = "cs"
	case langCU:
		Result = "cu"
	case langCV:
		Result = "cv"
	case langCY:
		Result = "cy"
	case langDA:
		Result = "da"
	case langDE:
		Result = "de"
	case langDV:
		Result = "dv"
	case langDZ:
		Result = "dz"
	case langEE:
		Result = "ee"
	case langEL:
		Result = "el"
	case langEN:
		Result = "en"
	case langEO:
		Result = "eo"
	case langES:
		Result = "es"
	case langET:
		Result = "et"
	case langEU:
		Result = "eu"
	case langFA:
		Result = "fa"
	case langFF:
		Result = "ff"
	case langFI:
		Result = "fi"
	case langFJ:
		Result = "fj"
	case langFO:
		Result = "fo"
	case langFR:
		Result = "fr"
	case langFY:
		Result = "fy"
	case langGA:
		Result = "ga"
	case langGD:
		Result = "gd"
	case langGL:
		Result = "gl"
	case langGN:
		Result = "gn"
	case langGU:
		Result = "gu"
	case langGV:
		Result = "gv"
	case langHA:
		Result = "ha"
	case langHE:
		Result = "he"
	case langHI:
		Result = "hi"
	case langHO:
		Result = "ho"
	case langHR:
		Result = "hr"
	case langHT:
		Result = "ht"
	case langHU:
		Result = "hu"
	case langHY:
		Result = "hy"
	case langHZ:
		Result = "hz"
	case langIA:
		Result = "ia"
	case langID:
		Result = "id"
	case langIE:
		Result = "ie"
	case langIG:
		Result = "ig"
	case langII:
		Result = "ii"
	case langIK:
		Result = "ik"
	case langIO:
		Result = "io"
	case langIS:
		Result = "is"
	case langIT:
		Result = "it"
	case langIU:
		Result = "iu"
	case langJA:
		Result = "ja"
	case langJV:
		Result = "jv"
	case langKA:
		Result = "ka"
	case langKG:
		Result = "kg"
	case langKI:
		Result = "ki"
	case langKJ:
		Result = "kj"
	case langKK:
		Result = "kk"
	case langKL:
		Result = "kl"
	case langKM:
		Result = "km"
	case langKN:
		Result = "kn"
	case langKO:
		Result = "ko"
	case langKR:
		Result = "kr"
	case langKS:
		Result = "ks"
	case langKU:
		Result = "ku"
	case langKV:
		Result = "kv"
	case langKW:
		Result = "kw"
	case langKY:
		Result = "ky"
	case langLA:
		Result = "la"
	case langLB:
		Result = "lb"
	case langLG:
		Result = "lg"
	case langLI:
		Result = "li"
	case langLN:
		Result = "ln"
	case langLO:
		Result = "lo"
	case langLT:
		Result = "lt"
	case langLU:
		Result = "lu"
	case langLV:
		Result = "lv"
	case langMG:
		Result = "mg"
	case langMH:
		Result = "mh"
	case langMI:
		Result = "mi"
	case langMK:
		Result = "mk"
	case langML:
		Result = "ml"
	case langMN:
		Result = "mn"
	case langMR:
		Result = "mr"
	case langMS:
		Result = "ms"
	case langMT:
		Result = "mt"
	case langMY:
		Result = "my"
	case langNA:
		Result = "na"
	case langNB:
		Result = "nb"
	case langND:
		Result = "nd"
	case langNE:
		Result = "ne"
	case langNG:
		Result = "ng"
	case langNL:
		Result = "nl"
	case langNN:
		Result = "nn"
	case langNO:
		Result = "no"
	case langNR:
		Result = "nr"
	case langNV:
		Result = "nv"
	case langNY:
		Result = "ny"
	case langOC:
		Result = "oc"
	case langOJ:
		Result = "oj"
	case langOM:
		Result = "om"
	case langOR:
		Result = "or"
	case langOS:
		Result = "os"
	case langPA:
		Result = "pa"
	case langPI:
		Result = "pi"
	case langPL:
		Result = "pl"
	case langPS:
		Result = "ps"
	case langPT:
		Result = "pt"
	case langQU:
		Result = "qu"
	case langRM:
		Result = "rm"
	case langRN:
		Result = "rn"
	case langRO:
		Result = "ro"
	case langRU:
		Result = "ru"
	case langRW:
		Result = "rw"
	case langSA:
		Result = "sa"
	case langSC:
		Result = "sc"
	case langSD:
		Result = "sd"
	case langSE:
		Result = "se"
	case langSG:
		Result = "sg"
	case langSH:
		Result = "sh"
	case langSI:
		Result = "si"
	case langSK:
		Result = "sk"
	case langSL:
		Result = "sl"
	case langSM:
		Result = "sm"
	case langSN:
		Result = "sn"
	case langSO:
		Result = "so"
	case langSQ:
		Result = "sq"
	case langSR:
		Result = "sr"
	case langSS:
		Result = "ss"
	case langST:
		Result = "st"
	case langSU:
		Result = "su"
	case langSV:
		Result = "sv"
	case langSW:
		Result = "sw"
	case langTA:
		Result = "ta"
	case langTE:
		Result = "te"
	case langTG:
		Result = "tg"
	case langTH:
		Result = "th"
	case langTI:
		Result = "ti"
	case langTK:
		Result = "tk"
	case langTL:
		Result = "tl"
	case langTN:
		Result = "tn"
	case langTO:
		Result = "to"
	case langTR:
		Result = "tr"
	case langTS:
		Result = "ts"
	case langTT:
		Result = "tt"
	case langTW:
		Result = "tw"
	case langTY:
		Result = "ty"
	case langUG:
		Result = "ug"
	case langUK:
		Result = "uk"
	case langUR:
		Result = "ur"
	case langUZ:
		Result = "uz"
	case langVE:
		Result = "ve"
	case langVI:
		Result = "vi"
	case langVO:
		Result = "vo"
	case langWA:
		Result = "wa"
	case langWO:
		Result = "wo"
	case langXH:
		Result = "xh"
	case langYI:
		Result = "yi"
	case langYO:
		Result = "yo"
	case langZA:
		Result = "za"
	case langZH:
		Result = "zh"
	case langZU:
		Result = "zu"
	default:
		return []byte(Result), errors.New("Not a valid ISO 639-1 language")
	}
	Result = "\"" + Result + "\""
	return []byte(Result), nil
}

func (l language) UnmarshalJSON(b []byte) error {
	lang := strings.Replace(string(b), "\"", "", -1)
	switch string(lang) {
	case "aa":
		l = langAA
	case "ab":
		l = langAB
	case "ae":
		l = langAE
	case "af":
		l = langAF
	case "ak":
		l = langAK
	case "am":
		l = langAM
	case "an":
		l = langAN
	case "ar":
		l = langAR
	case "as":
		l = langAS
	case "av":
		l = langAV
	case "ay":
		l = langAY
	case "az":
		l = langAZ
	case "ba":
		l = langBA
	case "be":
		l = langBE
	case "bg":
		l = langBG
	case "bh":
		l = langBH
	case "bi":
		l = langBI
	case "bm":
		l = langBM
	case "bn":
		l = langBN
	case "bo":
		l = langBO
	case "br":
		l = langBR
	case "bs":
		l = langBS
	case "ca":
		l = langCA
	case "ce":
		l = langCE
	case "ch":
		l = langCH
	case "co":
		l = langCO
	case "cr":
		l = langCR
	case "cs":
		l = langCS
	case "cu":
		l = langCU
	case "cv":
		l = langCV
	case "cy":
		l = langCY
	case "da":
		l = langDA
	case "de":
		l = langDE
	case "dv":
		l = langDV
	case "dz":
		l = langDZ
	case "ee":
		l = langEE
	case "el":
		l = langEL
	case "en":
		l = langEN
	case "eo":
		l = langEO
	case "es":
		l = langES
	case "et":
		l = langET
	case "eu":
		l = langEU
	case "fa":
		l = langFA
	case "ff":
		l = langFF
	case "fi":
		l = langFI
	case "fj":
		l = langFJ
	case "fo":
		l = langFO
	case "fr":
		l = langFR
	case "fy":
		l = langFY
	case "ga":
		l = langGA
	case "gd":
		l = langGD
	case "gl":
		l = langGL
	case "gn":
		l = langGN
	case "gu":
		l = langGU
	case "gv":
		l = langGV
	case "ha":
		l = langHA
	case "he":
		l = langHE
	case "hi":
		l = langHI
	case "ho":
		l = langHO
	case "hr":
		l = langHR
	case "ht":
		l = langHT
	case "hu":
		l = langHU
	case "hy":
		l = langHY
	case "hz":
		l = langHZ
	case "ia":
		l = langIA
	case "id":
		l = langID
	case "ie":
		l = langIE
	case "ig":
		l = langIG
	case "ii":
		l = langII
	case "ik":
		l = langIK
	case "io":
		l = langIO
	case "is":
		l = langIS
	case "it":
		l = langIT
	case "iu":
		l = langIU
	case "ja":
		l = langJA
	case "jv":
		l = langJV
	case "ka":
		l = langKA
	case "kg":
		l = langKG
	case "ki":
		l = langKI
	case "kj":
		l = langKJ
	case "kk":
		l = langKK
	case "kl":
		l = langKL
	case "km":
		l = langKM
	case "kn":
		l = langKN
	case "ko":
		l = langKO
	case "kr":
		l = langKR
	case "ks":
		l = langKS
	case "ku":
		l = langKU
	case "kv":
		l = langKV
	case "kw":
		l = langKW
	case "ky":
		l = langKY
	case "la":
		l = langLA
	case "lb":
		l = langLB
	case "lg":
		l = langLG
	case "li":
		l = langLI
	case "ln":
		l = langLN
	case "lo":
		l = langLO
	case "lt":
		l = langLT
	case "lu":
		l = langLU
	case "lv":
		l = langLV
	case "mg":
		l = langMG
	case "mh":
		l = langMH
	case "mi":
		l = langMI
	case "mk":
		l = langMK
	case "ml":
		l = langML
	case "mn":
		l = langMN
	case "mr":
		l = langMR
	case "ms":
		l = langMS
	case "mt":
		l = langMT
	case "my":
		l = langMY
	case "na":
		l = langNA
	case "nb":
		l = langNB
	case "nd":
		l = langND
	case "ne":
		l = langNE
	case "ng":
		l = langNG
	case "nl":
		l = langNL
	case "nn":
		l = langNN
	case "no":
		l = langNO
	case "nr":
		l = langNR
	case "nv":
		l = langNV
	case "ny":
		l = langNY
	case "oc":
		l = langOC
	case "oj":
		l = langOJ
	case "om":
		l = langOM
	case "or":
		l = langOR
	case "os":
		l = langOS
	case "pa":
		l = langPA
	case "pi":
		l = langPI
	case "pl":
		l = langPL
	case "ps":
		l = langPS
	case "pt":
		l = langPT
	case "qu":
		l = langQU
	case "rm":
		l = langRM
	case "rn":
		l = langRN
	case "ro":
		l = langRO
	case "ru":
		l = langRU
	case "rw":
		l = langRW
	case "sa":
		l = langSA
	case "sc":
		l = langSC
	case "sd":
		l = langSD
	case "se":
		l = langSE
	case "sg":
		l = langSG
	case "sh":
		l = langSH
	case "si":
		l = langSI
	case "sk":
		l = langSK
	case "sl":
		l = langSL
	case "sm":
		l = langSM
	case "sn":
		l = langSN
	case "so":
		l = langSO
	case "sq":
		l = langSQ
	case "sr":
		l = langSR
	case "ss":
		l = langSS
	case "st":
		l = langST
	case "su":
		l = langSU
	case "sv":
		l = langSV
	case "sw":
		l = langSW
	case "ta":
		l = langTA
	case "te":
		l = langTE
	case "tg":
		l = langTG
	case "th":
		l = langTH
	case "ti":
		l = langTI
	case "tk":
		l = langTK
	case "tl":
		l = langTL
	case "tn":
		l = langTN
	case "to":
		l = langTO
	case "tr":
		l = langTR
	case "ts":
		l = langTS
	case "tt":
		l = langTT
	case "tw":
		l = langTW
	case "ty":
		l = langTY
	case "ug":
		l = langUG
	case "uk":
		l = langUK
	case "ur":
		l = langUR
	case "uz":
		l = langUZ
	case "ve":
		l = langVE
	case "vi":
		l = langVI
	case "vo":
		l = langVO
	case "wa":
		l = langWA
	case "wo":
		l = langWO
	case "xh":
		l = langXH
	case "yi":
		l = langYI
	case "yo":
		l = langYO
	case "za":
		l = langZA
	case "zh":
		l = langZH
	case "zu":
		l = langZU
	default:
		return errors.New("Not a valid ISO 639-1 language")
	}
	return nil
}

func (l language) String() string {
	switch l {
	case langAA:
		return "Afar"
	case langAB:
		return "Abkhazian"
	case langAE:
		return "Avestan"
	case langAF:
		return "Afrikaans"
	case langAK:
		return "Akan"
	case langAM:
		return "Amharic"
	case langAN:
		return "Aragonese"
	case langAR:
		return "Arabic"
	case langAS:
		return "Assamese"
	case langAV:
		return "Avaric"
	case langAY:
		return "Aymara"
	case langAZ:
		return "Azerbaijani"
	case langBA:
		return "Bashkir"
	case langBE:
		return "Belarusian"
	case langBG:
		return "Bulgarian"
	case langBH:
		return "Bihari languages"
	case langBI:
		return "Bislama"
	case langBM:
		return "Bambara"
	case langBN:
		return "Bengali"
	case langBO:
		return "Tibetan"
	case langBR:
		return "Breton"
	case langBS:
		return "Bosnian"
	case langCA:
		return "Catalan"
	case langCE:
		return "Chechen"
	case langCH:
		return "Chamorro"
	case langCO:
		return "Corsican"
	case langCR:
		return "Cree"
	case langCS:
		return "Czech"
	case langCU:
		return "Church Slavic"
	case langCV:
		return "Chuvash"
	case langCY:
		return "Welsh"
	case langDA:
		return "Danish"
	case langDE:
		return "German"
	case langDV:
		return "Dhivehi"
	case langDZ:
		return "Dzongkha"
	case langEE:
		return "Ewe"
	case langEL:
		return "Modern Greek"
	case langEN:
		return "English"
	case langEO:
		return "Esperanto"
	case langES:
		return "Spanish"
	case langET:
		return "Estonian"
	case langEU:
		return "Basque"
	case langFA:
		return "Persian"
	case langFF:
		return "Fulah"
	case langFI:
		return "Finnish"
	case langFJ:
		return "Fijian"
	case langFO:
		return "Faroese"
	case langFR:
		return "French"
	case langFY:
		return "Western Frisian"
	case langGA:
		return "Irish"
	case langGD:
		return "Scottish Gaelic"
	case langGL:
		return "Galician"
	case langGN:
		return "Guarani"
	case langGU:
		return "Gujarati"
	case langGV:
		return "Manx"
	case langHA:
		return "Hausa"
	case langHE:
		return "Hebrew"
	case langHI:
		return "Hindi"
	case langHO:
		return "Hiri Motu"
	case langHR:
		return "Croatian"
	case langHT:
		return "Haitian"
	case langHU:
		return "Hungarian"
	case langHY:
		return "Armenian"
	case langHZ:
		return "Herero"
	case langIA:
		return "Interlingua"
	case langID:
		return "Indonesian"
	case langIE:
		return "Interlingue"
	case langIG:
		return "Igbo"
	case langII:
		return "Sichuan Yi"
	case langIK:
		return "Inupiaq"
	case langIO:
		return "Ido"
	case langIS:
		return "Icelandic"
	case langIT:
		return "Italian"
	case langIU:
		return "Inuktitut"
	case langJA:
		return "Japanese"
	case langJV:
		return "Javanese"
	case langKA:
		return "Georgian"
	case langKG:
		return "Kongo"
	case langKI:
		return "Kikuyu"
	case langKJ:
		return "Kuanyama"
	case langKK:
		return "Kazakh"
	case langKL:
		return "Kalaallisut"
	case langKM:
		return "Khmer"
	case langKN:
		return "Kannada"
	case langKO:
		return "Korean"
	case langKR:
		return "Kanuri"
	case langKS:
		return "Kashmiri"
	case langKU:
		return "Kurdish"
	case langKV:
		return "Komi"
	case langKW:
		return "Cornish"
	case langKY:
		return "Kirghiz"
	case langLA:
		return "Latin"
	case langLB:
		return "Luxembourgish"
	case langLG:
		return "Ganda"
	case langLI:
		return "Limburgan"
	case langLN:
		return "Lingala"
	case langLO:
		return "Lao"
	case langLT:
		return "Lithuanian"
	case langLU:
		return "Luba-Katanga"
	case langLV:
		return "Latvian"
	case langMG:
		return "Malagasy"
	case langMH:
		return "Marshallese"
	case langMI:
		return "Maori"
	case langMK:
		return "Macedonian"
	case langML:
		return "Malayalam"
	case langMN:
		return "Mongolian"
	case langMR:
		return "Marathi"
	case langMS:
		return "Malay"
	case langMT:
		return "Maltese"
	case langMY:
		return "Burmese"
	case langNA:
		return "Nauru"
	case langNB:
		return "Norwegian Bokmål"
	case langND:
		return "North Ndebele"
	case langNE:
		return "Nepali"
	case langNG:
		return "Ndonga"
	case langNL:
		return "Dutch"
	case langNN:
		return "Norwegian Nynorsk"
	case langNO:
		return "Norwegian"
	case langNR:
		return "South Ndebele"
	case langNV:
		return "Navajo"
	case langNY:
		return "Nyanja"
	case langOC:
		return "Occitan"
	case langOJ:
		return "Ojibwa"
	case langOM:
		return "Oromo"
	case langOR:
		return "Oriya"
	case langOS:
		return "Ossetian"
	case langPA:
		return "Panjabi"
	case langPI:
		return "Pali"
	case langPL:
		return "Polish"
	case langPS:
		return "Pushto"
	case langPT:
		return "Portuguese"
	case langQU:
		return "Quechua"
	case langRM:
		return "Romansh"
	case langRN:
		return "Rundi"
	case langRO:
		return "Romanian"
	case langRU:
		return "Russian"
	case langRW:
		return "Kinyarwanda"
	case langSA:
		return "Sanskrit"
	case langSC:
		return "Sardinian"
	case langSD:
		return "Sindhi"
	case langSE:
		return "Northern Sami"
	case langSG:
		return "Sango"
	case langSH:
		return "Serbo-Croatian"
	case langSI:
		return "Sinhala"
	case langSK:
		return "Slovak"
	case langSL:
		return "Slovenian"
	case langSM:
		return "Samoan"
	case langSN:
		return "Shona"
	case langSO:
		return "Somali"
	case langSQ:
		return "Albanian"
	case langSR:
		return "Serbian"
	case langSS:
		return "Swati"
	case langST:
		return "Southern Sotho"
	case langSU:
		return "Sundanese"
	case langSV:
		return "Swedish"
	case langSW:
		return "Swahili"
	case langTA:
		return "Tamil"
	case langTE:
		return "Telugu"
	case langTG:
		return "Tajik"
	case langTH:
		return "Thai"
	case langTI:
		return "Tigrinya"
	case langTK:
		return "Turkmen"
	case langTL:
		return "Tagalog"
	case langTN:
		return "Tswana"
	case langTO:
		return "Tonga"
	case langTR:
		return "Turkish"
	case langTS:
		return "Tsonga"
	case langTT:
		return "Tatar"
	case langTW:
		return "Twi"
	case langTY:
		return "Tahitian"
	case langUG:
		return "Uighur"
	case langUK:
		return "Ukrainian"
	case langUR:
		return "Urdu"
	case langUZ:
		return "Uzbek"
	case langVE:
		return "Venda"
	case langVI:
		return "Vietnamese"
	case langVO:
		return "Volapük"
	case langWA:
		return "Walloon"
	case langWO:
		return "Wolof"
	case langXH:
		return "Xhosa"
	case langYI:
		return "Yiddish"
	case langYO:
		return "Yoruba"
	case langZA:
		return "Zhuang"
	case langZH:
		return "Chinese"
	case langZU:
		return "Zulu"
	default:
		return "!!!!!ERROR!!!!!"
	}
}

const (
	langAA language = iota + 1
	langAB
	langAE
	langAF
	langAK
	langAM
	langAN
	langAR
	langAS
	langAV
	langAY
	langAZ
	langBA
	langBE
	langBG
	langBH
	langBI
	langBM
	langBN
	langBO
	langBR
	langBS
	langCA
	langCE
	langCH
	langCO
	langCR
	langCS
	langCU
	langCV
	langCY
	langDA
	langDE
	langDV
	langDZ
	langEE
	langEL
	langEN
	langEO
	langES
	langET
	langEU
	langFA
	langFF
	langFI
	langFJ
	langFO
	langFR
	langFY
	langGA
	langGD
	langGL
	langGN
	langGU
	langGV
	langHA
	langHE
	langHI
	langHO
	langHR
	langHT
	langHU
	langHY
	langHZ
	langIA
	langID
	langIE
	langIG
	langII
	langIK
	langIO
	langIS
	langIT
	langIU
	langJA
	langJV
	langKA
	langKG
	langKI
	langKJ
	langKK
	langKL
	langKM
	langKN
	langKO
	langKR
	langKS
	langKU
	langKV
	langKW
	langKY
	langLA
	langLB
	langLG
	langLI
	langLN
	langLO
	langLT
	langLU
	langLV
	langMG
	langMH
	langMI
	langMK
	langML
	langMN
	langMR
	langMS
	langMT
	langMY
	langNA
	langNB
	langND
	langNE
	langNG
	langNL
	langNN
	langNO
	langNR
	langNV
	langNY
	langOC
	langOJ
	langOM
	langOR
	langOS
	langPA
	langPI
	langPL
	langPS
	langPT
	langQU
	langRM
	langRN
	langRO
	langRU
	langRW
	langSA
	langSC
	langSD
	langSE
	langSG
	langSH
	langSI
	langSK
	langSL
	langSM
	langSN
	langSO
	langSQ
	langSR
	langSS
	langST
	langSU
	langSV
	langSW
	langTA
	langTE
	langTG
	langTH
	langTI
	langTK
	langTL
	langTN
	langTO
	langTR
	langTS
	langTT
	langTW
	langTY
	langUG
	langUK
	langUR
	langUZ
	langVE
	langVI
	langVO
	langWA
	langWO
	langXH
	langYI
	langYO
	langZA
	langZH
	langZU
)
