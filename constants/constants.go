package constants

type ECurrency int

// internal string values of the enumß
var currencynames = [...]string {
	"Ft",
	"HUF",
	"EUR",
	"CHF",
	"USD",
	"AUD",
	"AED",
	"BGN",
	"CAD",
	"CNY",
	"CZK",
	"DKK",
	"EEK",
	"GBP",
	"HRK",
	"ISK",
	"JPY",
	"LTL",
	"LVL",
	"NOK",
	"NZD",
	"PLN",
	"RON",
	"RUB",
	"SEK",
	"SKK",
	"UAH",
}

func (curr ECurrency) String() string {
	return currencynames[curr-1]
}

func (curr ECurrency) Value () string {
	return Currencies[curr].shortcode
}

// the currency constants
const (
	Ft ECurrency = 1 + iota
	HUF
	EUR
	CHF
	USD
	AUD
	AED
	BGN
	CAD
	CNY
	CZK
	DKK
	EEK
	GBP
	HRK
	ISK
	JPY
	LTL
	LVL
	NOK
	NZD
	PLN
	RON
	RUB
	SEK
	SKK
	UAH
)

type Currency struct {
	shortcode     string
	roundPriceExp int
	name          string
}

type ELanguage int

var languagenames = [...]string {
	"Hungarian",
	"English",
	"German",
	"Italian",
	"Romanian",
	"Slovak",
}

func (lang ELanguage) String() string {
	return languagenames[lang - 1]
}

func (lang ELanguage) Value () string {
	return Languages[lang].isoCountryCode
}

// the language constants
const (
	Hungarian ELanguage = 1 + iota
	English
	German
	Italian
	Romanian
	Slovak
)

type Language struct {
	isoCountryCode string
	name           string
}

type EPaymentMethod int

var paymentMethodNames = [...]string {
	"Cash",
	"BankTransfer",
	"CreditCard",
}

func (p EPaymentMethod) String() string {
	return paymentMethodNames[p- 1]
}

func (p EPaymentMethod) Object () PaymentMethod {
	return PaymentMethods[p]
}

func (p EPaymentMethod) Name () string {
	return PaymentMethods[p].name
}

func (p EPaymentMethod) Value () string {
	return PaymentMethods[p].value
}

// the payment method constants
const (
	Cash EPaymentMethod = 1 + iota
	BankTransfer
	CreditCard
)

type PaymentMethod struct {
	value string
	name  string
}

var Currencies = map[ECurrency]Currency{
	Ft:  {shortcode: Ft.String(), roundPriceExp: 0, name: "Hungarian Forint"},
	HUF: {shortcode: HUF.String(), roundPriceExp: 0, name: "Hungarian Forint"},
	EUR: {shortcode: EUR.String(), roundPriceExp: 2, name: "Euro"},
	CHF: {shortcode: CHF.String(), roundPriceExp: 2, name: "Swiss Franc"},
	USD: {shortcode: USD.String(), roundPriceExp: 2, name: "US Dollar"},
	AUD: {shortcode: AUD.String(), roundPriceExp: 2, name: "Australian Dollar"},
	AED: {shortcode: AED.String(), roundPriceExp: 2, name: "Emirati Dirham"},
	BGN: {shortcode: BGN.String(), roundPriceExp: 2, name: "Bulgarian Lev"},
	CAD: {shortcode: CAD.String(), roundPriceExp: 2, name: "Canadian Dollar"},
	CNY: {shortcode: CNY.String(), roundPriceExp: 2, name: "Chinese Yuan Renminbi"},
	CZK: {shortcode: CZK.String(), roundPriceExp: 2, name: "Czech Koruna"},
	DKK: {shortcode: DKK.String(), roundPriceExp: 2, name: "Danish Krone"},
	EEK: {shortcode: EEK.String(), roundPriceExp: 2, name: "Estonian Kroon"},
	GBP: {shortcode: GBP.String(), roundPriceExp: 2, name: "British Pound"},
	HRK: {shortcode: HRK.String(), roundPriceExp: 2, name: "Croatian Kuna"},
	ISK: {shortcode: ISK.String(), roundPriceExp: 2, name: "Icelandic Krona"},
	JPY: {shortcode: JPY.String(), roundPriceExp: 2, name: "Japanese Yen"},
	LTL: {shortcode: LTL.String(), roundPriceExp: 2, name: "Lithuanian Litas"},
	LVL: {shortcode: LVL.String(), roundPriceExp: 2, name: "Latvian Lats"},
	NOK: {shortcode: NOK.String(), roundPriceExp: 2, name: "Norwegian Krone"},
	NZD: {shortcode: NZD.String(), roundPriceExp: 2, name: "New Zealand Dollar"},
	PLN: {shortcode: PLN.String(), roundPriceExp: 2, name: "Polish Zloty"},
	RON: {shortcode: RON.String(), roundPriceExp: 2, name: "Romanian New Leu"},
	RUB: {shortcode: RUB.String(), roundPriceExp: 2, name: "Russian Ruble"},
	SEK: {shortcode: SEK.String(), roundPriceExp: 2, name: "Swedish Krona"},
	SKK: {shortcode: SKK.String(), roundPriceExp: 2, name: "Slovak Koruna"},
	UAH: {shortcode: UAH.String(), roundPriceExp: 2, name: "Ukrainian Hryvnia"},
}

var Languages = map[ELanguage]Language{
	Hungarian: {"hu", Hungarian.String()},
	English:   {"en", English.String()},
	German:    {"de", German.String()},
	Italian:   {"it", Italian.String()},
	Romanian:  {"ro", Romanian.String()},
	Slovak:    {"sk", Slovak.String()},
}

var PaymentMethods = map[EPaymentMethod]PaymentMethod {
	BankTransfer: {"bank transfer", "Átutalás"},
	Cash:         {"cash", "Készpénz"},
	CreditCard:   {"credit card", "Bankkártya"},
}

func GetCurrency(currency ECurrency) Currency {
	return Currencies[currency]
}

func GetLanguage(lang ELanguage) Language {
	return Languages[lang]
}

func GetPaymentMethod(method EPaymentMethod) PaymentMethod {
	return PaymentMethods[method]
}