package main

import (
	"Job/hasher"
	"Job/server/repo"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DataCurrent struct {
	FileName string
	Hash     string
}

func main() {
	var checkHashSumFile string
	flag.StringVar(&checkHashSumFile, "c", "", "check hash sum files in directory")
	flag.Parse()

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	allDataFromDB := repo.GetData()
	for _, data := range allDataFromDB {
		fmt.Println(data)
	}

	var allDataCurrent []DataCurrent
	if len(checkHashSumFile) > 0 {
		result := hasher.SearchFilePath(checkHashSumFile)
		for _, file := range result {
			var dataCurrent DataCurrent
			hashFile := hasher.CreateHash(file, "256")
			fmt.Println(file, hashFile)
			dataCurrent.FileName = file
			dataCurrent.Hash = hashFile
			allDataCurrent = append(allDataCurrent, dataCurrent)
		}
	}

	MatchData(allDataFromDB, allDataCurrent)
}

func MatchData(allDataFromDB []repo.DataFromDB, allDataCurrent []DataCurrent) {
	db, err := repo.ConnToDb(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatalln(err)
	}
	if allDataFromDB == nil {
		os.Exit(2)
	}
	for _, dataFromDB := range allDataFromDB {
		for _, dataCurrent := range allDataCurrent {
			if dataFromDB.FullFilePath == dataCurrent.FileName {
				if dataFromDB.Hash != dataCurrent.Hash {

					fmt.Printf("Modification data \n from db - %s %s, current - %s, %s", dataFromDB.FullFilePath, dataFromDB.Hash, dataCurrent.FileName, dataCurrent.Hash)
					//PutTable(dataCurrent.FileName, dataCurrent.Hash)
					_, err = db.Query("DROP TABLE hashfiles;")
					if err != nil {
						log.Fatalln(err)
					}
					os.Exit(1)

				}

			}
		}
	}

}

func PutTable(fileName, fileHash string) {
	db, err := repo.ConnToDb(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatalf("%v", err)
	}
	_, err = db.Exec("INSERT INTO hashfiles (fileName, fullFilePath, hashSum, algorithm)  VALUES ($1,$2,$3,'sha256');", fileName, fileName, fileHash)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println("successful insert")
}
