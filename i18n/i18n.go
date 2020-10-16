package i18n

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

//JSONToStruct coverter Json to equal struct
func JSONToStruct(name string) error {
	jsonStruct := make(map[string]interface{}) //mudar para estrutura
	loadedJSON, err := loadJSON(name)

	if nil != err {
		log.Print("Erro ao carregar json, ", err)
		return err
	}

	err = json.Unmarshal(loadedJSON, &jsonStruct)
	if nil != err {
		log.Print("Erro ao coverter Json, ", err)
		return err
	}

	return nil
}

func loadJSON(name string) ([]byte, error) {
	jsonFile, err := os.Open("/languages/" + name + ".json")
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
