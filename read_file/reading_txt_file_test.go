package read_file_test

import (
	"testing"

	"github.com/joshuatheokurniawansiregar/go-file-handling/read_file"
	"github.com/stretchr/testify/assert"
)

/*
	Testing cases:
	1. Testing file does not exist(return error)
	2. Testing file does exist(return data in string)
	3. Testing file does not exist(return error) using os.Open() with bufio.Reader struct
	4. Testing file does exist(return data in slice string) using os.Open() with bufio.Reader struct
*/

func TestReadFile(t *testing.T) {
	var err error =  read_file.OsReadFileNotExists()

	t.Run("1. Testing file does not exist(return error)", func(t *testing.T) {
		assert.Error(t, err)
	})
	t.Run("2. Testing fiel does exist(return data in string)", func(t *testing.T){
		var data string = read_file.OsReadFileExists()
		assert.Equal(t, `hello world`, data)
	})

	t.Run("3. Testing file does not exist(return error) using os.Open() with bufio.Reader struct", func(t *testing.T) {
		var err error = read_file.OsOpenFileNotExist()
		assert.Error(t, err)
	})

	t.Run("4. Testing file does exist(return data in slice string) using os.Open() with bufio.Reader struct", func(t *testing.T) {
		var data string = read_file.OsOpenFileExist()
		assert.Equal(t, "hello world", data)
	})
}