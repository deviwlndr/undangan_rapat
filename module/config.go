package module

import (
	"github.com/aiteung/atdb"
	"os"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "undangan_rapat",
}

var MongoConn = atdb.MongoConnect(MongoInfo)