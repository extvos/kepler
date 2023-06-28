package mysql

type Config struct {
	Name       string
	Hostname   string
	Port       int16
	Database   string
	Username   string
	Password   string
	Parameters map[string]interface{}
}
