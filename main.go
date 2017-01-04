package main

import (
    "os"
    "path/filepath"
    "strings"
    "encoding/json"
	_ "apiproj/routers"
    "apiproj/controllers"
	"github.com/astaxie/beego"
)

func getCurrentDirectory() string {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        beego.Error(err)
    }
    return strings.Replace(dir, "\\", "/", -1)
}

func getLogPath() string {
    path := getCurrentDirectory()
    if len(path) > 0 {
        return path + "/logs/test.log"
    }
    return "" 
}
func main() {
    logFile := getLogPath()
    logInfo, _ := json.Marshal(map[string]string{"filename":logFile})
    beego.SetLogger("file", string(logInfo))
    beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

