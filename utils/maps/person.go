package maps

type Address struct {
	streetNumber int
	streetName   string
	postalCode   int
	city         string
	country      string
}

type User struct {
	firstname string
	lastname  string
	email     string
	address   Address
}
