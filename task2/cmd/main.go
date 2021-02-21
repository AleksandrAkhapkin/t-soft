/*Задание 2
Создать функцию, которая принимает 2 параметра: путь к папке и название файла
Функция должна рекурсивно пробежаться по файлам указанной папки, а также по всем файлам всех подпапок.
Если будет обнаружен файл, имя которого совпадает с именем в параметре, то вернуть путь к этому файлу.
Если файл с указанным названием найден не был, вернуть nil.*/
package main

import (
	"fmt"
	"github.com/AleksandrAkhapkin/testTNS/task2/pkg/logger"
	"github.com/AleksandrAkhapkin/testTNS/task2/service"
	"github.com/pkg/errors"
	"strings"
)

func main() {

	var pathToDir string
	var nameFile string

	fmt.Print("Укажите путь к папке от корня: ")
	_, err := fmt.Scan(&pathToDir)
	if err != nil {
		logger.LogError(errors.Wrap(err, "err in main with Scan pathToDir"))
		return
	}

	if pathToDir == "/" {
		fmt.Println("Необходимо указать хотя бы одну папку от корня")
		return
	}
	if !strings.HasPrefix(pathToDir, "~/") {
		if !strings.HasPrefix(pathToDir, "/") {
			pathToDir = "/" + pathToDir
		}
	}

	fmt.Print("Укажите название файла для поиска: ")
	_, err = fmt.Scan(&nameFile)
	if err != nil {
		logger.LogError(errors.Wrap(err, "err in main with Scan nameFile"))
		return
	}

	pathToFile := service.FindFileInDirRecursive(pathToDir, nameFile)
	if pathToFile == nil {
		fmt.Println("Файл не найден")
		return
	}

	fmt.Printf("Путь до файла: %s\n", pathToFile)
}
