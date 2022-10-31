package cmd

const defaultConfigPath = "./clive.yml"

type configYaml struct {
	Actions []interface{}
}

type config struct {
	Actions []action
}
