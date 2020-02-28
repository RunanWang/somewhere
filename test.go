// package main

// import (
// 	"flag"
// 	"fmt"

// 	"github.com/somewhere/config"
// 	"github.com/somewhere/utils"
// )

// func main() {
// 	var configPath string
// 	flag.StringVar(&configPath, "config", "./conf/config.toml", "config path")
// 	flag.Parse()
// 	config.InitConfig(configPath)
// 	ans, err := utils.GetItemScoreFromUserID("5df9e1fe91560048ad6fb730")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(ans)
// }
