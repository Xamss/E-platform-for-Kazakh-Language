package main

import (
	"flag"
	"log"

	"xamss.diploma/config"
	"xamss.diploma/internal/repository/pgrepo"
	"xamss.diploma/internal/server"
	"xamss.diploma/internal/usecase"
	"xamss.diploma/pkg/jwttoken"
)

func main() {

	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	flag.IntVar(&cfg.HTTP.Port, "port", cfg.HTTP.Port, "API server port")
	flag.StringVar(&cfg.ENV, "env", cfg.ENV, "Environment (development|staging|production)")

	flag.StringVar(&cfg.DB.DSN, "db-dsn", cfg.DB.DSN, "PostgreSQL DSN")

	flag.IntVar(&cfg.DB.MaxOpenConns, "db-max-open-conns", cfg.DB.MaxOpenConns, "PostgreSQL max open connections")
	flag.IntVar(&cfg.DB.MaxIdleConns, "db-max-idle-conns", cfg.DB.MaxIdleConns, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.DB.MaxIdleTime, "db-max-idle-time", cfg.DB.MaxIdleTime, "PostgreSQL max connection idle time")

	flag.Parse()

	pool, err := pgrepo.NewRepoPG(struct {
		DSN          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}{
		cfg.DB.DSN,
		cfg.DB.MaxOpenConns,
		cfg.DB.MaxIdleConns,
		cfg.DB.MaxIdleTime,
	})
	if err != nil {
		log.Printf("connection to DB err: %s", err.Error())
	}
	log.Println("connection success")

	token := jwttoken.New(cfg.Token.SecretKey)
	srvs := usecase.New(pool, token, cfg)
	hndlr := server.New(srvs)
	server := server.NewHttpServer(
		hndlr.InitRouter(),
		struct {
			Port         int
			IdleTimeout  string
			ReadTimeout  string
			WriteTimeout string
		}{
			cfg.HTTP.Port,
			cfg.HTTP.IdleTimeout,
			cfg.HTTP.ReadTimeout,
			cfg.HTTP.WriteTimeout,
		})

	log.Println("server started")

	err = server.Serve()
	if err != nil {
		log.Fatal("Failed to start HTTP server")
	}

}
