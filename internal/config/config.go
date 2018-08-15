package config

import (
	"encoding/json"
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
	DATABASENAME      = "seemars"
	USERSCOLLECTION   = "users"
	UPLOADSCOLLECTION = "uploads"
	TAGCOLLECTION     = "tags"
)

//Config Holds the configuration settings for our application
type Config struct {
	MongoDB     string
	MongoServer string
	Port        string
	Session     *mgo.Session
	Encryption  Encryption
	GoogleCloud struct {
		Type                string
		ProjectID           string
		PrivateKeyID        string
		PrivateKey          string
		ClientEmail         string
		ClientID            string
		AuthURI             string
		TokenURI            string
		AuthProviderCertURL string
		ClientCertURL       string
	}
}

type Encryption struct {
	Public  []byte
	Private []byte
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

	config.Encryption.Public, err = ioutil.ReadFile("./internal/config/encryption_keys/mykey.pub")
	if err != nil {
		log.Println("Error reading public key")
	}

	config.Encryption.Private, err = ioutil.ReadFile("./internal/config/encryption_keys/mykey.pem")
	if err != nil {
		log.Println("Error reading private key")
	}

	gcBytes, err := ioutil.ReadFile("./internal/config/seemars-cloudstorage.json")
	if err != nil {
		log.Fatal("Error: Google Cloud storage credentials ", err)
	}
	var data map[string]interface{}
	if err := json.Unmarshal(gcBytes, &data); err != nil {
		log.Fatal("Error decoding json...", err)
	}

	config.GoogleCloud.Type = data["type"].(string)
	config.GoogleCloud.ProjectID = data["project_id"].(string)
	config.GoogleCloud.PrivateKeyID = data["private_key_id"].(string)
	config.GoogleCloud.PrivateKey = data["private_key"].(string)
	config.GoogleCloud.AuthURI = data["auth_uri"].(string)
	config.GoogleCloud.TokenURI = data["token_uri"].(string)
	config.GoogleCloud.AuthProviderCertURL = data["auth_provider_x509_cert_url"].(string)
	config.GoogleCloud.ClientCertURL = data["client_x509_cert_url"].(string)

	port := os.Getenv("PORT")
	if port == "" {
		config.Port = "4000"
	}
}

//Get Returns a pointer to the config
func Get() *Config {
	return &config
}
