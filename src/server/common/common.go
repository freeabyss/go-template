package common

import (
	"github.com/Sirupsen/logrus"
	"gopkg.in/redis.v5"
	"database/sql"
	"github.com/nsqio/go-nsq"
)

var (
	DB   *sql.DB
	Redis  *redis.Client
	Logger *logrus.Entry
	NsqProducer *nsq.Producer
)
