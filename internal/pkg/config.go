package pkg

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Authn Contains a list of users
type Authn struct {
	Users []User `yaml:"users"`
}

// User Identifies a user including the tenant
type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	OrgID    string `yaml:"orgid"`
}

// ParseConfig read a configuration file in the path `location` and returns an Authn object
func ParseConfig(locations *[]string) (*Authn, error) {
	authn := Authn{}

	for _, f := range *locations {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, err
		}

		_authn := Authn{}
		err = yaml.Unmarshal([]byte(data), &_authn)
		if err != nil {
			return nil, err
		}

		authn.Users = append(authn.Users, _authn.Users...)
	}

	return &authn, nil
}
