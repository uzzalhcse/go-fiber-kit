package config

import "github.com/spf13/viper"

type AmadeusConfig struct {
	APIKey    string
	APISecret string
}

func loadAmadeusConfig() AmadeusConfig {
	return AmadeusConfig{
		APIKey:    viper.GetString("AMADEUS_API_KEY"),
		APISecret: viper.GetString("AMADEUS_API_SECRET"),
	}
}
