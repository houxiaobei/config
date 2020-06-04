package config

// An Option configures a Config.
type Option interface {
	apply(*Config)
}

// optionFunc wraps a func so it satisfies the Option interface.
type optionFunc func(*Config)

func (f optionFunc) apply(cfg *Config) {
	f(cfg)
}

func FromFile() Option {
	return optionFunc(func(cfg *Config) {
		cfg.drivers = append(cfg.drivers, &FileDrive{})
	})
}
