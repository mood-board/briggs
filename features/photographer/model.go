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
	AvatarURL      string         `json:"avatar_url"`
	Migrate        string         `json:"migrate"` //How the user loggedin FB or Google+
	City           string         `json:"city"`
	Country        string         `json:"country"`
	DateOfBirth    string         `json:"dob"`
	IsActive       bool           `json:"is_active"`
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

//Find Gets the User Detail by username
func (p *Photographer) Find(conf *config.Config, username string) (Photographer, error) {
	session := conf.Session.Copy()
	defer session.Close()
	var photographer Photographer
	collection := session.DB(config.DATABASENAME).C(config.USERSCOLLECTION)
	err := collection.Find(bson.M{"username": username}).One(&photographer)
	return photographer, err
}

func (p *Photographer) FindByID(conf *config.Config, user_id string) (Photographer, error) {
	session := conf.Session.Copy()
	defer session.Close()
	var photographer Photographer
	collection := session.DB(config.DATABASENAME).C(config.USERSCOLLECTION)
	err := collection.Find(bson.M{"id": user_id}).One(&photographer)
	if err != nil {
		return photographer, err
	}
	return photographer, nil
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
	return photographers, err
}

//Exists Checks if the Users E-mail already exists
func (p *Photographer) Exists(conf *config.Config) bool {
	session := conf.Session.Copy()
	defer session.Close()
	var photographer Photographer
	collection := session.DB(config.DATABASENAME).C(config.USERSCOLLECTION)
	err := collection.Find(bson.M{"email": p.Email}).One(&photographer)
	return err == nil
}
