package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command

var loginCmd = &cobra.Command{
	Use:   "login {$servername}",
	Short: "Login to the given SensorThings server",
	Run: func(cmd *cobra.Command, args []string) {
		err := RunLogin(cmd, args)
		if err != nil {
			exitWithError(err)
		}

	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
}

// RunLogin set the login
func RunLogin(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return cmd.Help()
	}

	validURL := govalidator.IsURL(args[0])

	if !validURL {
		exitWithError(fmt.Errorf("Not a valid url: %s", args[0]))
	}

	serverURL := args[0]
	configPath := ""
	configName := ".sensorthings-cli"
	configType := "yaml"
	var configYaml []byte
	var madeConfigFile = false
	configFile := path.Join(configPath,
		(configName + "." + configType))

	if _, err := os.Stat(configFile); err != nil {
		var file, err = os.Create(configFile)
		defer file.Close()
		if err != nil {
			exitWithError(fmt.Errorf("Could not create config: %s", configFile))
		}
		configYaml = []byte("st_server: " + serverURL + "\n")

		defer file.Close()
		madeConfigFile = true
	} else {
		input, err := ioutil.ReadFile(configFile)
		if err != nil {
			exitWithError(err)
		}

		lines := strings.Split(string(input), "\n")

		isUpdate := false
		gitLine := "st_server: " + serverURL
		for i, line := range lines {
			if strings.Contains(line, "st_server: ") {
				isUpdate = true
				lines[i] = gitLine
			}
		}
		output := strings.Join(lines, "\n")
		if !isUpdate {
			output = output + "\n" + gitLine + "\n"
		}
		configYaml = []byte(output)
		if madeConfigFile {
			err_del := os.Remove(configFile)
			if err_del != nil {
				exitWithError(fmt.Errorf("Could not delete config: %s", configFile))
			}
		}

	}
	err := ioutil.WriteFile(configFile, configYaml, 0644)
	if err != nil {
		exitWithError(fmt.Errorf("Could not write config to %s", configFile))
	}

	if madeConfigFile {
		fmt.Printf("Login Succeeded")
	} else {
		fmt.Printf("Login Succeeded")
	}
	return nil
}

// exitWithError will terminate execution with an error result
// It prints the error to stderr and exits with a non-zero exit code
func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}
