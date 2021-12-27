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

O primeiro passo para a utilização da aplicação, é exportar como variáveis de ambiente as informações
necessárias no arquivo `env` em `hook/files`:

```
export SECRET_NAME=example
export SECRET_PREFIX=first
export SOURCE_NAMESPACE=namespace1
export INTENDED_NAMESPACE=namespace2
export SECRET_KEY_NAME=secretkeyname
export DOMAIN=domain
export INI_PATH=pod.ini
export MESSAGE_ID=example
export CHART=example
export K8S_ACCOUNT_ID=example
export K8S_MANAGED_BY=example
export K8S_NAME=example
export NAME=example
```