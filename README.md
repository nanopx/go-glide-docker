# go-glide-docker

Example using golang, [glide](https://github.com/Masterminds/glide) and [docker](https://docker.com), [docker-compose](https://docs.docker.com/compose/).  
This simply runs a golang web app under docker, and manages vendor packages using glide.  

## Usage
1. Install docker
2. Clone this repo, then initialize containers using:
```bash
$ docker-compose up --build
```