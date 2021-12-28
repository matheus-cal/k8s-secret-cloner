# Kubernetes Secret Cloner 

O Kubernetes Secret Cloner é uma aplicação open-source que permite que *secrets* sejam
clonadas e personalizadas entre *namespaces* Kubernetes de forma e ágil e confiável, mantendo
suas caracteristicas originais e cedendo a oportunidade de acrescentar e incorporar *labels*
adicionais.

Complementarmente, a aplicação é capaz de aferir se a *secret* já existe no *namespace*
pretendido, e já havendo, ela tem seus metadados e *labels* atualizado.

---

## Funcionamento

O Kubernetes Secret Cloner funciona com *jobs* unilaterais, ou seja, produzindo um
único *job* por *secret*.

O fluxo ocorre apartir do ponto no qual a aplicação checa se o *secret* existe no *namespace*
de origem, e o retorna se sim. Uma vez que o *secret* é retornado, se não houver uma *secret* de mesmo 
nome, ela é clonada no *namespace* pretendido. Se já houver uma *secret* de mesmo nome, ela é atualizada
com os *patchs* e *labels* necessários.

---

## Requerimentos

Para o funcionamento da aplicação, é necessário a instalação localmente das seguintes ferramentas
localmente:

| Ferramenta |      Versão     |
|:----------:|:---------------:|
|   [Go]     |  v1.17 ou maior  |
|[Kubernetes]|  v1.23 ou maior |

[Go]: https://github.com/golang/go
[Kubernetes]: https://github.com/kubernetes/kubernetes

---

## Começando

Os passos a seguir ocorrem considerando que as configurações de autorização e contexto da conexão com
o ambiente *Kubernetes* localmente já foram formadas. 

### Configuração

O primeiro passo para a utilização da aplicação, é exportar como variáveis de ambiente as informações
essenciais para funcionamento, no arquivo `env` em `hook/files`:

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

Cada uma destas variáveis de ambiente, possui uma importância no funcionamento da aplicação, e suas definições 
serão explicadas a seguir:

#### Variáveis da *Secret*

| *Variável de Ambiente* |            *Valor*                 |
|:----------------------:|:----------------------------------:|
|   SECRET_NAME          |  O nome da *secret* alvo.          |
|   SECRET_PREFIX        |  O prefixo do nome da *secret*.    |

#### Variáveis do *Namespace*

| *Variável de Ambiente* |            *Valor*                 |
|:----------------------:|:----------------------------------:|
|  SOURCE_NAMESPACE      |  O nome do *namespace* de origem da *secret*.                         |
|  INTENDED_NAMESPACE    |  O nome do *namespace* pretendido no qual a *secret* vai ser clonada. |

#### Variáveis das *Labels*

| *Variável de Ambiente* |              *Valor*                                              |
|:----------------------:|:-----------------------------------------------------------------:|
|   CHART                |  *label* responsável por nomear o *chart*.                        |
|   K8S_ACCOUNT_ID       |  *label* responsável por definir o ID do *Kubernetes Account*.    |
|   K8S_MANAGED_BY       |  *label* responsável por definir o gerenciador da *secret*.       |
|   K8S_NAME             |  *label* responsável por definir o nome *Kubernetes* da *secret*. |
|   NAME                 |  *label* responsável por definir o nome usual da *secret*.        |

### Execução

Uma vez que o ambiente local foi configurado com as variáveis de ambiente previamente descritas no passo
anterior, é agora o momento de executar a aplicação.

Dentro do repositório `hook/src`, execute o arquivo `main.go`.

```
go run main.go
```

Uma vez que o programa é executado, para acompanhar o sucesso da operação, acompanhamos as informações de 
de *log* disponibilizadas na saída:

```
2000/01/01 00:00:00 Found the secret 'first-example' in namespace 'namespace1'
2000/01/01 00:00:00 labels: map[app.kubernetes.io/instance:example app.kubernetes.io/managed-by:example app.kubernetes.io/name:example app.kubernetes.io/version:1.1.0 helm.sh/chart:example], annotations: map[meta.helm.sh/release-name:example meta.helm.sh/release-namespace:namespace1], creation: 2000-07-20 11:52:52 -0300 -03
2000/01/01 00:00:00 Secret cloned: {'first-example''namespace2''1515bd3c-2784-4b8f-9017-eba913498001'}
2000/01/01 00:00:00 Labels were patched.
```

---

## Releases

Os *releases* atuais mantidos, e também os futuros releases, estão e serão listados abaixo. 

| Release |  Current Patch  | Release Date |
|:-------:|:---------------:|:------------:|
|   v1.0  |     [v1.0.0]    | Dec 27, 2021 |

[v1.0.0]: https://github.com/matheus-cal/k8s-secret-cloner/releases/tag/1.0.0

---

## Licenciamento

O Kubernetes Secret Cloner é licenciado sobre [GNU] General Public License v3.0.
Os contribuídores não garantem o funcionamento desta aplicação.

[GNU]: https://github.com/matheus-cal/k8s-secret-cloner/blob/main/LICENSE
