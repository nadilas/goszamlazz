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
	return Currencies[curr].Shortcode
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
	Shortcode     string
	RoundPriceExp int
	Name          string
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
	return Languages[lang].IsoCountryCode
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
	IsoCountryCode string
	Name           string
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
	return PaymentMethods[p].Name
}

func (p EPaymentMethod) Value () string {
	return PaymentMethods[p].Value
}

// the payment method constants
const (
	Cash EPaymentMethod = 1 + iota
	BankTransfer
	CreditCard
)

type PaymentMethod struct {
	Value string
	Name  string
}

var Currencies = map[ECurrency]Currency{
	Ft:  {Shortcode: Ft.String(), RoundPriceExp: 0, Name: "Hungarian Forint"},
	HUF: {Shortcode: HUF.String(), RoundPriceExp: 0, Name: "Hungarian Forint"},
	EUR: {Shortcode: EUR.String(), RoundPriceExp: 2, Name: "Euro"},
	CHF: {Shortcode: CHF.String(), RoundPriceExp: 2, Name: "Swiss Franc"},
	USD: {Shortcode: USD.String(), RoundPriceExp: 2, Name: "US Dollar"},
	AUD: {Shortcode: AUD.String(), RoundPriceExp: 2, Name: "Australian Dollar"},
	AED: {Shortcode: AED.String(), RoundPriceExp: 2, Name: "Emirati Dirham"},
	BGN: {Shortcode: BGN.String(), RoundPriceExp: 2, Name: "Bulgarian Lev"},
	CAD: {Shortcode: CAD.String(), RoundPriceExp: 2, Name: "Canadian Dollar"},
	CNY: {Shortcode: CNY.String(), RoundPriceExp: 2, Name: "Chinese Yuan Renminbi"},
	CZK: {Shortcode: CZK.String(), RoundPriceExp: 2, Name: "Czech Koruna"},
	DKK: {Shortcode: DKK.String(), RoundPriceExp: 2, Name: "Danish Krone"},
	EEK: {Shortcode: EEK.String(), RoundPriceExp: 2, Name: "Estonian Kroon"},
	GBP: {Shortcode: GBP.String(), RoundPriceExp: 2, Name: "British Pound"},
	HRK: {Shortcode: HRK.String(), RoundPriceExp: 2, Name: "Croatian Kuna"},
	ISK: {Shortcode: ISK.String(), RoundPriceExp: 2, Name: "Icelandic Krona"},
	JPY: {Shortcode: JPY.String(), RoundPriceExp: 2, Name: "Japanese Yen"},
	LTL: {Shortcode: LTL.String(), RoundPriceExp: 2, Name: "Lithuanian Litas"},
	LVL: {Shortcode: LVL.String(), RoundPriceExp: 2, Name: "Latvian Lats"},
	NOK: {Shortcode: NOK.String(), RoundPriceExp: 2, Name: "Norwegian Krone"},
	NZD: {Shortcode: NZD.String(), RoundPriceExp: 2, Name: "New Zealand Dollar"},
	PLN: {Shortcode: PLN.String(), RoundPriceExp: 2, Name: "Polish Zloty"},
	RON: {Shortcode: RON.String(), RoundPriceExp: 2, Name: "Romanian New Leu"},
	RUB: {Shortcode: RUB.String(), RoundPriceExp: 2, Name: "Russian Ruble"},
	SEK: {Shortcode: SEK.String(), RoundPriceExp: 2, Name: "Swedish Krona"},
	SKK: {Shortcode: SKK.String(), RoundPriceExp: 2, Name: "Slovak Koruna"},
	UAH: {Shortcode: UAH.String(), RoundPriceExp: 2, Name: "Ukrainian Hryvnia"},
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