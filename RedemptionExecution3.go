package iso20022

// Execution of a redemption order.
type RedemptionExecution3 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType1 `xml:"OrdrTp,omitempty"`

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson2 `xml:"BnfcryDtls,omitempty"`

	// Number of investment funds units redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Net amount of money paid to the investor as a result of the redemption.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate,omitempty"`

	// Amount of money paid to the investor as a result of the redemption, including all charges, commissions, and tax.
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Price at which the order was executed.
	PriceDetails *UnitPrice5 `xml:"PricDtls"`

	// Indicates whether the order has been partially executed, ie, the confirmed quantity does not match the ordered quantity for a given financial instrument.
	PartiallyExecutedIndicator *YesNoIndicator `xml:"PrtlyExctdInd"`

	// Indicates whether the dividend is included, ie, cum-dividend, in the executed price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Part of the price deemed as accrued income or profit rather than capital. The interim profit amount is used for tax purposes.
	InterimProfitAmount *ProfitAndLoss1Choice `xml:"IntrmPrftAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms4 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Amount of money associated with a service.
	ChargeGeneralDetails *TotalCharges2 `xml:"ChrgGnlDtls,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionGeneralDetails *TotalCommissions2 `xml:"ComssnGnlDtls,omitempty"`

	// Tax related to an investment fund order.
	TaxGeneralDetails *TotalTaxes2 `xml:"TaxGnlDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters3 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Parameters of a physical delivery.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction18 `xml:"CshSttlmDtls,omitempty"`
}

func (r *RedemptionExecution3) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution3) SetDealReference(value string) {
	r.DealReference = (*Max35Text)(&value)
}

func (r *RedemptionExecution3) AddOrderType() *FundOrderType1 {
	newValue := new(FundOrderType1)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionExecution3) AddInvestmentAccountDetails() *InvestmentAccount13 {
	r.InvestmentAccountDetails = new(InvestmentAccount13)
	return r.InvestmentAccountDetails
}

func (r *RedemptionExecution3) AddBeneficiaryDetails() *IndividualPerson2 {
	r.BeneficiaryDetails = new(IndividualPerson2)
	return r.BeneficiaryDetails
}

func (r *RedemptionExecution3) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	r.UnitsNumber = new(FinancialInstrumentQuantity1)
	return r.UnitsNumber
}

func (r *RedemptionExecution3) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionExecution3) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution3) SetHoldingsRedemptionRate(value string) {
	r.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (r *RedemptionExecution3) SetGrossAmount(value, currency string) {
	r.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionExecution3) AddTradeDateTime() *DateAndDateTimeChoice {
	r.TradeDateTime = new(DateAndDateTimeChoice)
	return r.TradeDateTime
}

func (r *RedemptionExecution3) AddPriceDetails() *UnitPrice5 {
	r.PriceDetails = new(UnitPrice5)
	return r.PriceDetails
}

func (r *RedemptionExecution3) SetPartiallyExecutedIndicator(value string) {
	r.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution3) SetCumDividendIndicator(value string) {
	r.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution3) AddInterimProfitAmount() *ProfitAndLoss1Choice {
	r.InterimProfitAmount = new(ProfitAndLoss1Choice)
	return r.InterimProfitAmount
}

func (r *RedemptionExecution3) AddForeignExchangeDetails() *ForeignExchangeTerms4 {
	newValue := new(ForeignExchangeTerms4)
	r.ForeignExchangeDetails = append(r.ForeignExchangeDetails, newValue)
	return newValue
}

func (r *RedemptionExecution3) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionExecution3) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (r *RedemptionExecution3) AddChargeGeneralDetails() *TotalCharges2 {
	r.ChargeGeneralDetails = new(TotalCharges2)
	return r.ChargeGeneralDetails
}

func (r *RedemptionExecution3) AddCommissionGeneralDetails() *TotalCommissions2 {
	r.CommissionGeneralDetails = new(TotalCommissions2)
	return r.CommissionGeneralDetails
}

func (r *RedemptionExecution3) AddTaxGeneralDetails() *TotalTaxes2 {
	r.TaxGeneralDetails = new(TotalTaxes2)
	return r.TaxGeneralDetails
}

func (r *RedemptionExecution3) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionExecution3) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionExecution3) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionExecution3) AddCashSettlementDetails() *PaymentTransaction18 {
	r.CashSettlementDetails = new(PaymentTransaction18)
	return r.CashSettlementDetails
}
