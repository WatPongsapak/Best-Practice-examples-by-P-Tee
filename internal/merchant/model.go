package merchant

import "github.com/globalsign/mgo/bson"

type Merchant struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name" bson:"name"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
}

type RegisterResponse struct {
	ID string `json:"id"`
}
