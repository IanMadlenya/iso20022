package iso20022

// Account and holding of the next sub-level (Level 8).
type AccountSubLevel18 struct {

	// Unique and unambiguous identification for the sub-account between the account owner and the account servicer.
	AccountIdentification *SecuritiesAccount19 `xml:"AcctId"`

	// Party that legally owns the sub-account.
	AccountOwner *PartyIdentification100 `xml:"AcctOwnr"`

	// Party that manages the sub-account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification100 `xml:"AcctSvcr"`

	// Individual or entity that is ultimately entitled to the benefit of income and rights in a financial instrument, as opposed to a nominal or legal owner.
	BeneficialOwner []*BeneficialOwner2 `xml:"BnfclOwnr,omitempty"`

	// Report on the net position of a financial instrument on the sub-account (sub-account level 8),  for a certain date. The agent, for example, a trade intermediary, may also be specified.
	BalanceDetails []*AggregateHoldingBalance3 `xml:"BalDtls,omitempty"`

	// Holdings of level 9.
	AccountSubLevel9 []*AccountSubLevel19 `xml:"AcctSubLvl9,omitempty"`

	// Difference in holdings between the sub-account at level 8 and the sub-accounts of level 9.
	AccountSubLevel9Difference []*AggregateHoldingBalance2 `xml:"AcctSubLvl9Diff,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountSubLevel18) AddAccountIdentification() *SecuritiesAccount19 {
	a.AccountIdentification = new(SecuritiesAccount19)
	return a.AccountIdentification
}

func (a *AccountSubLevel18) AddAccountOwner() *PartyIdentification100 {
	a.AccountOwner = new(PartyIdentification100)
	return a.AccountOwner
}

func (a *AccountSubLevel18) AddAccountServicer() *PartyIdentification100 {
	a.AccountServicer = new(PartyIdentification100)
	return a.AccountServicer
}

func (a *AccountSubLevel18) AddBeneficialOwner() *BeneficialOwner2 {
	newValue := new(BeneficialOwner2)
	a.BeneficialOwner = append(a.BeneficialOwner, newValue)
	return newValue
}

func (a *AccountSubLevel18) AddBalanceDetails() *AggregateHoldingBalance3 {
	newValue := new(AggregateHoldingBalance3)
	a.BalanceDetails = append(a.BalanceDetails, newValue)
	return newValue
}

func (a *AccountSubLevel18) AddAccountSubLevel9() *AccountSubLevel19 {
	newValue := new(AccountSubLevel19)
	a.AccountSubLevel9 = append(a.AccountSubLevel9, newValue)
	return newValue
}

func (a *AccountSubLevel18) AddAccountSubLevel9Difference() *AggregateHoldingBalance2 {
	newValue := new(AggregateHoldingBalance2)
	a.AccountSubLevel9Difference = append(a.AccountSubLevel9Difference, newValue)
	return newValue
}

func (a *AccountSubLevel18) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
