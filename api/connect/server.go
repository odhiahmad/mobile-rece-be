package connect

import (
	"bri-rece/api/models"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	cfg *viper.Viper
	db  *gorm.DB
	err error
)

type Connect interface {
	SqlDb() *gorm.DB
	Config() *viper.Viper
	ApiServer(addr []string) string
}

type connect struct{}

func NewConnect() Connect {
	return &connect{}
}

func (i *connect) SqlDb() *gorm.DB {
	dbUser := i.Config().GetString("DB_USER")
	dbPassword := i.Config().GetString("DB_PASSWORD")
	dbHost := i.Config().GetString("DB_HOST")
	dbPort := i.Config().GetString("DB_PORT")
	dbName := i.Config().GetString("DB_NAME")

	db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s ", dbHost, dbUser, dbPassword, dbName, dbPort)))

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Account{}, &models.Wallet{}, &models.WalletHistory{}, &models.Otp{}, &models.Merchant{}, &models.UserMerchant{})

	return db
}

func (i *connect) Config() *viper.Viper {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	cfg = viper.GetViper()
	return cfg
}

func (i *connect) ApiServer(addr []string) string {
	switch len(addr) {
	case 0:
		if port := i.Config().GetString("HTTPPORT"); port != "" {
			debugPrint("Environment variable PORT=" + port)
			return ":" + port
		}
		debugPrint("Environment variable PORT is undefined. Using port :8081 by default")
		return ":8081"
	case 1:
		return addr[0]
	default:
		panic("too many parameters")
	}
}

func debugPrint(format string, values ...interface{}) {
	fmt.Println(format)
}
