/*
这个包存储了rabbitmq server的配置信息
*/

package common

import (
	"go-seckill/internal/logconf"

	"github.com/sirupsen/logrus"
)

// Ch作为全局变量, 可以被包外引用, rabbitmqServerName是rabbitmqServer容器在redisStore这个网络中的名称, 其他容器可以根据它的名字找到它
var rabbitmqServerName = "rabbitmq-server"
var Ch = GetChannel(rabbitmqServerName)
var rabbitmqServerUsername = "root"
var rabbitmqServerPassword = "12345678"
var rabbitmqServerPort = "5672"
var rabbitmqServerPath = "/root_vhost"

var logger = logconf.BaseLogger.WithFields(logrus.Fields{"component": "rabbitmq-common"})
