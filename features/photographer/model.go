package photographer

import (
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/ofonimefrancis/brigg/internal/config"
)

type onlineProfiles struct {
	Facebook   string `json:"profile_facebook"`
	Twitter    string `json:"profile_twitter"`
	GooglePlus string `json:"profile_google"`
	Instagram  string `json:"profile_instagram"`
	Website    string `json:"profile_website"`
}

//Photographer Represents a user that uploads images on the platform
type Photographer struct {
	ID             string         `json:"user_id"`
	Username       string         `json:"username"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	Password       string         `json:"password" bson:"-"`
	HashedPassword []byte         `json:"hash_password"`
	Email          string         `json:"email"`
	Migrate        string         `json:"migrate"` //How t
	City           string         `json:"city"`
	Country        string         `json:"country"`
	DateOfBirth    string         `json:"dob"`
	IsActive       string         `json:"is_active"`
	Description    string         `json:"about_me"`
	OnlineProfiles onlineProfiles `json:"online_profile"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

//Add Creates a new User/Photographer
func (p *Photographer) Add(conf *config.Config) error {
	session := conf.Session.Copy()
	defer session.Close()
	collection := session.DB(config.DATABASENAME).C(config.USERSCOLLECTION)
	return collection.Insert(p)
}

//Update Modifies a user credentials
func (p *Photographer) Update(conf *config.Config, update interface{}) error {
	session := conf.Session.Copy()
	defer session.Close()
	collection := session.DB(config.DATABASENAME).C(config.USERSCOLLECTION)
	return collection.Update(bson.M{"id": p.ID}, update)
}

//DisableAccount Disables a users account
func (p *Photographer) DisableAccount(conf *config.Config) error {
	session := conf.Session.Copy()
	defer session.Close()
	collection := session.DB(config.DATABASENAME).C(config.USERSCOLLECTION)
	change := bson.M{"$set": bson.M{"is_active": false, "updated_at": time.Now()}}
	return collection.Update(bson.M{"id": p.ID}, change)
}

//Listings A paginated listing of photographers sorted by time for now
//TODO: The photographers with the highest likes and favourite count will be seen at the top
func (p *Photographer) Listings(conf *config.Config, page int) ([]Photographer, error) {
	session := conf.Session.Copy()
	defer session.Close()
	var photographers []Photographer
	pageSize := 20
	collection := session.DB(config.DATABASENAME).C(config.USERSCOLLECTION)
	err := collection.Find(bson.M{}).Sort("-created_at").Skip(pageSize * (page - 1)).Limit(pageSize).All(&photographers)
	if err != nil {
		return photographers, err
	}
	return photographers, nil
}
