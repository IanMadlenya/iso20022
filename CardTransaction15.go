package iso20022

// Card transaction for which the authorisation is requested.
type CardTransaction15 struct {

	// Type of transaction being undertaken for the main service.
	// It correspond partially to the ISO 8583 field number 3.
	TransactionType *CardPaymentServiceType7Code `xml:"TxTp"`

	// Service in addition to the main service.
	// It correspond partially to the ISO 8583:2003 field number 22-3.
	AdditionalService []*CardPaymentServiceType8Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	// It correspond partially to the ISO 8583:2003 field number 22-3.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	// It correspond to the ISO 8583 field number 18 for the version 87, field numbers 18 and 26 for the version 93, and field number 26 for the version 2003.
	MerchantCategoryCode *Min3Max4NumericText `xml:"MrchntCtgyCd,omitempty"`

	// Identification of the reconciliation period between the acquirer and the issuer or their respective agents.
	Reconciliation *TransactionIdentifier2 `xml:"Rcncltn,omitempty"`

	// Date and time of the transaction transported unchanged by the acquirer from the card acceptor to the issuer. Corresponds to the CAPE data element TransactionIdentification/ TransactionDateTime.
	// It correspond to the ISO 8583 field number 12.
	AcceptorTransactionDateTime *ISODateTime `xml:"AccptrTxDtTm"`

	// Unique transaction identification generated by the acceptor or the acquirer transported unchanged by the acquirer from the card acceptor to the issuer. It is used to assist locating the original source information. Eventually it could be included in the cardholder statement. It corresponds to ISO 8583, field 37 and CAPE data element TransactionIdentification/TransactionReference.
	AcceptorTransactionIdentification *Max35Text `xml:"AccptrTxId"`

	// Number generated by the transaction Initiator to assist in identifying a transaction uniquely. This value remains unchanged for all messages within a message pair exchange, for instance an initiation/response. It corresponds to the ISO 8583 field number 11.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId"`

	// Unique identification to match transactions throughout their life cycle (for example, authorisation to financial presentment, financial presentment to chargeback). It shall contain the same value in all messages throughout a transaction’s life cycle. It corresponds partially to ISO 8583:2003 field number 21.
	TransactionLifeCycleIdentification *Max35Text `xml:"TxLifeCyclId,omitempty"`

	// Life cycle transaction sequence number when multiple authorisations are performed for the same presentment.
	// It corresponds partially to ISO 8583:2003, field number 21.
	TransactionLifeCycleSequenceNumber *Number `xml:"TxLifeCyclSeqNb,omitempty"`

	// Total number of transactions under the same life cycle transaction sequence number.
	TransactionLifeCycleSequenceCounter *Number `xml:"TxLifeCyclSeqCntr,omitempty"`

	// Identification of the original transaction.
	// It corresponds to ISO 8583 field number 90 for the version 87, and field number 56 for the other versions.
	OriginalTransaction *CardTransaction3 `xml:"OrgnlTx,omitempty"`

	// Details of the card transaction.
	TransactionDetails *CardTransactionDetail1 `xml:"TxDtls"`

	// Outcome of the authorisation.
	AuthorisationResult *AuthorisationResult7 `xml:"AuthstnRslt,omitempty"`
}

func (c *CardTransaction15) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType7Code)(&value)
}

func (c *CardTransaction15) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType8Code)(&value))
}

func (c *CardTransaction15) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardTransaction15) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4NumericText)(&value)
}

func (c *CardTransaction15) AddReconciliation() *TransactionIdentifier2 {
	c.Reconciliation = new(TransactionIdentifier2)
	return c.Reconciliation
}

func (c *CardTransaction15) SetAcceptorTransactionDateTime(value string) {
	c.AcceptorTransactionDateTime = (*ISODateTime)(&value)
}

func (c *CardTransaction15) SetAcceptorTransactionIdentification(value string) {
	c.AcceptorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardTransaction15) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardTransaction15) SetTransactionLifeCycleIdentification(value string) {
	c.TransactionLifeCycleIdentification = (*Max35Text)(&value)
}

func (c *CardTransaction15) SetTransactionLifeCycleSequenceNumber(value string) {
	c.TransactionLifeCycleSequenceNumber = (*Number)(&value)
}

func (c *CardTransaction15) SetTransactionLifeCycleSequenceCounter(value string) {
	c.TransactionLifeCycleSequenceCounter = (*Number)(&value)
}

func (c *CardTransaction15) AddOriginalTransaction() *CardTransaction3 {
	c.OriginalTransaction = new(CardTransaction3)
	return c.OriginalTransaction
}

func (c *CardTransaction15) AddTransactionDetails() *CardTransactionDetail1 {
	c.TransactionDetails = new(CardTransactionDetail1)
	return c.TransactionDetails
}

func (c *CardTransaction15) AddAuthorisationResult() *AuthorisationResult7 {
	c.AuthorisationResult = new(AuthorisationResult7)
	return c.AuthorisationResult
}
