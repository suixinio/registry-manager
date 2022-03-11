package conf

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestInitLoadDb(t *testing.T) {
	assert := assert.New(t)
	testCase := `[Database]
Type = mysql
User = root
Password = 233root
Host = 127.0.0.1:3306
Name = v3
TablePrefix = v3_`
	err := ioutil.WriteFile("testConf.ini", []byte(testCase), 0644)
	fmt.Println(err)
	defer func() { err = os.Remove("testConf.ini"); fmt.Println(err) }()
	assert.NotPanics(func() {
		Init("testConf.ini")
	})
}

func TestInitLoad(t *testing.T) {
	assert := assert.New(t)
	defer func() { err := os.Remove("conf.ini"); fmt.Println(err) }()

	assert.NotPanics(func() {
		Init("conf.ini")
	})
}
