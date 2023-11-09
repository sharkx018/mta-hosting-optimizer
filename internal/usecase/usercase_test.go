package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/mta-hosting-optimizer/internal/entity"
	"github.com/mta-hosting-optimizer/internal/mock_gen"
	"testing"
)

func Test_GetInactiveServers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ipconfigMock := mock_gen.NewMockIpConfigResource(ctrl)

	tests := []struct {
		name    string
		mock    func() *Usecase
		useSF   bool
		wantErr bool
	}{
		{
			name: "Error from ipconfig service",
			mock: func() *Usecase {

				ipconfigMock.EXPECT().GetIPConfigData(gomock.Any()).Return(entity.MockServiceResponse{}, errors.New("some error"))

				return &Usecase{
					ipConfigRepo: ipconfigMock,
				}
			},
			wantErr: true,
		},
		{
			name: "Success response",
			mock: func() *Usecase {

				ipconfigMock.EXPECT().GetIPConfigData(gomock.Any()).Return(entity.MockServiceResponse{
					Data: []entity.MtaData{
						{IP: "123", Hostname: "mta-1", Active: true},
						{IP: "124", Hostname: "mta-1", Active: false},
						{IP: "125", Hostname: "mta-2", Active: true},
						{IP: "126", Hostname: "mta-2", Active: true},
						{IP: "127", Hostname: "mta-3", Active: true},
						{IP: "128", Hostname: "mta-3", Active: true},
						{IP: "129", Hostname: "mta-4", Active: false},
						{IP: "130", Hostname: "mta-4", Active: true},
						{IP: "131", Hostname: "mta-4", Active: true},
						{IP: "132", Hostname: "mta-4", Active: true},
					},
				}, nil)

				return &Usecase{
					ipConfigRepo:    ipconfigMock,
					thresholdNumber: 2,
				}
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			module := test.mock()
			_, err := module.GetInactiveServers(context.TODO())

			if (err != nil) != test.wantErr {
				fmt.Println(test.name)
				t.Errorf("Usecase.GetInactiveServers() error = %v, wantErr %v", err, test.wantErr)
				return
			}

		})
	}

}
