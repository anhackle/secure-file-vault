package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()
	InitValidator()

	r := InitRouter()
	r.Run(":9000")

}
