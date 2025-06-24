package initialize

func Run() {
	LoadConfig()
	InitCronJob()
	InitLogger()
	InitMysql()
	InitRedis()
	InitValidator()
	InitMinio()

	r := InitRouter()
	r.Run(":9005")

}
