package config

import "os"

var Version = "1.0.0"

var Env = os.Getenv("ENV")
var Port = os.Getenv("PORT")

var PostgresHost = os.Getenv("POSTGRES_HOST")
var PostgresPort = os.Getenv("POSTGRES_PORT")
var PostgresDB = os.Getenv("POSTGRES_DB")
var PostgresUser = os.Getenv("POSTGRES_USER")
var PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
