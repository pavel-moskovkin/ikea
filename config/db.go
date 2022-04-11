package config

type DB struct {
	Address  string
	Username string
	Password string
	DBName   string
	Insecure bool
}

func NewDB() DB {
	return DB{
		Address:  "localhost:5432",
		Username: "root",
		Password: "toor",
		DBName:   "postgres",
		Insecure: true,
	}
}
