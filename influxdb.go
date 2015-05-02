package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/influxdb/influxdb/client"
)

const (
	InfluxDBHost    = "localhost"
	InfluxDBPort    = 8086
	InfluxDBPublic  = "celadon_public"
	InfluxDBPrivate = "celadon_private"
)

func Push_Data(pts []client.Point, isPublic bool) {

	con := get_Con()

	db := InfluxDBPrivate

	if isPublic {
		db = InfluxDBPublic
	}

	bps := client.BatchPoints{
		Points:          pts,
		Database:        db,
		RetentionPolicy: "default",
	}
	_, err := con.Write(bps)
	if err != nil {
		log.Fatal(err)
	}
}

func get_Con() *client.Client {

	u, err := url.Parse(fmt.Sprintf("%s:%d", InfluxDBHost, InfluxDBPort))
	if err != nil {
		log.Fatal(err)
	}

	conf := client.Config{
		URL:      *u,
		Username: os.Getenv("INFLUX_USER"),
		Password: os.Getenv("INFLUX_PWD"),
	}

	con, err := client.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	_, _, err = con.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return con

}
