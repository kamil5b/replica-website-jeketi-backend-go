package constant

const (
	EnvDatabaseUrl     = "DATABASE_URL"
	EnvPort            = "PORT"
	EnvTokenExpiryHour = "TOKEN_EXPIRY_HOUR"
)

var (
	EnvDefaultValues = map[string]string{
		EnvPort:            "3000",
		EnvTokenExpiryHour: "72",
	}
)
