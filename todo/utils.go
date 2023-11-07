package todo

import (
	"fmt"
	"os"
)

func ReadWholeFile(filePath string, fileName string) ([]byte, error){
	contents, err := os.ReadFile(filePath + fileName)

	if err != nil{
		fmt.Printf("Unable to find the file %s in path %s", fileName, filePath)
		return nil, err
	}

	return contents, nil
}

func CreateNewFile(filePath string, fileName string, fileContents string) error{
	file, err := os.Create(filePath + fileName)

	if err != nil{
		fmt.Printf("Unable to create the file %s in path %s", fileName, filePath)
		return err
	}

	defer file.Close()

	_, writeErr := file.WriteString(fileContents)

	if writeErr != nil{
		fmt.Printf("Unable to write to the file %s in path %s", fileName, filePath)
		return writeErr
	}

	return nil
}