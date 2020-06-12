package config

import (
	"errors"
	"gopkg.in/ini.v1"
	"strconv"
)

type ErrorsIni map[string]map[int]string

func ReadErrorsIniFile(file *ini.File) (*ErrorsIni, error) {

	sections := file.SectionStrings()
	sections = append(sections[:0], sections[1:]...) // Убираем из среза секцию "DEFAULT"

	if len(sections) == 0 {
		return nil, errors.New("empty errors.ini file")
	}
	var errorsIni ErrorsIni
	errorsMap := make(map[string]map[int]string)
	for _, section := range sections {
		keys := file.Section(section).Keys()
		errorMap := make(map[int]string)

		for _, key := range keys {

			keyInt, err := strconv.ParseInt(key.Name(), 10, 64)
			if err != nil {
				return nil, err
			}
			errorMap[int(keyInt)] = key.Value()
		}
		errorsMap[section] = errorMap
	}
	errorsIni = errorsMap

	return &errorsIni, nil
}
