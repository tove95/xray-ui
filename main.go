package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	_ "unsafe"
	"xray-ui/config"
	"xray-ui/database"
	"xray-ui/logger"
	"xray-ui/v2ui"
	"xray-ui/web"
	"xray-ui/web/global"
	"xray-ui/web/service"
	"github.com/op/go-logging"
)

func runWebServer() {
	log.Printf("%v %v", config.GetName(), config.GetVersion())

	switch config.GetLogLevel() {
	case config.Debug:
		logger.InitLogger(logging.DEBUG)
	case config.Info:
		logger.InitLogger(logging.INFO)
	case config.Warn:
		logger.InitLogger(logging.WARNING)
	case config.Error:
		logger.InitLogger(logging.ERROR)
	default:
		log.Fatal("unknown log level:", config.GetLogLevel())
	}

	err := database.InitDB(config.GetDBPath())
	if err != nil {
		log.Fatal(err)
	}

	var server *web.Server

	server = web.NewServer()
	global.SetWebServer(server)
	err = server.Start()
	if err != nil {
		log.Println(err)
		return
	}

	sigCh := make(chan os.Signal, 1)
	//信号量捕获处理
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGKILL)
	for {
		sig := <-sigCh

		switch sig {
		case syscall.SIGHUP:
			err := server.Stop()
			if err != nil {
				logger.Warning("stop server err:", err)
			}
			server = web.NewServer()
			global.SetWebServer(server)
			err = server.Start()
			if err != nil {
				log.Println(err)
				return
			}
		default:
			server.Stop()
			return
		}
	}
}

func resetSetting() {
	err := database.InitDB(config.GetDBPath())
	if err != nil {
		fmt.Println(err)
		return
	}

	settingService := service.SettingService{}
	err = settingService.ResetSettings()
	if err != nil {
		fmt.Println("reset setting failed:", err)
	} else {
		fmt.Println("reset setting success")
	}
}

func showSetting(show bool) {
	if show {
		settingService := service.SettingService{}
		port, err := settingService.GetPort()
		if err != nil {
			fmt.Println("get current port fialed,error info:", err)
		}
		listen, err := settingService.GetListen()
		if err != nil {
			fmt.Println("get current listen fialed,error info:", err)
		}
		webBasePath, err := settingService.GetBasePath()
		if err != nil {
			fmt.Println("get webBasePath failed, error info:", err)
		}
		GetCertFile, err := settingService.GetCertFile()
		if err != nil {
			fmt.Println("get GetCertFile failed, error info:", err)
		}
		GetKeyFile, err := settingService.GetKeyFile()
		if err != nil {
			fmt.Println("get GetKeyFile failed, error info:", err)
		}
		GetCaFile, err := settingService.GetCaFile()
		if err != nil {
			fmt.Println("get GetCaFile failed, error info:", err)
		}
		userService := service.UserService{}
		userModel, err := userService.GetFirstUser()
		if err != nil {
			fmt.Println("get current user info failed,error info:", err)
		}
		username := userModel.Username
		userpasswd := userModel.Password
		if (username == "") || (userpasswd == "") {
			fmt.Println("current username or password is empty")
		}
		fmt.Println("当前面板信息设置如下:")
		fmt.Println("登录用户名:", username)
		fmt.Println("登录密码:", userpasswd)
		fmt.Println("登录端口:", port)
		fmt.Println("监听地址:", listen)
		if webBasePath != "" {
			fmt.Println("Web 路径:", webBasePath)
		} else {
			fmt.Println("Web 路径 is not set")
		}
		if GetCertFile != "" {
			fmt.Println("证书文件:", GetCertFile)
			fmt.Println("私钥文件:", GetKeyFile)
		} else {
			fmt.Println("证书 is not set")
		}
		if GetCaFile != "" {
			fmt.Println("CA mTLS 开启:", GetCaFile)
		} else {
			fmt.Println("CA is not set")
		}

	}
}
 
