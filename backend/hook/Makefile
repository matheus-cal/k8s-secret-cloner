TAG:=0.2.0
IMAGE:=k8s-utils
COMMIT_ID=$(shell git log -n 1 --decorate=no --no-merges | head -1 | cut -f2 -d' ')

build:
	docker build --build-arg COMMIT_ID=${COMMIT_ID} --build-arg APP_VERSION=${TAG} -t ${IMAGE}:${TAG} .