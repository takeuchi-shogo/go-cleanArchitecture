package infrastructure

type Config struct {
	DB struct {
		Production struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {
	c := new(Config)

	c.DB.Production.Host = "sns_mysql"
	c.DB.Production.Username = "sns_test"
	c.DB.Production.Password = "password"
	c.DB.Production.DBName = "sns_sample"

	c.Routing.Port = ":8080"

	return c
}
