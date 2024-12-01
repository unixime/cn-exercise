package main

import (
	"cn-exercise/cmd"
	"cn-exercise/internal/client"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func init() {

	pflag.String("url", "https://vault.immudb.io/", "")
	pflag.String("api-key", "", "")

	viper.SetEnvPrefix("CN")
	viper.AutomaticEnv()
	viper.BindPFlags(pflag.CommandLine)
	pflag.Parse()

}

func main() {

	fmt.Println(os.Getenv("CN_URL"))
	fmt.Println(viper.Get("url"))
	cmd.StartServer(client.Client{
		URL:    "https://vault.immudb.io",
		ApiKey: "default.AIkWyayo4M8uOBVUbce3zg.DyHDJbEg9chloDI6deZ2ldxERsi_z-fxifUqgkNuzsH5TZ3y",
	})
}
