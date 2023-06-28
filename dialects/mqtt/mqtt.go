package mqtt

type Config struct {
	Name       string
	Broker     string
	ClientId   string
	Username   string
	Password   string
	Parameters map[string]interface{}
}
