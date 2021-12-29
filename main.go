package main

import (
	"time"

	log "github.com/wimark/liblog"
	mongo "github.com/wimark/libmongo"
	wimark "github.com/wimark/libwimark"

	"github.com/robfig/cron/v3"
)

const coll = "manufacturer_mac"

var (
	Version = "NO VERSION SET"
	Commit  = "NO COMMIT SET"
	Build   = "NO BUILD SET"
	Service = "manufacturer-parser"
)

//Conf конфигурация приложения
var Conf Config

func main() {
	log.InitSingleStr(Service)
	var version = wimark.MakeVersion(Version, Commit, Build)
	log.Info("Start %s", Service)
	log.Info("Version of daemon is: %+v", version)
	Conf.Init()
	worker()
	c := cron.New()
	c.AddFunc("@daily", worker)
	c.Run()
}

func worker() {
	m, err := getAndParseData()
	if err != nil {
		log.Error("getAndParseData error: %s", err.Error())
	}
	err = intertToDB(m)
	if err != nil {
		log.Error("intertToDB error: %+s", err.Error())
	}
}

func intertToDB(data []ManufData) error {
	var mongoDb = new(mongo.MongoDb)
	err := mongoDb.ConnectWithTimeout(Conf.DBURL, 20*time.Second)
	if err != nil {
		return err
	}
	defer mongoDb.Disconnect()

	c, _ := mongoDb.Count(coll, map[string]interface{}{})
	if c > 0 && (c == len(data) || len(data) < 1000) {
		log.Info("intertToDB update not need")
		return nil
	} else if c == 0 && len(data) == 0 {
		log.Info("intertToDB rise gitInitData")
	}

	var createQuery = make([]interface{}, len(data))

	for i, v := range data {
		createQuery[i] = v
	}
	if c > 0 {
		mongoDb.RemoveAll(coll)
	}
	mongoDb.InsertBulk(coll, createQuery...)
	return nil

}
