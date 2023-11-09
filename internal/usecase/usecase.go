package usecase

import (
	"context"
	"fmt"
	"github.com/mta-hosting-optimizer/internal/entity"
)

func (uc *Usecase) GetInactiveServers(ctx context.Context) (entity.InefficentMTAResponseParams, error) {

	// Simulate fetching IP configuration data from a mock service
	ipConfigs, err := uc.ipConfigRepo.GetIPConfigData(ctx)

	if err != nil {
		fmt.Println("[GetInactiveServers] Error while getting the data from GetIPConfigData", err.Error())
		return entity.InefficentMTAResponseParams{}, err
	}

	threshold := uc.thresholdNumber
	hostnames := make(map[string]int)

	for _, config := range ipConfigs.Data {

		_, ok := hostnames[config.Hostname]
		if !ok {
			hostnames[config.Hostname] = 0
		}

		if config.Active {
			hostnames[config.Hostname]++
		}
	}

	var result []string
	for hostname, activeCount := range hostnames {
		if activeCount <= threshold {
			result = append(result, hostname)
		}
	}

	return entity.InefficentMTAResponseParams{
		Data: result,
	}, nil
}
