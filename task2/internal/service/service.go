package service

import (
	"github.com/AleksandrAkhapkin/testTNS/task2/pkg/logger"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

//Функция рекурсивно пробегается по файлам указанной папки, а также по всем файлам всех подпапок.
//Если будет обнаружен файл, имя которого совпадает с именем в параметре, то вернутся путь к этому файлу.
//Если файл с указанным названием найден не был, вернется ошибка
func FindFileInDirRecursive(pathToDir, nameFile string) interface{} {

	//данная функция сделанна на основе задания,
	//я бы делал из функции возврат строки + ошибки, вместо использования интерфейса.

	var pathToFile string

	err := filepath.Walk(pathToDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				if info != nil {
					if !info.Mode().IsRegular() {
						return nil
					}
					logger.LogError(errors.Wrap(err, "err with WalkFunc info == nil"))
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
