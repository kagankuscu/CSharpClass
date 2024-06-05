package readwrite

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Read(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func ReadCsv(filename string) [][]string {
	csvFile, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(csvFile)

	var csvData [][]string
	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		csvData = append(csvData, line)
	}
	return csvData
}

func Write(filename *string, text *string) {
	data := []byte(*text)
	_, file := filepath.Split(*filename)

	savePath := "/Users/kagankuscu/createCsharpClass/"
	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		os.Mkdir(savePath, 0755)
	}

	err := os.WriteFile(savePath+strings.Split(file, ".")[0]+".txt", data, 0777)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("file written successfully")
}

func WriteJson(filename string, data []byte) {
	err := os.WriteFile("Users/kagankuscu/createCsharpClass/"+filename+"_written_by_go.json", data, 0777)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Json file written successfully")
}
