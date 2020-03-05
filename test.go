// package main

// import (
// 	"flag"
// 	"fmt"

// 	"github.com/somewhere/db"
// 	"github.com/somewhere/model"

// 	"github.com/somewhere/config"
// 	"github.com/somewhere/utils"
// )

// func main() {
// 	var configPath string
// 	flag.StringVar(&configPath, "config", "./conf/config.toml", "config path")
// 	flag.Parse()
// 	config.InitConfig(configPath)
// 	db.InitDatabase()
// 	ans, err := utils.GetItemScoreFromUserID("5df9e1fe91560048ad6fb730")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(ans)
// 	plist, err := model.GetAllProducts()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	ans2 := utils.SortItemByScore(ans, plist)
// 	fmt.Println(ans2)
// }
