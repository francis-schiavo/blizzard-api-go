package blizzard_api

type Namespace string

const (
	NONE_NS = Namespace("")
	STATIC_NS = Namespace("static-%s")
	CLASSIC_NS = Namespace("static-classic-%s")
	DYNAMIC_NS = Namespace("dynamic-%s")
	PROFILE_NS = Namespace("profile-%s")
)
func (namespace Namespace) String() string {
	return string(namespace)
}

type EndpointType int8

const (
	_ EndpointType = iota
	COMMUNITY
	PROFILE
	DATA
	MEDIA
	OAUTH
	SEARCH
)

type Game string
const (
	WoW = Game("wow")
	D3 = Game("d3")
	SC2 = Game("sc2")
	HS = Game("hs")
)

type Locale string
const (
	All = Locale("")
	EnUS = Locale("en_US")
	EsMX = Locale("es_MX")
	PtBR = Locale("pt_BR")
	EnGB = Locale("en_GB")
	EsES = Locale("es_ES")
	FrFR = Locale("fr_FR")
	RuRU = Locale("ru_RU")
	DeDE = Locale("de_DE")
	PtPT = Locale("pt_PT")
	ItIT = Locale("it_IT")
	KoKR = Locale("ko_KR")
	ZhTW = Locale("zh_TW")
	ZhCN = Locale("zh_CN")
)

func (locale Locale) String() string {
	return string(locale)
}

type Region string
const (
	US = Region("us")
	EU = Region("eu")
	KR = Region("kr")
	TW = Region("tw")
)
func (region Region) String() string {
	return string(region)
}