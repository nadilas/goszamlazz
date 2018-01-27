# goszamlazz
A Go client for Szamlazz.hu based on [https://github.com/ewngs/szamlazz.js](https://github.com/ewngs/szamlazz.js)

## Installation

```
glide get http://github.com/nadilas/goszamlazz
```

## Usage

```go
package main

import "github.com/nadilas/goszamlazz

```

You can reuse this client to issue invoices.

### Create a client

```go
client := szamlazz.Client{
    user: "USERNAME",
    password: "PASSWORD",
    eInvoice: false, // create e-invoice. optional, default: false
    passphrase: "", // passphrase for e-invoice. optional
    requestInvoiceDownload: true, // downloads the issued pdf invoice. optional, default: false
    downloadedInvoiceCount: 1, // optional, default: 1
    responseVersion: 1, // optional, default: 1
}
```

### Create a seller

```go
seller := szamlazz.Seller{
	bank: szamlazz.Bank{
		name: "Test Bank Name",
		accountNumber: "11111111-11111111-11111111",
	},
	email: szamlazz.Email{
		replyToAddress: "test@gmail.com",
		subject: "Invoice email subject",
		message: "This is an email message",
	},
	issuerName: "",
}
```

### Create a buyer

```go
buyer := szamlazz.Buyer{
	name: "Some buyer name " + Math.random(),
	country: "",
	zip: "",
	city: "",
	address: "",
	taxNumber: "",
	postAddress: szamlazz.PostAddress{
		name: "Some buyer name",
		zip: "",
		city: "",
		address: "",
	},
	issuerName: "",
	identifier: 1,
	phone: "",
	comment: "",
}
```

### Create an invoice item

With net unit price

```go
soldItem1 := szamlazz.Item{
	label: "first item",
	quantity: 2,
	unit: "qt",
	vat: 27,
	netUnitPrice: 100.55, // calculates gross and net values from per item net
	comment: "",
}
```

With gross unit price:

```go
soldItem1 := szamlazz.Item{
	label: "first item",
	quantity: 2,
	unit: "qt",
	vat: 27,
	grossUnitPrice: 1270, // calculates gross and net values from per item net
	comment: "",
}
```

### Create an invoice

You can create an invoice with the instances created above:

```go
invoice := szamlazz.Invoice{
	paymentMethod: szamlazz.PaymentMethod.BankTransfer, // optional, default: BankTransfer
	currency: szamlazz.Currency.Ft, // optional, default: Ft
	language: szamlazz.Language.Hungarian, // optional, default: Hungarian
	seller: seller, // required
	buyer: buyer, // required
	items: []szamlazz.Item{soldItem1}
}
```

To issue the invoice with szamlazz.hu:

```go
result, err := client.issueInvoice(invoice)
if err != nil {
	log.Fatal(err)
}

if result.pdf != nil {
	// a Buffer with the pdf data is available if requestInvoiceDownload === true
}
```

## Constants

### PaymentMethod

The following payment methods are supported by szamlazz.hu:

```
szamlazz.PaymentMethod.Cash
szamlazz.PaymentMethod.BankTransfer
szamlazz.PaymentMethod.CreditCard
```

### Currency

The following currencies are recognized by szamlazz.hu:

```
szamlazz.Currency.Ft
szamlazz.Currency.HUF
szamlazz.Currency.EUR
szamlazz.Currency.CHF
szamlazz.Currency.USD
szamlazz.Currency.AUD
szamlazz.Currency.AED
szamlazz.Currency.BGN
szamlazz.Currency.CAD
szamlazz.Currency.CNY
szamlazz.Currency.CZK
szamlazz.Currency.DKK
szamlazz.Currency.EEK
szamlazz.Currency.GBP
szamlazz.Currency.HRK
szamlazz.Currency.ISK
szamlazz.Currency.JPY
szamlazz.Currency.LTL
szamlazz.Currency.LVL
szamlazz.Currency.NOK
szamlazz.Currency.NZD
szamlazz.Currency.PLN
szamlazz.Currency.RON
szamlazz.Currency.RUB
szamlazz.Currency.SEK
szamlazz.Currency.SKK
szamlazz.Currency.UAH
```

### Language

The accepted languages are:

```
szamlazz.Currency.Hungarian
szamlazz.Currency.English
szamlazz.Currency.German
szamlazz.Currency.Italian
szamlazz.Currency.Romanian
szamlazz.Currency.Slovak
```