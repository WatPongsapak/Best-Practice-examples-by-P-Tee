package db

import (
	"fmt"
	"strings"

	"Advance-Golang-Programming/advanced/topic_2_framework/workshop_structure_answer/internal/config"
	"github.com/globalsign/mgo"
)

const MongoDB = "mongodb"

type (
	Database struct {
		Type        string
		MongoDBConn MongoDBConn
	}

	MongoDBConn struct {
		Config  config.MongoDB
		Session *mgo.Session
	}
)

func Init(conf config.Database) (*Database, error) {
	var d Database

	switch strings.ToLower(conf.Type) {
	case MongoDB:
		d = Database{
			Type:        MongoDB,
			MongoDBConn: MongoDBConn{Config: conf.MongoDB},
		}

		info := mgo.DialInfo{
			Addrs:    conf.MongoDB.Addresses,
			Username: conf.MongoDB.Username,
			Password: conf.MongoDB.Password,
			Database: conf.MongoDB.Database,
			Timeout:  conf.MongoDB.Timeout,
		}
		sess, err := mgo.DialWithInfo(&info)
		if err != nil {
			return nil, fmt.Errorf("dial with info error: %s %+v", err, conf.MongoDB)
		}
		sess.SetMode(mgo.Monotonic, true)

		d.MongoDBConn.Session = sess
	default:
		return nil, fmt.Errorf("db type not implement: %s", conf.Type)
	}
	return &d, nil
}

func (d *Database) Close() error {
	switch d.Type {
	case MongoDB:
		if d.MongoDBConn.Session != nil {
			d.MongoDBConn.Session.Close()
		}
	default:
		return fmt.Errorf("db type not implement: %s", d.Type)
	}
	return nil
}