func updateTgbotEnableSts(status bool) {
	settingService := service.SettingService{}
	currentTgSts, err := settingService.GetTgbotenabled()
	if err != nil {
		fmt.Println(err)
		return
	}
	logger.Infof("current enabletgbot status[%v],need update to status[%v]", currentTgSts, status)
	if currentTgSts != status {
		err := settingService.SetTgbotenabled(status)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			logger.Infof("SetTgbotenabled[%v] success", status)
		}
	}
	return
}

func updateTgbotSetting(tgBotToken string, tgBotChatid int, tgBotRuntime string) {
	err := database.InitDB(config.GetDBPath())
	if err != nil {
		fmt.Println(err)
		return
	}

	settingService := service.SettingService{}

	if tgBotToken != "" {
		err := settingService.SetTgBotToken(tgBotToken)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			logger.Info("updateTgbotSetting tgBotToken success")
		}
	}

	if tgBotRuntime != "" {
		err := settingService.SetTgbotRuntime(tgBotRuntime)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			logger.Infof("updateTgbotSetting tgBotRuntime[%s] success", tgBotRuntime)
		}
	}

	if tgBotChatid != 0 {
		err := settingService.SetTgBotChatId(tgBotChatid)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			logger.Info("updateTgbotSetting tgBotChatid success")
		}
	}
}

func updateSetting(port int, username string, password string, listen  string, webBasePath string) {
	err := database.InitDB(config.GetDBPath())
	if err != nil {
		fmt.Println(err)
		return
	}
 
	settingService := service.SettingService{}
 
	if port > 0 {
		err := settingService.SetPort(port)
		if err != nil {
			fmt.Println("set port failed:", err)
		} else {
			fmt.Printf("set port %v success", port)
		}
	}
	if username != "" || password != "" {
		userService := service.UserService{}
		err := userService.UpdateFirstUser(username, password)
		if err != nil {
			fmt.Println("set username and password failed:", err)
		} else {
			fmt.Println("set username and password success")
		}
	}
	if listen != "" {
		err := settingService.SetListen(listen)
		if err != nil {
			fmt.Println("set listen failed:", err)
		} else {
			fmt.Printf("set listen %v success", listen)
		}		 
	}
	if webBasePath != "" {
		err := settingService.SetBasePath(webBasePath)
		if err != nil {
			fmt.Println("Failed to set base URI path:", err)
		} else {
			fmt.Println("Base URI path set successfully")
		}
	}
}

func UpdateAllip() {

	err := database.InitDB(config.GetDBPath())
	if err != nil {
		fmt.Println(err)
		return
	}	

	serverService := service.ServerService{} // 创建 ServerService 实例

	version, err := serverService.GetLatestVersion()
	if err != nil {
		fmt.Printf("Error getting the latest version: %v\n", err)
		return
	}

	err = serverService.UpdateGeoipip(version)
	if err != nil {
		fmt.Printf("Error updating GeoIP file for version %s: %v\n", version, err)
		return
	}

	err = serverService.UpdateGeositeip(version)
	if err != nil {
		fmt.Printf("Error updating Geosite file for version %s: %v\n", version, err)
		return
	}

	fmt.Printf("GeoIP and Geosite files for version %s downloaded and updated successfully!\n", version)
	
	GeoipVersion := service.GeoipVersion{}
	err = GeoipVersion.UpVersion(version)
	if err != nil {
		fmt.Println("get current UpVersion failed,error info:", err)

	}
}

