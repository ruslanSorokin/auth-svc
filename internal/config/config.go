package config

type (
	// DocTable stores information for document oriented database's tables
	DocTable struct {
		URI   string
		DB    string
		Table string
	}

	// Handler stores information for network exposure
	Handler struct {
		Host string
		Port int
	}
)
