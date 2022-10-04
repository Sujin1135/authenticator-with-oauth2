package configs

import (
	"github.com/kamva/mgm/v3"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMgm() {
	err := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI("mongodb+srv://root:kZlVwPbY2936GSjV@cluster0.vt02bgc.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		logrus.Errorln("failed to initialize mgm default configurations:", err)
		panic("occurs panic when failed to initialize mgm")
	}

	logrus.Infoln("Success to initialize mgm default configurations")
}
