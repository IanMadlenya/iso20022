package iso20022

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification25 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Exact4AlphaNumericText `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max4AlphaNumericText `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification25) SetIdentification(value string) {
	g.Identification = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification25) SetIssuer(value string) {
	g.Issuer = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification25) SetSchemeName(value string) {
	g.SchemeName = (*Max4AlphaNumericText)(&value)
}
