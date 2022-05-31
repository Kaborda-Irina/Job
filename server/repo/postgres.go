package repo

import (
	"log"
	"os"
)

type DataFromDB struct {
	FileName     string
	FullFilePath string
	Hash         string
	Algorithm    string
}

func GetData() []DataFromDB {
	db, err := ConnToDb(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatalln(err)
	}

	var res []DataFromDB
	sql, err := db.Query("SELECT fileName, fullFilePath, hashSum, algorithm FROM hashfiles;")
	if err != nil {
		log.Fatalln(err)
	}
	for sql.Next() {
		var dataFromDB DataFromDB
		err = sql.Scan(&dataFromDB.FileName, &dataFromDB.FullFilePath, &dataFromDB.Hash, &dataFromDB.Algorithm)
		res = append(res, dataFromDB)
	}
	return res
}

func PutTable() string {
	db, err := ConnToDb(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	_, err1 := db.Exec("INSERT INTO hashfiles (fileName, fullFilePath, hashSum, algorithm) VALUES ('5.txt','5/5.txt','4566','sha256');")
	if err1 != nil {
		log.Fatalf("%v", err)
	}
	res := "Succesful creation of table"
	return res
}
