package config

func InitConfig() error {
	errEnv := InitEnv()
	if errEnv != nil {
		return errEnv
	}

	errDatabase := InitDatabase()
	if errDatabase != nil {
		return errDatabase
	}

	return nil
}