func updateCert(publicKey string, privateKey string, clientCa string) {
	err := database.InitDB(config.GetDBPath())
	if err != nil {
		fmt.Println(err)
		return
	}

	if (privateKey != "" && publicKey != "") || (privateKey == "" && publicKey == "") {
		settingService := service.SettingService{}
		err = settingService.SetCertFile(publicKey)
		if err != nil {
			fmt.Println("set certificate public key failed:", err)
		} else {
			fmt.Println("set certificate public key success")
		}

		err = settingService.SetKeyFile(privateKey)
		if err != nil {
			fmt.Println("set certificate private key failed:", err)
		} else {
			fmt.Println("set certificate private key success")
		}
	} else {
		fmt.Println("both public and private key should be entered.")
	}

	if (clientCa != "") || (clientCa == "") {
		settingService := service.SettingService{}
		err = settingService.SetCaFile(clientCa)
		if err != nil {
			fmt.Println("set mTLS ca failed:", err)
		} else {
			fmt.Println("set mTLS ca success")
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		runWebServer()
		return
	}

	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")

	runCmd := flag.NewFlagSet("run", flag.ExitOnError)

	v2uiCmd := flag.NewFlagSet("v2-ui", flag.ExitOnError)
	var dbPath string
	v2uiCmd.StringVar(&dbPath, "db", "/etc/v2-ui/v2-ui.db", "set v2-ui db file path")

	settingCmd := flag.NewFlagSet("setting", flag.ExitOnError)

	geoipCmd := flag.NewFlagSet("geoip", flag.ExitOnError)

	var port int
	var listen string
	var username string
	var password string
	var webBasePath string
	var webCertFile string
	var webKeyFile string
	var webCAFile string
	var tgbottoken string
	var tgbotchatid int
	var enabletgbot bool
	var tgbotRuntime string
	var reset bool
	var show bool
	settingCmd.BoolVar(&reset, "reset", false, "reset all settings")
	settingCmd.BoolVar(&show, "show", false, "show current settings")
	settingCmd.IntVar(&port, "port", 0, "set panel port")
	settingCmd.StringVar(&listen, "listen", "", "set panel listen")
	settingCmd.StringVar(&username, "username", "", "set login username")
	settingCmd.StringVar(&password, "password", "", "set login password")
	settingCmd.StringVar(&webBasePath, "webBasePath", "", "Set base path for Panel")
	settingCmd.StringVar(&webCertFile, "webCert", "", "Set path to public key file for panel")
	settingCmd.StringVar(&webKeyFile, "webCertKey", "", "Set path to private key file for panel")
	settingCmd.StringVar(&webCAFile, "webCa", "", "Set path to mTLS ca file for panel")
	settingCmd.StringVar(&tgbottoken, "tgbottoken", "", "set telegrame bot token")
	settingCmd.StringVar(&tgbotRuntime, "tgbotRuntime", "", "set telegrame bot cron time")
	settingCmd.IntVar(&tgbotchatid, "tgbotchatid", 0, "set telegrame bot chat id")
	settingCmd.BoolVar(&enabletgbot, "enabletgbot", false, "enable telegram bot notify")
 
	oldUsage := flag.Usage
	flag.Usage = func() {
		oldUsage()
		fmt.Println()
		fmt.Println("Commands:")
		fmt.Println("    run            run web panel")
		fmt.Println("    v2-ui          migrate form v2-ui")
		fmt.Println("    setting        set settings")
		fmt.Println("    geoip          down geoip and geosite")
	}

	flag.Parse()
	if showVersion {
		fmt.Println(config.GetVersion())
		return
	}

	switch os.Args[1] {
	case "run":
		err := runCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			return
		}
		runWebServer()
	case "v2-ui":
		err := v2uiCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			return
		}
		err = v2ui.MigrateFromV2UI(dbPath)
		if err != nil {
			fmt.Println("migrate from v2-ui failed:", err)
		}
	case "setting":
		err := settingCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			return
		}
		if reset {
			resetSetting()
		} else {
			updateSetting(port, username, password, listen, webBasePath)
		}
		if show {
			showSetting(show)
		}
		if (tgbottoken != "") || (tgbotchatid != 0) || (tgbotRuntime != "") {
			updateTgbotSetting(tgbottoken, tgbotchatid, tgbotRuntime)
		}
	case "geoip":
		err := geoipCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			return
		}
		UpdateAllip()
	case "cert":
		err := settingCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			return
		}
		if reset {
			updateCert("", "", "")
		} else {
			updateCert(webCertFile, webKeyFile , webCAFile)
		}
	default:
		fmt.Println("except 'run' or 'v2-ui' or 'setting' subcommands")
		fmt.Println()
		runCmd.Usage()
		fmt.Println()
		v2uiCmd.Usage()
		fmt.Println()
		settingCmd.Usage()
		geoipCmd.Usage()
		fmt.Println()
	}
}
