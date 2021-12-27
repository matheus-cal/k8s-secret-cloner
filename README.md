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
|   [Go]     |  v1.6 ou maior  |
|[Kubernetes]|  v1.23 ou maior |

[Go]: https://github.com/golang/go
[Kubernetes]: https://github.com/kubernetes/kubernetes

---

## Começando

Os passos a seguir ocorrem considerando que as configurações de autorização e contexto da conexão com
o ambiente *Kubernetes* localmente já foram formadas. 

O primeiro passo para a utilização da aplicação, é exportar como variáveis de ambiente as informações
essenciais para funcionamento no arquivo `env` em `hook/files`:

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

### Variáveis da *Secret*

| *Variável de Ambiente* |            *Valor*                 |
|:----------------------:|:----------------------------------:|
|   SECRET_NAME          |  O nome da *secret* alvo.          |
|   SECRET_PREFIX        |  O prefixo do nome da *secret*.    |

### Variáveis do *Namespace*

| *Variável de Ambiente* |            *Valor*                 |
|:----------------------:|:----------------------------------:|
|  SOURCE_NAMESPACE      |  O nome do *namespace* de origem da *secret*.                         |
|  INTENDED_NAMESPACE    |  O nome do *namespace* pretendido no qual a *secret* vai ser clonada. |

### Variáveis das *Labels*

| *Variável de Ambiente* |              *Valor*                                                |
|:----------------------:|:-----------------------------------------------------------------:|
|   CHART                |  *label* responsável por nomear o *chart*.                        |
|   K8S_ACCOUNT_ID       |  *label* responsável por definir o ID do *Kubernetes Account*.    |
|   K8S_MANAGED_BY       |  *label* responsável por definir o gerenciador da *secret*.       |
|   K8S_NAME             |  *label* responsável por definir o nome *Kubernetes* da *secret*. |
|   NAME                 |  *label* responsável por definir o nome usual da *secret*.        |