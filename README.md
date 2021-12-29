# Kubernetes Secret Cloner 

The Kubernetes Secret Cloner is an open-source application that allows secrets to be 
cloned and customized between Kubernetes' namespaces in a agile and reliable way, keeping
its original features and providing a chance to add and embed additional labels.

Additionally, the application is able to check whether the secret already exists in that
intended place. If it already exists, it has its metadata and labels updated.

---

## Behavior

The Kubernetes Secret Cloner works with one-sided jobs. In other words, it produces
a single job per secret.

The flow occurs from the point at which the application checks if the secret exists in the
source namespace and retrieves it if there is. Once the secret is retrieved, if there's not a secret
with the same name, it's cloned in the intended namespace. If there is a secret with the same name, it is
updated with necessary patchs and labels.

---

## Requirements

For the application to work fully, the following tools must be locally installed:

|    Tool    |      Version    |
|:----------:|:---------------:|
|   [Go]     |  v1.17 or later |
|[Kubernetes]|  v1.23 or later |

[Go]: https://github.com/golang/go
[Kubernetes]: https://github.com/kubernetes/kubernetes

---

## Getting Started

The following steps take place considering that the settings related to authorization and context of
your Kubernetes' environment were locally formed.

### Settings

The first step in using the application is exporting all the essential information for operation 
as environment variables. They are located in the file `env` in `hook/files`:

```
export SECRET_NAME=example
export SECRET_PREFIX=first
export SOURCE_NAMESPACE=namespace1
export INTENDED_NAMESPACE=namespace2
export CHART=chart-example
export K8S_ACCOUNT_ID=example
export K8S_MANAGED_BY=example
export K8S_NAME=example
export NAME=example
```

You must edit its values according with the information you need.

Each of these environment variables has its own importance to the application operation, and its definitions
will be explained below:

#### *Secret* Variables

| *Environment Variable* |            *Value*                 |
|:----------------------:|:----------------------------------:|
|   SECRET_NAME          |  The name of target *secret*.      |
|   SECRET_PREFIX        |  Prefix name of target *secret*.   |

#### *Namespace* Variables

| *Environment Variable* |            *Value*                 |
|:----------------------:|:----------------------------------:|
|  SOURCE_NAMESPACE      |  The name of source *namespace* where *secret* is.                    |
|  INTENDED_NAMESPACE    |  The name of intended *namespace* where cloned *secret* is going to.  |

#### *Labels* Variables

| *Environment Variable* |                             *Value*                                  |
|:----------------------:|:--------------------------------------------------------------------:|
|   CHART                |  *label* responsible for naming the *chart*.                         |
|   K8S_ACCOUNT_ID       |  *label* responsible for defining the ID of *Kubernetes Account*.    |
|   K8S_MANAGED_BY       |  *label* responsible for defining the *secret* manager.              |
|   K8S_NAME             |  *label* responsible for defining the *Kubernetes* name of *secret*. |
|   NAME                 |  *label* responsible for defining the usual name of *secret*.        |

### Execution

Once the local environment is set with the environment variables described in the previous step, is the 
moment to execute the application.

Inside the repository `hook/src`, execute the file `main.go`.

```
go run main.go
```

Since the application is executed, to follow success of the operation just follow the log information
made available in the output:

```
2000/01/01 00:00:00 Found the secret 'first-example' in namespace 'namespace1'
2000/01/01 00:00:00 labels: map[app.kubernetes.io/instance:example app.kubernetes.io/managed-by:example app.kubernetes.io/name:example app.kubernetes.io/version:1.1.0 helm.sh/chart:example], annotations: map[meta.helm.sh/release-name:example meta.helm.sh/release-namespace:namespace1], creation: 2000-07-20 11:52:52 -0300 -03
2000/01/01 00:00:00 Secret cloned: {'first-example''namespace2''1515bd3c-2784-4b8f-9017-eba913498001'}
2000/01/01 00:00:00 Labels were patched.
```

---

## Releases

The current and future *releases* will be listed below: 

| Release |  Current Patch  | Release Date |
|:-------:|:---------------:|:------------:|
|   v1.0  |     [v1.0.0]    | Dec 27, 2021 |

[v1.0.0]: https://github.com/matheus-cal/k8s-secret-cloner/releases/tag/1.0.0

---

## Licensing

The Kubernetes Secret Cloner is licensed under [GNU] General Public License v3.0.

The contributors do not guarantee the operation of this application.

[GNU]: https://github.com/matheus-cal/k8s-secret-cloner/blob/main/LICENSE
