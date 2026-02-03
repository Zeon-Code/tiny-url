package config

type Configuration interface {
	Log() Log
	Cache() DatabaseConfiguration
	PrimaryDatabase() DatabaseConfiguration
	ReplicaDatabase() DatabaseConfiguration
}

type AppConfiguration struct{}

func NewConfiguration() Configuration {
	return AppConfiguration{}
}

func (c AppConfiguration) Cache() DatabaseConfiguration {
	return RedisConfig{}
}

func (c AppConfiguration) PrimaryDatabase() DatabaseConfiguration {
	return newPostgresConfig("DB")
}

func (c AppConfiguration) ReplicaDatabase() DatabaseConfiguration {
	return newPostgresConfig("DB_REPLICA")
}

func (c AppConfiguration) Log() Log {
	return newLogConfig()
}
