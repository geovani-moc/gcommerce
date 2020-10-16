package i18n

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//Dictionary for translater
type Dictionary struct {
	Pages struct {
		Home    string `json:"home"`
		Product string `json:"product"`
		Profile string `json:"profile"`
		Stock   string `json:"stock"`
	}
}

//JSONToStruct coverter Json to equal struct
func JSONToStruct(path string) (Dictionary, error) {

	var jsonStruct Dictionary
	loadedJSON, err := loadJSON(path)

	if nil != err {
		log.Print("Erro ao carregar json, ", err)
		return jsonStruct, err
	}

	err = json.Unmarshal(loadedJSON, &jsonStruct)
	if nil != err {
		log.Print("Erro ao coverter Json, ", err)
		return jsonStruct, err
	}

	return jsonStruct, nil
}

func loadJSON(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	if nil != err {
		log.Print("Erro ao abrir arquivo json, ", err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if nil != err {
		log.Print("Erro ao ler arquivo json, ", err)
		return nil, err
	}

	return byteValue, nil
}

//GetAllDictionaries return all dictionaries
func GetAllDictionaries() ([]Dictionary, error) {
	var dictionaries []Dictionary

	err := filepath.Walk("./i18n/languages", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".json") {
			dictionary, err := JSONToStruct(path)
			if nil != err {
				log.Print("Error de listagem de arquivos", err)
			} else {
				dictionaries = append(dictionaries, dictionary)
			}
		}
		return err
	})

	if nil != err {
		log.Print("Erro ao carregar dicionarios, ", err)
	}
	return dictionaries, err
}
