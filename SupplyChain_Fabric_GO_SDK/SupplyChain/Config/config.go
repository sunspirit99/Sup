package Config

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath("/home/sunspirit/Desktop/se06-24.4/SupplyChain_Fabric_GO_SDK/SupplyChain/Config")
	err := config.ReadInConfig()
	if err != nil {
		fmt.Println("config not found !", err)
		return nil
	}
	// fmt.Println(config.GetString("ServerName"))
	return config
}
