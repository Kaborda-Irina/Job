package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func ConnToDb(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*sql.DB, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := sql.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	//ticker := time.NewTicker(5 * time.Second)
	//go func() {
	//	if true {
	//		for range ticker.C {
	//			_, err := db.Exec("SELECT * FROM hashfiles;")
	//			if err != nil {
	//				log.Fatalln(err)
	//				os.Exit(1)
	//			}
	//		}
	//	}
	//}()

	return db, nil
}
