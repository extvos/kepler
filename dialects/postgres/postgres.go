package postgres

type Config struct {
	Name       string
	Hostname   string
	Port       int16
	Database   string
	Schema     string
	Username   string
	Password   string
	Parameters map[string]interface{}
}

type Postgres struct {
}
