package utils

import (
	"errors"

	"github.com/ofonimefrancis/brigg/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
	storage "google.golang.org/api/storage/v1"
)

//GetGoogleCloud Retrieves an instance of Google Cloud storage
func GetGoogleCloud() (service *storage.Service, err error) {
	conf := config.Get()
	authConf := &jwt.Config{
		Email:      conf.GoogleCloud.ClientEmail,
		PrivateKey: []byte(conf.GoogleCloud.PrivateKey),
		Scopes:     []string{storage.DevstorageReadWriteScope},
		TokenURL:   conf.GoogleCloud.TokenURI,
	}

	client := authConf.Client(oauth2.NoContext)
	service, err = storage.New(client)
	if err != nil {
		return service, errors.New("Problem authenticating to GCS")
	}
	return service, nil
}
