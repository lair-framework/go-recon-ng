package reconng

import "encoding/json"

// Results contains all the data for a Recon-ng output
type Results struct {
	Contacts    []Contact    `json:"contacts"`
	NetBlocks   []Netblock   `json:"netblocks"`
	Domains     []Domain     `json:"domains"`
	Credentials []Credential `json:"credentials"`
	Hosts       []Host       `json:"hosts"`
}

// Contact information
type Contact struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Title      string `json:"title"`
	Country    string `json:"country"`
	Region     string `json:"region"`
	Module     string `json:"module"`
	Email      string `json:"email"`
}

// Netblock information
type Netblock struct {
	Netblock  string `json:"netblock"`
	Module    string `json:"module"`
	Email     string `json:"email"`
	OrgHandle string `json:"org_handle"`
}

// Domain information
type Domain struct {
	Name   string `json:"domain"`
	Module string `json:"module"`
}

// Credential from leaks
type Credential struct {
	Username string `json:"username"`
	Hash     string `json:"hash"`
	Leak     string `json:"leak"`
	Module   string `json:"module"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

// Host information
type Host struct {
	Country   string `json:"country"`
	Region    string `json:"region"`
	Longitude string `json:"longitude"`
	Name      string `json:"host"`
	Module    string `json:"module"`
	Lattitude string `json:"lattitude"`
	IPAddress string `json:"ip_address"`
}

// Parse recon-ng json data
func Parse(content []byte) (*Results, error) {
	r := &Results{}
	err := json.Unmarshal(content, r)
	if err != nil {
		return r, err
	}
	return r, nil
}
