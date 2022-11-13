package config

import (
	"fmt"
	"testing"

	fruit "github.com/gowhale/go-test-data/pkg/fruits"
	"github.com/stretchr/testify/suite"
)

type configSuite struct {
	suite.Suite

	mockJSONReader *mockReaderJSON
}

func (c *configSuite) SetupTest() {
	c.mockJSONReader = new(mockReaderJSON)
}

func (*configSuite) AfterTest() {

}

func TestTerminalSuite(t *testing.T) {
	suite.Run(t, new(configSuite))
}

func (c *configSuite) Test_DisplayMatrix_ReadFile_Pass() {
	fileContents := `{
		"cols": [5,6,13,19,4,27,22,17],
		"rows": [23,20,18,16,12,26,21,25]
	}`
	fileBytes := []byte(fileContents)

	c.mockJSONReader.On("ReadFile", "config-files/Grape").Return(fileBytes, nil)

	cfg, err := loadConfigImpl(c.mockJSONReader, fruit.Grape)

	c.Equal(PinConfig{
		ColPins: []int{5, 6, 13, 19, 4, 27, 22, 17},
		RowPins: []int{23, 20, 18, 16, 12, 26, 21, 25},
	}, cfg)
	c.Nil(err)
}

func (c *configSuite) Test_DisplayMatrix_JSON_Error() {
	c.mockJSONReader.On("ReadFile", "config-files/Grape").Return(nil, nil)

	cfg, err := loadConfigImpl(c.mockJSONReader, fruit.Grape)

	c.Equal(PinConfig{}, cfg)
	c.EqualError(err, "unexpected end of JSON input")
}

func (c *configSuite) Test_DisplayMatrix_ReadFile_Error() {
	cfg := PinConfig{
		RowPins: []int{1, 2, 3},
		ColPins: []int{1, 2, 3, 4, 5, 6},
	}
	c.Equal(cfg.ColCount(), 6)
	c.Equal(cfg.RowCount(), 3)
}

func (c *configSuite) Test_PinConfig_Getters() {
	c.mockJSONReader.On("ReadFile", "config-files/Grape").Return(nil, fmt.Errorf("read file error"))

	cfg, err := loadConfigImpl(c.mockJSONReader, fruit.Grape)

	c.Equal(PinConfig{}, cfg)
	c.EqualError(err, "read file error")
}
