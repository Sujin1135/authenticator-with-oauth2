package configs

import (
	"fmt"
	"github.com/kamva/mgm/v3"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func InitDB() {
	uri := fmt.Sprintf("%s://%s/%s?%s", os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_DATABASE_NAME"), os.Getenv("DB_OPTION_STRINGS"))
	err := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI(uri))
	if err != nil {
		logrus.Fatalln("failed to initialize mgm default configurations:", err)
	}

	logrus.Infoln("Success to initialize mgm default configurations")
}
