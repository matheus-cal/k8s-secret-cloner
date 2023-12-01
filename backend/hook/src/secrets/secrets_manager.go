package secrets

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"

	config "github.com/matheus-cal/k8s-secret-cloner/hook/src/config"
)

func GetSecret(clientset *kubernetes.Clientset, namespace string) *v1.Secret {
	secrets, err := clientset.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	envConfig, _ := config.EnvConfig()

	secretName := fmt.Sprintf("%s-%s", envConfig["SECRET_PREFIX"], envConfig["SECRET_NAME"])
	secret, err := clientset.CoreV1().Secrets(namespace).Get(context.TODO(), secretName, metav1.GetOptions{})

	if errors.IsNotFound(err) {
		log.Printf("Secret '%s' not found in namespace '%s'\n", secretName, namespace)
		var all []string

		for _, secret := range secrets.Items {
			all = append(all, secret.Name)
		}

		log.Printf("All secrets found: %s\n", strings.Join(all, ", "))
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		log.Printf("Error getting secret '%s' in namespace '%s': %v\n",
			secretName, namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		log.Printf("Found the secret '%s' in namespace '%s'\n", secretName, namespace)
		// TODO: make this string more generic instead of Golang structures
		log.Printf("labels: %s, annotations: %s, creation: %s\n", secret.Labels, secret.Annotations, secret.CreationTimestamp)
	}

	return secret
}

func ModifySecret(secret *v1.Secret) (*v1.Secret, error) {
	originInterface := *secret
	content, _ := json.Marshal(originInterface)

	var modifiedSecret *v1.Secret
	err := json.Unmarshal(content, &modifiedSecret)

	if err != nil {
		return nil, err
	}

	envConfig, _ := config.EnvConfig()
	modifiedSecret.Namespace = envConfig["INTENDED_NAMESPACE"]
	modifiedSecret.ResourceVersion = ""
	modifiedSecret.UID = ""

	return modifiedSecret, nil
}

func CloneSecret(clientset *kubernetes.Clientset, modifiedSecret *v1.Secret) *v1.Secret {
	envConfig, _ := config.EnvConfig()
	intended_namespace := envConfig["INTENDED_NAMESPACE"]

	done, err := clientset.CoreV1().Secrets(intended_namespace).Create(context.TODO(), modifiedSecret, metav1.CreateOptions{})

	if err != nil {
		log.Printf("Error in cloning secret: '%s'\n", err)

		if strings.Contains(err.Error(), "already exists") {
			done, err := clientset.CoreV1().Secrets(intended_namespace).Update(context.TODO(), modifiedSecret, metav1.UpdateOptions{})

			if err != nil {
				panic(err.Error())
			}

			log.Printf("Secret updated: {'%s''%s''%s'}\n", done.Name, done.Namespace, done.UID)

			return done
		}
	} else {
		log.Printf("Secret cloned: {'%s''%s''%s'}\n", done.Name, done.Namespace, done.UID)
	}

	return done
}

func PatchLabelSecret(clientset *kubernetes.Clientset, secret *v1.Secret, namespace string, secretName string) {
	envConfig, _ := config.EnvConfig()
	secret.Labels["helm.sh/chart"] = envConfig["CHART"]
	secret.Labels["app.kubernetes.io/account-id"] = envConfig["K8S_ACCOUNT_ID"]
	secret.Labels["app.kubernetes.io/managed-by"] = envConfig["K8S_MANAGED_BY"]
	secret.Labels["app.kubernetes.io/name"] = envConfig["K8S_NAME"]
	secret.Labels["name"] = envConfig["NAME"]

	payloadBytes, _ := json.Marshal(secret)
	_, err := clientset.CoreV1().Secrets(namespace).Patch(context.TODO(), secretName, types.StrategicMergePatchType, payloadBytes, metav1.PatchOptions{})

	if err != nil {
		log.Fatalf("Fail to patch: %s", err)
	}

	log.Println("Labels were patched.")
}
