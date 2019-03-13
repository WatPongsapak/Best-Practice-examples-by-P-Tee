package merchant

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type MongoDB struct {
	session *mgo.Session
}

func NewMongoDB(session *mgo.Session) *MongoDB {
	return &MongoDB{session: session}
}

func (m MongoDB) MerchantInsert(merchant Merchant) (string, error) {
	ss := m.session.Clone()
	defer ss.Close()

	merchant.ID = bson.NewObjectId()
	if err := ss.DB("").C("merchants").Insert(merchant); err != nil {
		return "", err
	}
	return merchant.ID.Hex(), nil
}

func (m MongoDB) FindMerchantByID(id string) (Merchant, error) {
	ss := m.session.Clone()
	defer ss.Close()

	var rtn Merchant
	if err := ss.DB("").C("merchants").Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&rtn); err != nil {
		return Merchant{}, err
	}
	return rtn, nil
}
