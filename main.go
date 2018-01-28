package szamlazz

import (
	"log"
	"os"

	"github.com/nadilas/goszamlazz/constants"
)

func main() {
	log.Println("Starting dummy client...")
	opts := ClientOptions{}
	c, err := NewClient(opts)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	seller := Seller{

	}
	buyer := Buyer{

	}
	soldItem1 := Item{

	}
	in, err := NewInvoice(InvoiceOptions{
		paymentMethod: constants.Cash,
		seller: seller,
		buyer: buyer,
		items: []Item{soldItem1},
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	err, inresponse, _ := c.IssueInvoice(*in)
	if err != nil {
		log.Fatal(err)
		os.Exit(3)
	}

	log.Println(inresponse)
}