package configs

import (
	"os"
	"strings"
)

func ReadConfig(path string) (map[string]string, error) {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var dat []byte

	_, err = f.Read(dat)
	if err != nil {
		return nil, err
	}

	config := make(map[string]string)
	data := strings.Split(string(dat), "\n")
	for _, conf := range data {

		field := strings.TrimSpace(conf)
		confArr := strings.Split(field, "=")

		config[confArr[0]] = confArr[1]
	}
	return config, nil
}
