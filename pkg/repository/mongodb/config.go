package mongodb

import "github.com/talento90/gorpo/pkg/errors"

// Configuration for MongoDB
type Configuration struct {
	//Database name
	Database string

	//MongoURL for mongo
	MongoURL string
}

// Validate validates if the configuration is valid
func (c *Configuration) Validate() error {

	if c.Database == "" {
		return errors.EValidation("Missing Database", nil)
	}

	if c.MongoURL == "" {
		return errors.EValidation("Missing ConnectionString", nil)
	}

	return nil
}
