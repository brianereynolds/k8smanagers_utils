package k8smanagers_utils

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

// compareVersions will compare the versions by breaking them down into major, minor, and patch components
// They are then converting them to integers, and comparing them in sequence.
// It returns 1 when version1 is greater than version2, -1 when version1 is less than version2 and 0 when they are equal
func compareVersions(version1, version2 string) (int, error) {

	v1Parts := strings.Split(version1, ".")
	v2Parts := strings.Split(version2, ".")

	if len(v1Parts) < 3 || len(v2Parts) < 3 {
		err := errors.New("version numbers must be of the format V.v.n")
		return 0, err
	}

	for i := 0; i < 3; i++ {
		v1, _ := strconv.Atoi(v1Parts[i])
		v2, _ := strconv.Atoi(v2Parts[i])

		if v1 > v2 {
			return 1, nil
		}
		if v1 < v2 {
			return -1, nil
		}
	}
	return 0, nil
}

// getContainerServiceClient will create a Container Service Client
func getManagedClusterClient(ctx context.Context, sub string) (*armcontainerservice.ManagedClustersClient, error) {
	cred, err := getAzureCredential(ctx)
	if err != nil {
		return nil, err
	}

	managedClustersClient, err := armcontainerservice.NewManagedClustersClient(sub, cred, nil)
	if err != nil {
		return nil, err
	}

	return managedClustersClient, nil
}

// getAzureCredential will create a new Azure Credential
func getAzureCredential(ctx context.Context) (*azidentity.DefaultAzureCredential, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	return cred, nil
}
