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
	paymentMethod   constants.EPaymentMethod
	currency        constants.ECurrency
	language        constants.ELanguage
	seller          Seller
	buyer           Buyer
	items           []Item
	issueDate       time.Time
	fulfillmentDate time.Time
	dueDate         time.Time
	comment         string
	orderNumber		string
	proforma		bool
	invoiceIdPrefix	string
	paid			bool
}

var defaultInvoiceOptions = InvoiceOptions{
	paymentMethod: constants.BankTransfer,
	currency:      constants.Ft,
	language:      constants.Hungarian,
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

	if i.options.issueDate.IsZero() {
		i.options.issueDate = time.Now()
	}

	if i.options.fulfillmentDate.IsZero() {
		i.options.fulfillmentDate = time.Now()
	}

	if i.options.dueDate.IsZero() {
		i.options.dueDate = time.Now()
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

	err, buyerString := in.options.buyer.generateXML()
	if err != nil {
		return err, ""
	}
	err, sellerString := in.options.seller.generateXML()
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
		issueDate:       in.options.issueDate,
		fulfillmentDate: in.options.fulfillmentDate,
		dueDate:         in.options.dueDate,
		paymentMethod:   in.options.paymentMethod.Value(),
		currency:        in.options.currency.Value(),
		language:        in.options.language.Value(),
		comment:         in.options.comment,
		// exchangeRateBank:	?
		// exchangeRate:	?
		orderNumber:	 in.options.orderNumber,
		// retainerInvoice:	?
		// finalInvoice: ?
		proforma:		 in.options.proforma,
		invoiceIdPrefix: in.options.invoiceIdPrefix,
		paid:			 in.options.paid,
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
