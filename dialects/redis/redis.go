package redis

type Config struct {
	Name       string
	Hostname   string
	Port       int16
	Database   int16
	Username   string
	Password   string
	Parameters map[string]interface{}
}
