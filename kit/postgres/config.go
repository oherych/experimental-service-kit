package postgres

type Config struct {
	DSN string `mapstructure:"MAIN_POSTGRES_URL"`
}
