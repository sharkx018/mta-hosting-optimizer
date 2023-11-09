package repo

import (
	"context"
	"encoding/json"
	"github.com/mta-hosting-optimizer/internal/entity"
	"io/ioutil"
)

func (ip *IpConfigModule) GetIPConfigData(ctx context.Context) (entity.MockServiceResponse, error) {

	data, err := fetchDataFromJSONFile("mock_data.json")
	if err != nil {
		return entity.MockServiceResponse{}, err
	}

	return data, nil

}

var fetchDataFromJSONFile = func(filePath string) (entity.MockServiceResponse, error) {
	// Read JSON file
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return entity.MockServiceResponse{}, err
	}

	// Unmarshal JSON data into struct
	var mockResponse entity.MockServiceResponse
	err = json.Unmarshal(jsonData, &mockResponse)
	if err != nil {
		return entity.MockServiceResponse{}, err
	}

	return mockResponse, nil
}
