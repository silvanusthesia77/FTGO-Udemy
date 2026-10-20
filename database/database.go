package database

var connection string

func init() {
	connection = "MySQL"
}
func GetDataBases() string {
	return connection
}

// Error 53
