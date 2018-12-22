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
	bytes, err := ioutil.ReadFile(*flagConfigFile)
	if err != nil {
		if env == LOCAL_ENV_ID {
			return getDefaultLocalConfig()
		}
		die("Could not open config file '%s' containing environment details.\n\n%s", *flagConfigFile, err.Error())
	}

	confFile, err := jsoncodec.UnmarshalConfFile(bytes)
	if err != nil {
		die("Failed parsing config json file '%s'.\n\n%s", *flagConfigFile, err.Error())
	}

	if len(confFile.Environments) == 0 {
		if env == LOCAL_ENV_ID {
			return getDefaultLocalConfig()
		}
		die("Key 'Environments' does not contain data in config file '%s'.", *flagConfigFile)
	}

	confEnv, found := confFile.Environments[env]
	if !found {
		if env == LOCAL_ENV_ID {
			return getDefaultLocalConfig()
		}
		die("Environment with id '%s' not found in config file '%s'.", env, *flagKeyFile)
	}

	return confEnv
}
