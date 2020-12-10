package pokey

type Subject struct {
	CommonName       string `json:"common_name,omitempty"`
	Country          string `json:"country,omitempty"`
	State            string `json:"state,omitempty"`
	Locality         string `json:"locality,omitempty"`
	Organization     string `json:"organization,omitempty"`
	OrganizationUnit string `json:"organization_unit,omitempty"`
}
