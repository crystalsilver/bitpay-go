package client_test

import (
	. "github.com/bitpay/bitpay-go/client"
	ku "github.com/bitpay/bitpay-go/key_utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"time"
)

var _ = Describe("RetrieveInvoice", func() {
	It("Retrieves an invoice from the server with an id", func() {
		time.Sleep(5)
		pm := ku.GeneratePem()
		apiuri := os.ExpandEnv("$RCROOTADDRESS")
		webClient := Client{ApiUri: apiuri, Insecure: true, Pem: pm}
		code := os.ExpandEnv("RETRIEVEPAIR")
		token, _ := webClient.PairWithCode(code)
		webClient.Token = token
		response, err := webClient.CreateInvoice(10, "USD")
		if err != nil {
			println("the retrieve test errored while creating an invoice: Error - " + err.Error())
		}
		invoiceId := response.Id
		retrievedInvoice, err := webClient.GetInvoice(invoiceId)
		if err != nil {
			println(webClient.ApiUri + ", " + webClient.Token.token + " errored retrieving an invoice: Error - " + err.Error())
		}
		Expect(retrievedInvoice.Id).To(Equal(invoiceId))
		Expect(retrievedInvoice.Price).To(Equal(response.Price))
	})
})
