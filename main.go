package main

import (
	_ "cloudauto/routers"
	"cloudauto/work"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"path/filepath"
)

func main() {
	beego.Run()
	data := "/data/apps/mysql"
	fmt.Println(work.Version)
	fmt.Println(filepath.Join(data, "./static"))
}
