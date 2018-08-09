package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/mgo.v2"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
	config     Config
)

const (
	DATABASENAME         = "seemars"
	USERSCOLLECTION      = "users"
	CATEGORIESCOLLECTION = "categories"
)

//Config Holds the configuration settings for our application
type Config struct {
	MongoDB     string
	MongoServer string
	Database    *mgo.Database
	Port        string
	JwtKey      string
	Session     *mgo.Session
	Encryption  Encryption
}

type Encryption struct {
	PublicKey  []byte
	PrivateKey []byte
}

type JWTConfig struct {
	Secret         string `json:"secret"`
	PublicKeyPath  string `json:"public_key_path"`
	PrivateKeyPath string `json:"private_key_path"`
	Expiration     string
}

//Init Configures the Config struct
func Init() {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"127.0.0.1:27017", "127.0.0.1:27018"},
		Username: "",
		Password: "",
	})

	config = Config{}

	if err != nil {
		log.Fatal("Error connecting to Database")
	}

	config.Session = session
	config.MongoDB = DATABASENAME

	config.Encryption.PublicKey, err = ioutil.ReadFile("./config/encryption_keys/app.rsa.pub")
	if err != nil {
		log.Println("Error reading public key")
	}

	config.Encryption.PrivateKey, err = ioutil.ReadFile("./config/encryption_keys/app.rsa")
	if err != nil {
		log.Println("Error reading private key")
	}

	config.JwtKey = "mysecretjwtkeynotsoperfect"

	port := os.Getenv("PORT")
	if port == "" {
		config.Port = "4000"
	}
}

//Get Returns a pointer to the config
func Get() *Config {
	return &config
}
