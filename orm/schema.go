package orm

type AddressInfo struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Address   string `sql:"not null;unique"`
}
