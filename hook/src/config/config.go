package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func EnvConfig() (map[string]string, error) {
	envs := [9]string{
		"SECRET_NAME",
		"SECRET_PREFIX",
		"SOURCE_NAMESPACE",
		"INTENDED_NAMESPACE",
		"CHART",
		"K8S_ACCOUNT_ID",
		"K8S_MANAGED_BY",
		"K8S_NAME",
		"NAME",
	}

	config := map[string]string{}

	for _, envName := range envs {
		val, ok := os.LookupEnv(envName)

		if !ok {
			log.Printf("Missing the environment variable '%s'", envName)
			return config, fmt.Errorf("Missing the environment variable '%s'", envName)
		}

		if val == "" {
			log.Printf("The value of the environment variable '%s' cannot be an empty string", envName)
			return config, fmt.Errorf("The value of the environment variable '%s' cannot be an empty string", envName)
		}
		config[envName] = val
	}

	return config, nil
}

func CreateClientset() (*kubernetes.Clientset, string) {
	var config *rest.Config

	// first tries to configure itself from Kubernetes environment
	config, err := rest.InClusterConfig()

	if err != nil {
		var kubeconfig *string

		// nope, consider tself running outside the cluster
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}

		flag.Parse()

		// use the current context in kubeconfig
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)

		if err != nil {
			panic(err.Error())
		}
	}

	envConfig, _ := EnvConfig()
	clientset, err := kubernetes.NewForConfig(config)
	namespace := envConfig["SOURCE_NAMESPACE"]

	if err != nil {
		panic(err.Error())
	}

	return clientset, namespace
}
