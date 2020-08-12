package servlet

type SqlConnector func(cfg Config) (SQL, error)

type RedisConnector func(cfg Config) (Redis, error)

type PublishConnector func(cfg Config) (Publisher, error)

type SubscribeConnector func(cfg Config) (Subscriber, error)
