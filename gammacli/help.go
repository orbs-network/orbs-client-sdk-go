package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func commandShowHelp() {
	fmt.Fprintf(os.Stderr, "Usage:\n\n")
	fmt.Fprintf(os.Stderr, "gamma-cli COMMAND [OPTIONS]\n\n")

	fmt.Fprintf(os.Stderr, "Commands:\n\n")
	sortedCommands := sortCommands()
	for _, name := range sortedCommands {
		cmd := commands[name]
		fmt.Fprintf(os.Stderr, "  %s %s %s\n", name, strings.Repeat(" ", 15-len(name)), cmd.desc)
		if cmd.args != "" {
			fmt.Fprintf(os.Stderr, "  %s  options: %s\n", strings.Repeat(" ", 15), cmd.args)
		}
		if cmd.example != "" {
			fmt.Fprintf(os.Stderr, "  %s  example: %s\n", strings.Repeat(" ", 15), cmd.example)
		}
		fmt.Fprintf(os.Stderr, "\n")
	}
	fmt.Fprintf(os.Stderr, "\n")

	fmt.Fprintf(os.Stderr, "Options:\n\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")

	fmt.Fprintf(os.Stderr, "Multiple environments (eg. local and testnet) can be defined in orbs-gamma-config.json configuration file.\n")
	fmt.Fprintf(os.Stderr, "See https://github.com/orbs-network/orbs-contract-sdk for more info.\n")
	fmt.Fprintf(os.Stderr, "\n")

	os.Exit(2)
}

func commandVersion() {
	log("gamma-cli version %s", GAMMA_CLI_VERSION)
	// log("gamma server version %s", "TODO")
}

func sortCommands() []string {
	res := make([]string, len(commands))
	for name, cmd := range commands {
		res[cmd.sort] = name
	}
	return res
}
