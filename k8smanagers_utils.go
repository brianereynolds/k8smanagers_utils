package k8smanagers_utils

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

// CompareVersions will compare the versions by breaking them down into major, minor, and patch components
// They are then converting them to integers, and comparing them in sequence.
// It returns 1 when version1 is greater than version2, -1 when version1 is less than version2 and 0 when they are equal
func CompareVersions(version1, version2 string) (int, error) {

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

func getDefaultAzureCredential(ctx context.Context) (*azidentity.DefaultAzureCredential, error) {

	if ctx.Value("DefaultAzureCredential") != nil {
		return ctx.Value("DefaultAzureCredential").(*azidentity.DefaultAzureCredential), nil
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}
	context.WithValue(ctx, "DefaultAzureCredential", cred)
	return cred, nil
}

// GetManagedClusterClient will create a Managed Clusters Client
// The first time the function is called it will cache the object, and return this on subsequent calls.
func GetManagedClusterClient(ctx context.Context, sub string) (*armcontainerservice.ManagedClustersClient, error) {

	if ctx.Value("ManagedClusterClient") != nil {
		return ctx.Value("ManagedClusterClient").(*armcontainerservice.ManagedClustersClient), nil
	}

	cred, err := getDefaultAzureCredential(ctx)
	if err != nil {
		return nil, err
	}

	managedClustersClient, err := armcontainerservice.NewManagedClustersClient(sub, cred, nil)
	if err != nil {
		return nil, err
	}

	context.WithValue(ctx, "ManagedClusterClient", managedClustersClient)

	return managedClustersClient, nil
}

// GetAgentPoolClient will create a Agent Pools Client
// The first time the function is called it will cache the object, and return this on subsequent calls.
func GetAgentPoolClient(ctx context.Context, sub string) (*armcontainerservice.AgentPoolsClient, error) {

	if ctx.Value("AgentPoolsClient") != nil {
		return ctx.Value("AgentPoolsClient").(*armcontainerservice.AgentPoolsClient), nil
	}

	cred, err := getDefaultAzureCredential(ctx)
	if err != nil {
		return nil, err
	}

	agentPoolClient, err := armcontainerservice.NewAgentPoolsClient(sub, cred, nil)
	if err != nil {
		return nil, err
	}
	context.WithValue(ctx, "AgentPoolsClient", agentPoolClient)
	return agentPoolClient, nil
}
