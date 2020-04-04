image = 353bytes/go

all:
	$(error please pick a target)

build-docker:
	docker build -t $(image) .

run-docker:
	docker run -v ${PWD}:/code --rm -it $(image)

