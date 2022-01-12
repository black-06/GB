package main

import (
	"encoding/json"
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
)

func UpdateReadme() error {
	gbMap, err := ReadSource()
	if err != nil {
		return err
	}

	for id := range gbMap {
		gb := Gb{}
		err = json.Unmarshal(gbMap[id], &gb)
		if err != nil {
			return err
		}

		path := "../data/" + id
		if _, err = os.Stat(path); errors.Is(err, fs.ErrNotExist) {
			err = os.Mkdir(path, fs.ModePerm)
			if err != nil {
				return err
			}
		}
		err = ioutil.WriteFile(path+"/README.md", []byte(gb.String()), 0666)
		if err != nil {
			return err
		}
	}
	return nil
}
