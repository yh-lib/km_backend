// 配置层 存放程序的配置信息
package config

import (
	"km_backend/utils/logs"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 配置参数
const (
	// 日志时间格式
	timeFormat = string("2006-01-02 15:04:05")
)

var (
	// 监听端口号
	Port string
	// Jwt 签名
	JwtSignKey string
	JwtExpTime int64 // jwt token 过期时间，单位：分钟
	Username   string
	Password   string
)

// 规范返回给前端的数据
type ReturnData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// 构造函数
func NewRetrunData() ReturnData {
	returnData := ReturnData{}
	returnData.Status = 200
	data := make(map[string]interface{})
	returnData.Data = data
	return returnData
}

// 配置日志输出格式
func initLogConfig(logLevel string) {
	// 控制日志的输出级别
	switch logLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	}
	// 添加文件名和行号
	logrus.SetReportCaller(true)
	// 配置日志输出为json格式
	// 原先："file":"D:/Workspace/BaiduSyncdisk/MyNotes/2_IT/DEV/golang/goStudySrc/src/km_backend/utils/logs/logs.go:11"
	// 改为："file":"logs.go"
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: timeFormat,
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			fileName := path.Base(f.File)
			return f.Function, fileName
		},
	})

}

func init() {
	logs.Info(nil, "Start loading the program configuration...")
	// 配置环境变量默认值
	viper.SetDefault("LOG_LEVEL", "debug")                           // 日志输出级别
	viper.SetDefault("PORT", ":8080")                                // 程序监听端口
	viper.SetDefault("JWT_SIGN_KEY", "加密用的SECRET")                   // 获取jwt加密的secret
	viper.SetDefault("JWT_EXPIRE_TIME", 120)                         // 获取jwt过期时间
	viper.SetDefault("USERNAME", "E3AFED0047B08059D0FADA10F400C1E5") // 默认用户名：Admin	通过MD5加密
	viper.SetDefault("PASSWORD", "E64B78FC3BC91BCBC7DC232BA8EC59E0") // 默认密码：Admin123	通过MD5加密
	// 绑定环境变量到配置
	viper.AutomaticEnv()
	// 获取环境变量值并绑定到程序变量
	logLevel := viper.GetString("LOG_LEVEL")
	Port = viper.GetString("PORT")
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	JwtExpTime = viper.GetInt64("JWT_EXPIRE_TIME")
	Username = viper.GetString("USERNAME") // 获取用户名
	Password = viper.GetString("PASSWORD") // 获取密码
	// 加载日志输出级别
	initLogConfig(logLevel)
	// 日志格式加载完成提示
	logs.Info(nil, "The log configuration loading is complete.")

}
