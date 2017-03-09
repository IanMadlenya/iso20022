package iso20022

// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
type Header31 struct {

	// Identifies the type of process related to the message.
	MessageFunction *ATMMessageFunction2 `xml:"MsgFctn"`

	// Version of the ATM protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *Max35Text `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *Max35Text `xml:"RcptPty,omitempty"`

	// State of the sender of the message inside the process flow.
	ProcessState *Max35Text `xml:"PrcStat,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability4 `xml:"Tracblt,omitempty"`
}

func (h *Header31) AddMessageFunction() *ATMMessageFunction2 {
	h.MessageFunction = new(ATMMessageFunction2)
	return h.MessageFunction
}

func (h *Header31) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header31) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header31) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header31) SetInitiatingParty(value string) {
	h.InitiatingParty = (*Max35Text)(&value)
}

func (h *Header31) SetRecipientParty(value string) {
	h.RecipientParty = (*Max35Text)(&value)
}

func (h *Header31) SetProcessState(value string) {
	h.ProcessState = (*Max35Text)(&value)
}

func (h *Header31) AddTraceability() *Traceability4 {
	newValue := new(Traceability4)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}
