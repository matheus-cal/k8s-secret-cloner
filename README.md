# Secret Cloner

This is an implementation of an application built to get a secret from a source namespace 
and clone it -- or update it if there is one already -- in a intended namespace using Golang.

It's main objective is to run after the Keycloak client is created for a customer 
(as a post-install or post-upgrade job) to keep secrets up to date in different namespaces.

## How it works

This application creates an unilateral job, in other words, it produces one job
per secret.

It's composed by an one-sidedly flow where the application checks if the target secret
from exists in the source namespace, and retrives it if it does. Once the secret is retrived,
if there's not a secret with the same name, it's cloned in the intended namespace. 

Once there is a secret with the same name as the target one, it s updated with new relevant
information.

Lastly, the required labels are patched. 

You can use the file 'env.sh' in the directory 'scripts' for development, when running locally.

## References

* https://github.com/kubernetes/client-go/
