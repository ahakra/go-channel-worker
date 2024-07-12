package filereader

import (
	"encoding/csv"
	"encoding/json"
	"go-channel-worker/model"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

type FileReader struct {
	Location string
	Channel  chan string
}

func NewFileReader(location string, channel chan string) *FileReader {
	return &FileReader{Location: location, Channel: channel}
}

func (f *FileReader) InitReader() (*csv.Reader, error) {
	file, err := os.Open(f.Location)

	if err != nil {
		return nil, err
	}

	return csv.NewReader(file), nil
}

func (f *FileReader) ReadFileToChannel() error {

	reader, err := f.InitReader()
	if err != nil {
		return err
	}

	for {
		dateParts, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return err
		}
		stringify, err := f.StringifyStruct(dateParts)
		if err != nil {
			log.Fatal(err)
		}
		f.Channel <- stringify

	}
	return nil
}

func (f *FileReader) StringifyStruct(dataParts []string) (string, error) {
	var myData model.MyData

	myData.ID = dataParts[0]
	myData.CallingCallReference = CleanString(dataParts[1])
	myData.CalledCallReference = CleanString(dataParts[2])
	myData.CalledCallReference = CleanString(dataParts[2])
	myData.CallingRecordType = CleanString(dataParts[3])
	myData.CallingNumber = CleanString(dataParts[7])
	myData.CalledNumber = CleanString(dataParts[8])
	myData.CallDuration = CleanString(dataParts[9])
	myData.SetupTime = CleanString(dataParts[10])

	jsonData, err := json.Marshal(myData)
	if err != nil {
		return "", err
	}

	jsonString := string(jsonData)
	return jsonString, nil
}

func CleanString(input string) string {
	cleaned := strings.Map(func(r rune) rune {
		if r == '\x00' || !utf8.ValidString(string(r)) {
			return -1
		}
		return r
	}, input)
	return cleaned
}
