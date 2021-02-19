package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Perezonance/tinder-copycat/profile-management-service/internal/controllers"
	"github.com/Perezonance/tinder-copycat/profile-management-service/internal/server"
	"github.com/Perezonance/tinder-copycat/profile-management-service/internal/storage"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

const (
	//TODO: Pull from Config
	timeoutDurVal = time.Second * 15
)

//NewConfig returns a new decoded Config struct
func newConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}

//ValidateConfigPath ensures the path given is a proper file; not a dir
func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}

	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

/*ParseFlags will create and parts the CLI flags and returns the path to
be used*/
func parseFlags() (string, error) {
	var configPath string

	flag.StringVar(&configPath, "config", "./config.yml", "path to config"+
		"file")

	flag.Parse()

	if err := validateConfigPath(configPath); err != nil {
		return "", err
	}
	return configPath, nil
}

/*NewRouter initializes a Gorilla mux router and assigns all the endpoints
to the appropriate handlers*/
func newRouter(c *controllers.Controller) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/profiles", c.PostProfilesHandler).Methods(http.MethodPost)
	r.HandleFunc("/profiles", c.GetProfilesHandler).Methods(http.MethodGet)

	return r
}

//Run initializes the server with the given config
func (c Config) run() {
	db := storage.NewMockDynamo()

	s := server.NewServer(db)

	ct := controllers.NewController(s)

	var runChan = make(chan os.Signal, 1)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		c.Server.Timeout.Server,
	)
	defer cancel()

	server := &http.Server{
		Addr:         c.Server.Host + ":" + c.Server.Port,
		Handler:      newRouter(ct),
		ReadTimeout:  c.Server.Timeout.Read,
		WriteTimeout: c.Server.Timeout.Write,
		IdleTimeout:  c.Server.Timeout.Idle,
	}

	signal.Notify(runChan, os.Interrupt, syscall.SIGTSTP)

	log.Printf("Server is being initialized on %v\n", server.Addr)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatalf("Server failed to initialize due to errror:%v", err)
			}
		}
	}()

	interrupt := <-runChan

	log.Printf("Server is shutting down due to %v\n", interrupt)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server was unable to gracefully shutdown due to"+
			"error:%v\n", err)
	}
}

func main() {

	cfgPath, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfgPath)
	cfg, err := newConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg.run()
}
