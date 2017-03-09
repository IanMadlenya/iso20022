package acmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02100101 struct {
	XMLName xml.Name                                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.021.001.01 Document"`
	Message *AccountClosingAdditionalInformationRequestV01 `xml:"AcctClsgAddtlInfReq"`
}

func (d *Document02100101) AddMessage() *AccountClosingAdditionalInformationRequestV01 {
	d.Message = new(AccountClosingAdditionalInformationRequestV01)
	return d.Message
}

// Scope
// The AccountClosingAdditionalInformationRequest message is sent from a financial institution to an organisation as part of the account closing process.
// Usage
// This message is sent in response to an closing request message from the organisation, if the business content is valid, but additional information is required.
type AccountClosingAdditionalInformationRequestV01 struct {

	// Set of elements for the identification of the message and related references.
	References *iso20022.References3 `xml:"Refs"`

	// Identifier for an organisation.
	OrganisationIdentification []*iso20022.OrganisationIdentification6 `xml:"OrgId"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	AccountIdentification *iso20022.AccountForAction1 `xml:"AcctId"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	//
	AccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification4 `xml:"AcctSvcrId"`

	// Identification of the account to which the remaining positive balance of the account to be closed must be transferred or account from which funds can be moved to the account to be closed and which balance is negative. This account must be held in the same financial institution as the account to be closed if the transfer account is used to compensate a negative balance. For a positive balance to be transferred, an account in another financial institution might be used. In that case the account servicer is mandatory.
	BalanceTransferAccount *iso20022.AccountForAction1 `xml:"BalTrfAcct,omitempty"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme, that is the servicer of the transfer account.
	TransferAccountServicerIdentification *iso20022.BranchAndFinancialInstitutionIdentification4 `xml:"TrfAcctSvcrId,omitempty"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*iso20022.PartyAndSignature1 `xml:"DgtlSgntr,omitempty"`
}

func (a *AccountClosingAdditionalInformationRequestV01) AddReferences() *iso20022.References3 {
	a.References = new(iso20022.References3)
	return a.References
}

func (a *AccountClosingAdditionalInformationRequestV01) AddOrganisationIdentification() *iso20022.OrganisationIdentification6 {
	newValue := new(iso20022.OrganisationIdentification6)
	a.OrganisationIdentification = append(a.OrganisationIdentification, newValue)
	return newValue
}

func (a *AccountClosingAdditionalInformationRequestV01) AddAccountIdentification() *iso20022.AccountForAction1 {
	a.AccountIdentification = new(iso20022.AccountForAction1)
	return a.AccountIdentification
}

func (a *AccountClosingAdditionalInformationRequestV01) AddAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification4 {
	a.AccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification4)
	return a.AccountServicerIdentification
}

func (a *AccountClosingAdditionalInformationRequestV01) AddBalanceTransferAccount() *iso20022.AccountForAction1 {
	a.BalanceTransferAccount = new(iso20022.AccountForAction1)
	return a.BalanceTransferAccount
}

func (a *AccountClosingAdditionalInformationRequestV01) AddTransferAccountServicerIdentification() *iso20022.BranchAndFinancialInstitutionIdentification4 {
	a.TransferAccountServicerIdentification = new(iso20022.BranchAndFinancialInstitutionIdentification4)
	return a.TransferAccountServicerIdentification
}

func (a *AccountClosingAdditionalInformationRequestV01) AddDigitalSignature() *iso20022.PartyAndSignature1 {
	newValue := new(iso20022.PartyAndSignature1)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}
