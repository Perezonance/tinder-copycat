package main

import "time"

/*Config is the structure used for configuring the server and its
dependencies*/
type Config struct {
	//Server provides server config options
	Server struct {
		//Host is the local machine IP Address to bidn the HTTP Server to
		Host string `yaml:"host"`

		//Port is the local machine TCP Port to bind the HTTP Server to
		Port string `yaml:"port"`

		//Timeout is used for all timeout properties for the server
		Timeout struct {
			/*Server is the general server timeout to use for graceful
			shutdowns*/
			Server time.Duration `yaml:"server"`

			/*Write is the amount of time to wait until an HTTP Server
			write operation is cancelled*/
			Write time.Duration `yaml:"write"`

			/*Read is the amount of time to wait until an HTTP server
			read opertaion is cancelled*/
			Read time.Duration `yaml:"read"`

			/*Idle is the amount of time to wait until and IDLE HTTP
			Session is closed*/
			Idle time.Duration `yaml:"idle"`
		} `yaml:"timeout"`
	} `yaml:"server"`
}
