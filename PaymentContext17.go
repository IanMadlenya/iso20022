package iso20022

// Payment context in which the transaction is performed.
type PaymentContext17 struct {

	// Indicates whether the transaction has been initiated by a card physically present or not.
	CardPresent *TrueFalseIndicator `xml:"CardPres,omitempty"`

	// Indicates whether the transaction has been initiated in presence of the cardholder or not.
	CardholderPresent *TrueFalseIndicator `xml:"CrdhldrPres,omitempty"`

	// On-line or off-line context of the transaction.
	OnLineContext *TrueFalseIndicator `xml:"OnLineCntxt,omitempty"`

	// Human attendance at the POI location during the transaction.
	AttendanceContext *AttendanceContext1Code `xml:"AttndncCntxt,omitempty"`

	// Indicates the environment of the transaction.
	TransactionEnvironment *TransactionEnvironment1Code `xml:"TxEnvt,omitempty"`

	// Identifies the type of the communication channels used by the cardholder to the acceptor system.
	TransactionChannel *TransactionChannel3Code `xml:"TxChanl,omitempty"`

	// Entry mode of the card data.
	CardDataEntryMode *CardDataReading5Code `xml:"CardDataNtryMd"`

	// Indicator of a card entry mode fallback.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`
}

func (p *PaymentContext17) SetCardPresent(value string) {
	p.CardPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext17) SetCardholderPresent(value string) {
	p.CardholderPresent = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext17) SetOnLineContext(value string) {
	p.OnLineContext = (*TrueFalseIndicator)(&value)
}

func (p *PaymentContext17) SetAttendanceContext(value string) {
	p.AttendanceContext = (*AttendanceContext1Code)(&value)
}

func (p *PaymentContext17) SetTransactionEnvironment(value string) {
	p.TransactionEnvironment = (*TransactionEnvironment1Code)(&value)
}

func (p *PaymentContext17) SetTransactionChannel(value string) {
	p.TransactionChannel = (*TransactionChannel3Code)(&value)
}

func (p *PaymentContext17) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading5Code)(&value)
}

func (p *PaymentContext17) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}
