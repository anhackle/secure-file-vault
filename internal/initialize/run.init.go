package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()
	InitValidator()
	InitMinio()

	r := InitRouter()
	r.Run(":9005")

}
