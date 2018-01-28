package szamlazz

import (
	"bytes"
	"encoding/xml"
	"errors"
	"log"

	"github.com/imdario/mergo"
	"github.com/jung-kurt/gofpdf"
	"github.com/parnurzeal/gorequest"
)

type ClientOptions struct {
	XMLName				   xml.Name
	eInvoice               bool		`xml:"eszamla"`
	requestInvoiceDownload bool		`xml:"szamlaLetoltes"`
	downloadedInvoiceCount int		`xml:"szamlaLetoltesPld"`
	responseVersion        int 		`xml:"valaszVerzio"`
	user				   string 	`xml:"felhasznalo"`
	password			   string 	`xml:"jelszo"`
	passphrase			   string	`xml:"kulcstartojelszo"`
	writePdfFile		   bool
}

const (
	xmlHeader = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<xmlszamla xmlns=\"http://www.szamlazz.hu/xmlszamla\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:schemaLocation=\"http://www.szamlazz.hu/xmlszamla xmlszamla.xsd\">\n"
	xmlFooter = "</xmlszamla>"
	szamlazzURL = "https://www.szamlazz.hu/szamla/"
)

var defaultOptions = ClientOptions{
	XMLName: xml.Name{Local:"beallitasok"},
	eInvoice: false,
	requestInvoiceDownload: false,
	downloadedInvoiceCount: 1,
	responseVersion: 1,
	user: "",
	password: "",
}

type Client struct {
	options ClientOptions

}

func NewClient (options ClientOptions) (*Client, error) {
	c := new(Client)
	c.options = defaultOptions
	// merge with defaults
	if err := mergo.Merge(&c.options, options, mergo.WithOverride); err != nil {
		log.Println("Failed to merge options into default options")
		return nil, err
	}
	// validate config
	if len(c.options.user) < 1 {
		return nil, errors.New("valid User field missing from client options")
	}
	if len(c.options.password) < 1 {
		return nil, errors.New("valid password field missing from client options")
	}

	return c, nil
}

func (c *Client) IssueInvoice (invoice Invoice) (error, InvoiceResponse, gorequest.Response) {
	var iresp InvoiceResponse
	content, err := c.generateInvoiceXML(invoice)
	if err != nil {
		return err, iresp, nil
	}
	err, ir, resp := c.sendRequest("action-xmlagentxmlfile", content)
	original, ok := ir.(InvoiceResponse)
	if ok {
		return err, original, resp
	}
	log.Fatal("InvoiceResponse is not of type")
	return err, iresp, resp
}

func (c *Client) IssueDeliveryNote (deliveryNote DeliveryNote) (error, DeliveryNoteResponse, gorequest.Response) {
	var dn DeliveryNoteResponse
	content, err := c.generateDeliveryNoteXML(deliveryNote)
	if err != nil {
		return err, dn, nil
	}
	err, dnr, resp := c.sendRequest("action-xmlagentxmlfile", content)
	original, ok := dnr.(DeliveryNoteResponse)
	if ok {
		return err, original, resp
	}
	log.Fatal("DeliveryNoteResponse is not of type")
	return err, dn, resp
}

func (c *Client) setRequestInvoiceDownload(value bool) {
	c.options.requestInvoiceDownload = value
}

func (c *Client) generateInvoiceXML (invoice Invoice) (string, error) {
	err, settings := c.generateSettingsXML()
	if err != nil {
		log.Fatal("failed to generate settings header")
		return "", err
	}
	err, invoicePart := invoice.generateXML()
	if err != nil {
		log.Fatal("failed to gernerate invoice xml part")
		return "", err
	}

	// merge parts together and return
	return xmlHeader + settings + invoicePart + xmlFooter, nil
}

func (c *Client) generateSettingsXML() (error, string) {
	w := bytes.NewBufferString("")
	enc := xml.NewEncoder(w)
	enc.Indent("  ", "    ")
	if err := enc.Encode(&c.options); err != nil {
		// oooops ?
		log.Fatal("failed to generate config xml part")
		return err, ""
	}
	return nil, w.String()
}

func (c *Client) generateDeliveryNoteXML (deliveryNote DeliveryNote) (string, error) {
	err, settings := c.generateSettingsXML()
	if err != nil {
		log.Fatal("failed to generate settings header")
		return "", nil
	}
	err, deliveryNotePart := deliveryNote.generateXML()
	return xmlHeader + settings + deliveryNotePart + xmlFooter, nil
}

// returns: error, body, httpResponse
func (c *Client) sendRequest (fileFieldName string, data string) (error, interface{}, gorequest.Response) {
	var iresp InvoiceResponse

	formData := map[string]interface{} {
		fileFieldName: FormData{
			value: data,
			options: FormDataOptions{
				filename: "request.xml",
				contentType: "text/xml",
			},
		},
	}
	resp, body, errs := gorequest.New().Post(szamlazzURL).
		Type("multipart").
		Send(formData).
		End()
	if errs != nil {
		return errs[0], iresp, resp
	}

	if resp.StatusCode != 200 {
		return errors.New(resp.Status), iresp, resp
	}

	// if there was an error
	szlahuErrCode := resp.Header.Get("szlahu_error_code")
	if len(szlahuErrCode) > 0 {
		return errors.New(resp.Header.Get("szlahu_error")), iresp, resp
	}

	// if there was no error.... extract the information we need
	iresp = InvoiceResponse{
		invoiceId: resp.Header.Get("szlahu_szamlaszam"),
		netTotal: resp.Header.Get("szlahu_nettovegosszeg"),
		grossTotal: resp.Header.Get("szlahu_bruttovegosszeg"),
	}

	if c.options.requestInvoiceDownload {
		if c.options.responseVersion == 2 {
			pdf := gofpdf.New("P", "mm", "A4", "")
			pdf.AddPage()
			pdf.SetFont("Arial", "B", 16)
			pdf.Cell(40, 10, "OK")
			if pdf.Error() != nil {
				log.Fatal("not expecting error when rendering text")
				return pdf.Error(), iresp, nil
			}

			iresp.pdf = &bytes.Buffer{}
			errbuffer := pdf.Output(iresp.pdf)
			if errbuffer != nil {
				return errbuffer, iresp, nil
			}

			if c.options.writePdfFile {
				err := pdf.OutputFileAndClose("invoice_"+iresp.invoiceId+".pdf")
				if err != nil {
					return err, iresp, nil
				}
			}
		} else {
			b := bytes.NewBufferString(body)
			iresp.pdf = b

			return nil, iresp, resp
		}
	}
	// otherwise return the responses
	return nil, iresp, resp
}