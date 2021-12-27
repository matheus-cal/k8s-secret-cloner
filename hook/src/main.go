package main

import (
	config "github.com/matheus-cal/k8s-secret-cloner/hook/src/config"
	secrets "github.com/matheus-cal/k8s-secret-cloner/hook/src/secrets"
)

func main() {
	envConfig := config.EnvConfig()

	// Create clientset
	clientset, namespace := config.CreateClientset()

	// Return the secret
	secret := secrets.GetSecret(clientset, namespace)

	// modify secret
	modifiedSecret, err := secrets.ModifySecret(secret)

	if err != nil {
		panic(err.Error())
	}

	// clone secret in intended namespace or update it if there is one already
	done := secrets.CloneSecret(clientset, modifiedSecret)

	// patching labels in previously cloned/updated secret
	secrets.PatchLabelSecret(clientset, done, envConfig["INTENDED_NAMESPACE"], modifiedSecret.Name)
}
