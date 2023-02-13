package configs

var Config envConfigs

func InitializeConfigs() {
	initializeEnvConfigs()

	// Can add more config files here
	Config = envConfig

}
