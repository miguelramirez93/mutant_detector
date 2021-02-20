package contracts

// DbProvider represents the methods that every db provider should have.
type DbProvider interface {
	Connect() (interface{}, error)
	Migrate(connInstance interface{}) error
}

// DbDriverManager epresents the methods that every db driver manager should have.
type DbDriverManager interface {
	GetProvider(driverName string) DbProvider
}
