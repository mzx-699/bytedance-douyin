package util

import "os"

func CreateDirIfNotExist(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
		return nil
	}
	return err
}
