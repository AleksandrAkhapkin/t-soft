package service

import (
	"github.com/AleksandrAkhapkin/testTNS/task2/pkg/logger"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

//Функция рекурсивно пробегается по файлам указанной папки, а также по всем файлам всех подпапок.
//Если будет обнаружен файл, имя которого совпадает с именем в параметре, то вернутся путь к этому файлу.
//Если файл с указанным названием найден не был, вернется nil
func FindFileInDirRecursive(pathToDir, nameFile string) interface{} {

	//данная функция сделана согласно поставленному заданию, однако,
	//я бы предпочел написать функцию которая будет возвращать string и error:
	//- В случае нахождения файла возвращать строку + nil,
	//- В случае не найденного файла возвращать "" + error

	var pathToFile string

	err := filepath.Walk(pathToDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				if info != nil {
					if !info.Mode().IsRegular() {
						return nil
					}
					return nil
				}
				logger.LogError(errors.Wrap(err, "err with WalkFunc"))
				return nil
			}
			if info.Name() == nameFile {
				pathToFile = path
				return nil
			}

			return nil
		})
	if err != nil {
		return nil
	}

	if pathToFile == "" {
		return nil
	}

	return pathToFile
}
