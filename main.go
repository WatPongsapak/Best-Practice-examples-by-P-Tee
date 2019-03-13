package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"Advance-Golang-Programming/advanced/topic_2_framework/workshop_structure_answer/internal/api"
	"Advance-Golang-Programming/advanced/topic_2_framework/workshop_structure_answer/internal/config"
	"Advance-Golang-Programming/advanced/topic_2_framework/workshop_structure_answer/internal/db"
)

func main() {
	port := flag.String("port", "8080", "default port: 8080")
	stage := flag.String("stage", "dev", "product stage: local/dev/sit/prod")
	configPath := flag.String("config", "configs", "configuration path")

	flag.Parse()

	// override config if provide by environment
	if s := os.Getenv("CONFIG_STATE"); s != "" {
		*stage = s
	}

	*stage = config.ParseStage(*stage)
	conf := &config.Config{}
	if err := conf.Init(*stage, *configPath); err != nil {
		panic(err)
	}

	log.Printf("config: %+v", conf)

	dbConn, err := db.Init(conf.Database)
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	route := api.Init(conf, dbConn)
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", *port),
		Handler: route,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
