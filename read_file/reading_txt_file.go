package read_file

import (
	"bufio"
	"os"
)

//Reading file using os.ReadFile() method
func OsReadFileNotExists()error{
	var(
		// data []byte
		err error
	)
	// data = make([]byte, 0)
	_, err = os.ReadFile("..test/files/data.txt")
	if err != nil{
		return err
	}
	return nil
}

func OsReadFileExists()string{
	var(
		data []byte
		err error
	)
	data, err = os.ReadFile("../files/hello_world.txt")
	if err != nil{
		return ""
	}
	return string(data)
}

func OsOpenFileNotExist()error{
	var(
		file *os.File
		bufioReader *bufio.Reader
		data []byte
		err error
	)
	file , err = os.Open("../files/helloworld.txt")
	if err != nil{
		return err
	}
	bufioReader = bufio.NewReader(file)
	data = make([]byte, 11)
	_, err = bufioReader.Read(data)
	if err != nil{
		return err
	}
	return err
}

func OsOpenFileExist()string{
	var(
		file *os.File
		reader *bufio.Reader
		data string
		bufferedData []byte
		err error
	)
	file, err = os.Open("../files/hello_world.txt")
	if err != nil{
		return data
	}
	
	reader = bufio.NewReader(file)
	bufferedData = make([]byte, 11)
	_, err = reader.Read(bufferedData)
	if err != nil {
		return data
	}
	data = string(bufferedData)
	return data
}