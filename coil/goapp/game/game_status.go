package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var statusMutext = &sync.Mutex{}

// 게임 상태를 저장하는것과 각 고루틴 별로 게임 진행 과정을 표시하기 위한 구조체
type mapStatus struct {
	Width     int
	Heigtt    int
	GameMap   [][]int
	deadPoint int
}

func (ms *mapStatus) SaveGmaeStatus(filePath string) error {
	statusMutext.Lock()

	json, _ := json.Marshal(ms)
	err := ioutil.WriteFile(filePath, json, 0644)

	if err != nil {
		fmt.Println(err)
		return err
	}

	statusMutext.Unlock()
	return nil
}

func (ms *mapStatus) LoadGameStatus(filePath string) error {
	// Open our jsonFile
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println("config file load error!!", err)

		return err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, ms)

	if err != nil {
		return err
	}

	fmt.Println("CONFIG ", string(byteValue), ms)
	return nil
}
