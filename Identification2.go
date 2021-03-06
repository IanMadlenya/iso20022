package iso20022

// Unique identifier of a document, message or transaction.
type Identification2 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterIdentification *Max35Text `xml:"MstrId,omitempty"`

	// Identification of a basket trade.
	BasketIdentification *Max35Text `xml:"BsktId,omitempty"`

	// Reference identifying a index trade.
	IndexIdentification *Max35Text `xml:"IndxId,omitempty"`

	// Unique identifier for a list, as assigned by the trading party. The identifier must be unique within a single trading day.
	ListIdentification *Max35Text `xml:"ListId,omitempty"`

	// Program reference which identifies a program trade.
	ProgramIdentification *Max35Text `xml:"PrgmId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`
}

func (i *Identification2) SetAccountOwnerTransactionIdentification(value string) {
	i.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification2) SetAccountServicerTransactionIdentification(value string) {
	i.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification2) SetMarketInfrastructureTransactionIdentification(value string) {
	i.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (i *Identification2) SetCommonIdentification(value string) {
	i.CommonIdentification = (*Max35Text)(&value)
}

func (i *Identification2) AddTradeIdentification(value string) {
	i.TradeIdentification = append(i.TradeIdentification, (*Max35Text)(&value))
}

func (i *Identification2) SetMasterIdentification(value string) {
	i.MasterIdentification = (*Max35Text)(&value)
}

func (i *Identification2) SetBasketIdentification(value string) {
	i.BasketIdentification = (*Max35Text)(&value)
}

func (i *Identification2) SetIndexIdentification(value string) {
	i.IndexIdentification = (*Max35Text)(&value)
}

func (i *Identification2) SetListIdentification(value string) {
	i.ListIdentification = (*Max35Text)(&value)
}

func (i *Identification2) SetProgramIdentification(value string) {
	i.ProgramIdentification = (*Max35Text)(&value)
}

func (i *Identification2) SetPoolIdentification(value string) {
	i.PoolIdentification = (*Max35Text)(&value)
}

func (i *Identification2) SetCorporateActionEventIdentification(value string) {
	i.CorporateActionEventIdentification = (*Max35Text)(&value)
}
