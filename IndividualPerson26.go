package iso20022

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson26 struct {

	// Name received at birth, for example, maiden name.
	BirthName *Max35Text `xml:"BirthNm"`

	// First name of the person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Unique and unambiguous identification of the person, for example, passport.
	Identification *PersonIdentification6 `xml:"Id,omitempty"`

	// Postal address of the party.
	Address *LongPostalAddress2Choice `xml:"Adr,omitempty"`

	// Organisation represented by the person, or for which the person works.
	EmployingParty *PartyIdentification40Choice `xml:"EmplngPty,omitempty"`

	// Specifies details related to the attendance card.
	AttendanceCardDetails *AttendanceCard2 `xml:"AttndncCardDtls"`
}

func (i *IndividualPerson26) SetBirthName(value string) {
	i.BirthName = (*Max35Text)(&value)
}

func (i *IndividualPerson26) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson26) AddIdentification() *PersonIdentification6 {
	i.Identification = new(PersonIdentification6)
	return i.Identification
}

func (i *IndividualPerson26) AddAddress() *LongPostalAddress2Choice {
	i.Address = new(LongPostalAddress2Choice)
	return i.Address
}

func (i *IndividualPerson26) AddEmployingParty() *PartyIdentification40Choice {
	i.EmployingParty = new(PartyIdentification40Choice)
	return i.EmployingParty
}

func (i *IndividualPerson26) AddAttendanceCardDetails() *AttendanceCard2 {
	i.AttendanceCardDetails = new(AttendanceCard2)
	return i.AttendanceCardDetails
}
