package szamlazz

import (
	"bytes"
	"encoding/xml"
	"log"
	"time"

	"github.com/imdario/mergo"
	"github.com/nadilas/goszamlazz/constants"
)

type Buyer struct {
}

func (b *Buyer) generateXML() (error, string) {
	// todo
	return nil, ""
}

type Seller struct {
}

func (s *Seller) generateXML() (error, string) {
	// todo
	return nil, ""
}

type Item struct {
}

type InvoiceOptions struct {
	PaymentMethod   constants.EPaymentMethod
	Currency        constants.ECurrency
	Language        constants.ELanguage
	Seller          Seller
	Buyer           Buyer
	Items           []Item
	IssueDate       time.Time
	FulfillmentDate time.Time
	DueDate         time.Time
	Comment         string
	OrderNumber     string
	Proforma        bool
	InvoiceIdPrefix string
	Paid            bool
}

var defaultInvoiceOptions = InvoiceOptions{
	PaymentMethod: constants.BankTransfer,
	Currency:      constants.Ft,
	Language:      constants.Hungarian,
}

type Invoice struct {
	options InvoiceOptions
}

type InvoiceHeader struct {
	XMLName         xml.Name
	issueDate       time.Time	`json:"keltDatum"`
	fulfillmentDate time.Time	`json:"teljesitesDatum"`
	dueDate         time.Time	`json:"fizetesiHataridoDatum"`
	paymentMethod   string		`json:"fizmod"`
	currency        string		`json:"penznem"`
	language        string		`json:"szamlaNyelve"`
	comment         string		`json:"megjegyzes"`
	orderNumber     string		`json:"rendelesSzam"`
	proforma        bool		`json:"dijbekero"`
	invoiceIdPrefix string		`json:"szamlaszamElotag"`
	paid            bool		`json:"fizetve"`
}

func NewInvoice(opts InvoiceOptions) (*Invoice, error) {
	i := new(Invoice)
	i.options = defaultInvoiceOptions
	if err := mergo.Merge(&i.options, opts, mergo.WithOverride); err != nil {
		log.Fatal("failed to merge options into default options")
		return nil, err
	}

	if i.options.IssueDate.IsZero() {
		i.options.IssueDate = time.Now()
	}

	if i.options.FulfillmentDate.IsZero() {
		i.options.FulfillmentDate = time.Now()
	}

	if i.options.DueDate.IsZero() {
		i.options.DueDate = time.Now()
	}

	return i, nil
}

func (in *Invoice) generateXML() (error, string) {
	b := bytes.NewBufferString("")
	enc := xml.NewEncoder(b)
	enc.Indent("  ", "    ")
	if err := enc.Encode(in.getHeader()); err != nil {
		return err, ""
	}

	err, buyerString := in.options.Buyer.generateXML()
	if err != nil {
		return err, ""
	}
	err, sellerString := in.options.Seller.generateXML()
	if err != nil {
		return err, ""
	}
	headerString := b.String()
	itemsString := in.getItemsXML()

	xmlString := headerString + sellerString + buyerString + itemsString

	return nil, xmlString
}

func (in *Invoice) getItemsXML() string {
	// todo
	return ""
}

func (in *Invoice) getHeader() (InvoiceHeader) {
	return InvoiceHeader{
		XMLName:         xml.Name{Local: "fejlec"},
		issueDate:       in.options.IssueDate,
		fulfillmentDate: in.options.FulfillmentDate,
		dueDate:         in.options.DueDate,
		paymentMethod:   in.options.PaymentMethod.Value(),
		currency:        in.options.Currency.Value(),
		language:        in.options.Language.Value(),
		comment:         in.options.Comment,
		// exchangeRateBank:	?
		// exchangeRate:	?
		orderNumber:	 in.options.OrderNumber,
		// retainerInvoice:	?
		// finalInvoice: ?
		proforma:		 in.options.Proforma,
		invoiceIdPrefix: in.options.InvoiceIdPrefix,
		paid:			 in.options.Paid,
	}
}

type DeliveryNote struct {
}

func (dn *DeliveryNote) generateXML() (error, string) {
	// todo
	return nil, ""
}

type FormData struct {
	value   string
	options FormDataOptions
}

type FormDataOptions struct {
	filename    string
	contentType string
}

type InvoiceResponse struct {
	invoiceId  string
	netTotal   string
	grossTotal string
	pdf        *bytes.Buffer
}

type DeliveryNoteResponse struct {
}
