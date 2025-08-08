package intl

import "golang.org/x/text/language"

var preferred12Hour = map[string]struct{}{
	"AS": {}, "BT": {}, "DJ": {}, "ER": {}, "GH": {}, "IN": {}, "LS": {}, "PG": {}, "PW": {}, "SO": {},
	"TO": {}, "VU": {}, "WS": {}, "AL": {}, "TD": {}, "CY": {}, "GR": {}, "419": {}, "AR": {}, "BO": {},
	"CL": {}, "CO": {}, "CR": {}, "CU": {}, "DO": {}, "EC": {}, "GT": {}, "HN": {}, "KP": {}, "KR": {},
	"MX": {}, "NA": {}, "NI": {}, "PA": {}, "PE": {}, "PR": {}, "PY": {}, "SV": {}, "UY": {}, "VE": {},
	"BD": {}, "PK": {}, "AG": {}, "AU": {}, "BB": {}, "BM": {}, "BS": {}, "CA": {}, "DM": {}, "en-001": {},
	"en-HK": {}, "en-MY": {}, "FJ": {}, "FM": {}, "GD": {}, "GM": {}, "GU": {}, "GY": {}, "JM": {}, "KI": {},
	"KN": {}, "KY": {}, "LC": {}, "LR": {}, "MH": {}, "MP": {}, "MW": {}, "NZ": {}, "SB": {}, "SG": {},
	"SL": {}, "SS": {}, "SZ": {}, "TC": {}, "TT": {}, "UM": {}, "US": {}, "VC": {}, "VG": {}, "VI": {},
	"ZM": {}, "AE": {}, "ar-001": {}, "BH": {}, "DZ": {}, "EG": {}, "EH": {}, "HK": {}, "IQ": {}, "JO": {},
	"KW": {}, "LB": {}, "LY": {}, "MO": {}, "MR": {}, "OM": {}, "PH": {}, "PS": {}, "QA": {}, "SA": {},
	"SD": {}, "SY": {}, "TN": {}, "YE": {}, "hi-IN": {}, "kn-IN": {}, "ml-IN": {}, "te-IN": {}, "KH": {}, "ta-IN": {},
	"BN": {}, "MY": {}, "ET": {}, "gu-IN": {}, "mr-IN": {}, "pa-IN": {}, "TW": {},
}

func is12HourLocale(locale language.Tag) bool {
	if _, ok := preferred12Hour[locale.String()]; ok {
		return true
	}
	region, _ := locale.Region()
	if _, ok := preferred12Hour[region.String()]; ok {
		return true
	}
	return false
}
