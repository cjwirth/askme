package server

type Config struct {

	// Database
	DBDriver     string
	DBDataSource string

	// JSON
	JSONPrettyPrint bool
}
