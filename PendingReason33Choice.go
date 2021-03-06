package iso20022

// Choice between a standard code and a proprietary code to specify the reason why the instruction/event has a pending status.
type PendingReason33Choice struct {

	// Standard code to specify the reason why the instruction/event has a pending status.
	Code *PendingReason13Code `xml:"Cd"`

	// Proprietary identification of the reason why the instruction/event has a pending status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (p *PendingReason33Choice) SetCode(value string) {
	p.Code = (*PendingReason13Code)(&value)
}

func (p *PendingReason33Choice) AddProprietary() *GenericIdentification30 {
	p.Proprietary = new(GenericIdentification30)
	return p.Proprietary
}
