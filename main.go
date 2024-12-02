package main

import (
	"cn-exercise/cmd"
	"cn-exercise/internal/client"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

func init() {

	pflag.String("url", "https://vault.immudb.io", "")
	pflag.String("api-key", "", "")

	viper.SetEnvPrefix("CN")
	viper.AutomaticEnv()
	if e := viper.BindPFlags(pflag.CommandLine); e != nil {
		fmt.Printf("%s: url", e.Error())
		os.Exit(3)
	}
	pflag.Parse()

}

func getParamValue(userInput, envVar string) (string, error) {
	ui := viper.GetString(userInput)
	if ui != "" {
		return ui, nil
	}

	ev := os.Getenv(envVar)
	if ev != "" {
		return ev, nil
	}

	return "", fmt.Errorf("missing parameter")
}

func main() {

	url, e1 := getParamValue("url", "CN_URL")
	if e1 != nil {
		fmt.Printf("%s: url", e1.Error())
		os.Exit(1)
	}

	ak, e2 := getParamValue("api-key", "CN_API_KEY")
	if e2 != nil {
		fmt.Printf("%s: api-key", e2.Error())
		os.Exit(2)
	}

	if err := cmd.StartServer(client.Client{
		URL:    url,
		ApiKey: ak,
	}); err != nil {
		fmt.Printf("%s: api-key", err.Error())
		os.Exit(4)
	}
}
