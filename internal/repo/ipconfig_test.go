package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/mta-hosting-optimizer/internal/entity"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func Test_GetIPConfigData(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
		want    entity.MockServiceResponse
		mock    func()
	}{
		{
			name:    "SUCCESS",
			wantErr: false,
			want: entity.MockServiceResponse{
				Data: []entity.MtaData{
					{IP: "123", Hostname: "mta-1", Active: false},
				},
			},
			mock: func() {
				fetchDataFromJSONFile = func(filePath string) (entity.MockServiceResponse, error) {
					return entity.MockServiceResponse{
						Data: []entity.MtaData{
							{IP: "123", Hostname: "mta-1", Active: false},
						},
					}, nil
				}
			},
		},
		{
			name:    "Error From File",
			wantErr: true,
			want:    entity.MockServiceResponse{},
			mock: func() {
				fetchDataFromJSONFile = func(filePath string) (entity.MockServiceResponse, error) {
					return entity.MockServiceResponse{}, errors.New("some error")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mock()
			d := IpConfigModule{}

			got, err := d.GetIPConfigData(context.TODO())
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}

			assert.Nil(t, err)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestFetchDataFromJSONFile(t *testing.T) {
	tests := []struct {
		name             string
		jsonContent      string
		expectedResponse entity.MockServiceResponse
		expectedError    error
		wantErr          bool
	}{
		{
			name:             "Non-existent file",
			jsonContent:      "",
			expectedResponse: entity.MockServiceResponse{},
			expectedError:    os.ErrNotExist,
			wantErr:          true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			// Create a temporary JSON file for testing
			err := ioutil.WriteFile("temp.json", []byte(test.jsonContent), 0644)
			if err != nil {
				t.Fatalf("Error creating temporary JSON file: %v", err)
			}

			defer func() {
				// Clean up: Remove the temporary JSON file after the test
				err := os.Remove("temp.json")
				if err != nil {
					t.Fatalf("Error cleaning up temporary JSON file: %v", err)
				}
			}()

			_, err = fetchDataFromJSONFile("temp.json")

			// Check for errors
			if (err != nil) != test.wantErr {
				fmt.Println(test.name)
				t.Errorf("Usecase.GetInactiveServers() error = %v, wantErr %v", err, test.wantErr)
				return
			}

		})
	}
}
