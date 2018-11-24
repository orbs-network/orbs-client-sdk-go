package main

import (
	"github.com/orbs-network/orbs-client-sdk-go/gammacli/jsoncodec"
	"io/ioutil"
)

func getDefaultLocalConfig() *jsoncodec.ConfEnv {
	return &jsoncodec.ConfEnv{
		VirtualChain: 42,
		Endpoints:    []string{"localhost"},
	}
}

func getEnvironmentFromConfigFile(env string) *jsoncodec.ConfEnv {
	bytes, err := ioutil.ReadFile(CONFIG_FILENAME)
	if err != nil {
		if env == LOCAL_ENV_ID {
			return getDefaultLocalConfig()
		}
		die("could not open config file '%s' containing environment details\n\n%s", CONFIG_FILENAME, err.Error())
	}

	confFile, err := jsoncodec.UnmarshalConfFile(bytes)
	if err != nil {
		die("failed parsing config json file '%s'\n\n%s", CONFIG_FILENAME, err.Error())
	}

	if len(confFile.Environments) == 0 {
		if env == LOCAL_ENV_ID {
			return getDefaultLocalConfig()
		}
		die("key 'Environments' does not contain data in config file '%s'", CONFIG_FILENAME)
	}

	confEnv, found := confFile.Environments[env]
	if !found {
		if env == LOCAL_ENV_ID {
			return getDefaultLocalConfig()
		}
		die("environment with id '%s' not found in config file '%s'", env, *flagKeyFile)
	}

	return confEnv
}
