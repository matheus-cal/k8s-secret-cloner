package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/rest"
)

func TestEnvConfigNil(t *testing.T) {
	_, err := EnvConfig()
	assert.EqualError(t, err, "Missing the environment variable 'SECRET_NAME'")

}

func TestEnvConfigEmpty(t *testing.T) {
	t.Setenv("SECRET_NAME", "")

	_, err := EnvConfig()
	assert.EqualError(t, err, "The value of the environment variable 'SECRET_NAME' cannot be an empty string")
}

func TestEnvConfigOK(t *testing.T) {
	t.Setenv("SECRET_NAME", "test")
	t.Setenv("SECRET_PREFIX", "test")
	t.Setenv("SOURCE_NAMESPACE", "test")
	t.Setenv("INTENDED_NAMESPACE", "test")
	t.Setenv("CHART", "test")
	t.Setenv("K8S_ACCOUNT_ID", "test")
	t.Setenv("K8S_MANAGED_BY", "test")
	t.Setenv("K8S_NAME", "test")
	t.Setenv("NAME", "test")

	ex := make(map[string]string)
	ex["CHART"] = "test"
	ex["INTENDED_NAMESPACE"] = "test"
	ex["K8S_ACCOUNT_ID"] = "test"
	ex["K8S_MANAGED_BY"] = "test"
	ex["K8S_NAME"] = "test"
	ex["NAME"] = "test"
	ex["SECRET_NAME"] = "test"
	ex["SECRET_PREFIX"] = "test"
	ex["SOURCE_NAMESPACE"] = "test"

	envConfig, _ := EnvConfig()
	assert.Equal(t, ex, envConfig)
}

func fakeGetInclusterConfig() (*rest.Config, error) {
	return nil, nil
}

// func TestConfigure(t *testing.T) {
// 	origGetInclusterConfig := rest.InClusterConfig
// 	getInclusterConfigFunc := fakeGetInclusterConfig

// 	defer func() {
// 		getInclusterConfigFunc = origGetInclusterConfig
// 	}()

// 	client := Configure()
// 	assert.NotNil(t, client, getInclusterConfigFunc)

// }

func TestCreateClientset(t *testing.T) {
	t.Setenv("SECRET_NAME", "test")
	t.Setenv("SECRET_PREFIX", "test")
	t.Setenv("SOURCE_NAMESPACE", "test")
	t.Setenv("INTENDED_NAMESPACE", "test")
	t.Setenv("CHART", "test")
	t.Setenv("K8S_ACCOUNT_ID", "test")
	t.Setenv("K8S_MANAGED_BY", "test")
	t.Setenv("K8S_NAME", "test")
	t.Setenv("NAME", "test")

	origGetInclusterConfig := rest.InClusterConfig
	getInclusterConfigFunc := fakeGetInclusterConfig

	defer func() {
		getInclusterConfigFunc = origGetInclusterConfig
	}()

	client, namespace := CreateClientset()
	assert.NotNil(t, client, getInclusterConfigFunc)
	assert.NotNil(t, namespace, getInclusterConfigFunc)
}
