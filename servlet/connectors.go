package servlet

// SqlConnector
// SQL Database Connector function definition
type SqlConnector func(cfg Config) (SQL, error)

// RedisConnector
// Redis Connector function definition
type RedisConnector func(cfg Config) (Redis, error)

// PublishConnector
// Publisher Connector function definition
type PublishConnector func(cfg Config) (Publisher, error)

// SubscribeConnector
// Subscriber Connector function definition
type SubscribeConnector func(cfg Config) (Subscriber, error)
