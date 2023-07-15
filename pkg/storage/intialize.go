package storage

import (
	"strconv"

	"github.com/bjoydeep/simple-microservice-proto/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// need to figure out how to use this with gorm
// var DBPool *pgxpool.Pool
var DB_ *gorm.DB

func SetupStorage() {

	// need to figure out how to use this with gorm
	// databaseUrl := "postgres://dbuid:dbpwd@dbhost:5432/dbname"
	// dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	// if err != nil {
	// 	println("Unable to connect to database: %v\n", err)
	// 	//os.Exit(1)
	// } else {

	// 	println("Connnected to DB ---")
	// 	DBPool = dbPool
	// }

	//db, err := gorm.Open(postgres.New(postgres.Config{Conn: DBPool}), &gorm.Config{})
	// db, err := gorm.Open(postgres.New(postgres.Config{
	// 	Conn:                  stdlib.NewConn(dbPool),
	// 	Preparer:              stdlib.NewDefaultPreparer(),
	// 	PreferSimpleProtocol:  true,
	// }), &gorm.Config{})

	//dsn := "host=somehost user=dbuser password=dbpwd dbname=somename port=5432 sslmode=disable TimeZone=America/Los_Angeles"
	dsn := "host=" + config.Cfg.DBHost + " user=" + config.Cfg.DBUser + " password=" + config.Cfg.DBPwd + " dbname=" + config.Cfg.DBName +
		" port=" + strconv.Itoa(config.Cfg.DBPort) + " sslmode=" + config.Cfg.DBSSL + " TimeZone=" + config.Cfg.DBTmz

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB_ = db

	if err != nil {
		// be careful - this prints out the whole dsn string!!
		println("Unable to connect to database: %v\n", err)
	}

	//to close DB pool
	//defer dbPool.Close()

}
