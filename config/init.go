package config

import (
	"path"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
)

func init() {
	currentAbPath := getCurrentAbPathByCaller()
	tomlFile, err := filepath.Abs(currentAbPath + "/configV21.toml")
	if err != nil {
		panic("read toml file error: " + err.Error())
	}
	if _, err := toml.DecodeFile(tomlFile, &Config); err != nil {
		panic("decode toml file error: " + err.Error())
	}
}

func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
