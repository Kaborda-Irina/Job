package main

import (
	"Job/db"
	"Job/hasher"
	"Job/server"
	"flag"
	"fmt"
	"os"
)

var dirPath string
var doHelp bool
var printText string
var checkHashSumFile bool

//initializes the binding of the flag to a variable that must run before the main() function
func init() {
	flag.StringVar(&dirPath, "d", "", "a specific file or directory")
	flag.BoolVar(&doHelp, "h", false, "help")
	flag.StringVar(&printText, "p", "", "print text")
	flag.BoolVar(&checkHashSumFile, "c", false, "check hash sum files in directory")
}

func main() {
	flag.Parse()

	switch {
	case len(printText) > 0:
		PrintText(printText)
	case doHelp:
		customHelpFlag()
	case len(dirPath) > 0:
		result := hasher.SearchFilePath(dirPath)
		for _, file := range result {
			hashFile := hasher.CreateHash(file, "256")
			fmt.Println(file, hashFile)
			db.PutDatabase(file, hashFile)
		}
	case checkHashSumFile:
		result := db.GetfromDB()
		for _, file := range result {
			fmt.Println(file)
		}
	}

	server.Run()
}
func PrintText(printText string) {
	fmt.Println(printText)
}

func customHelpFlag() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Custom help %s:\nYou can use the following flag:\n", os.Args[0])

		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "  flag -%v \n       %v\n", f.Name, f.Usage)
		})
	}
	flag.Usage()
}
