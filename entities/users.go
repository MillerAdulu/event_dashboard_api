package entities

// UserRegistrationData -
type UserRegistrationData struct {
	Name   string  `json:"name" rethinkdb:"name"`
	Number float64 `json:"number" rethinkdb:"number"`
	IP     string  `json:"ip,omitempty" rethinkdb:"ip_address"`
}
