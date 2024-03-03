![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)&nbsp;&nbsp; 
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)&nbsp;&nbsp;

# go-go-power-rangers
A simple golang API to show Power Rangers seasons

## :bookmark: Requirements
- [Go](https://go.dev/)
- [Docker](https://docs.docker.com/engine/install/ubuntu/) 

## :triangular_flag_on_post: Starting
Let's start by cloning the repository

```bash
# Clone repository
git clone https://github.com/arielroque/go-go-power-rangers.git

# Go to go-go-power-rangers
cd go-go-power-rangers
```

## :runner: Running
Let's run the API

### Debug version

```bash
# Run debug version
go run main.go
```

### Release version

```bash
# Build image
docker build -t go-go-power-rangers . 

# Run container
docker run -p 8080:8080 go-go-power-rangers
```

## :cityscape: Populate data (optional)
If you want to populate the API before use, you can use the following script

```bash
# Populate api
./populate_api.sh
```

## :rowboat: Request Samples
Finally, to make requests to API, you can use the following commands

```bash
# Post
curl -X POST -H "Content-Type: application/json" -d '{
  "title": "Power Ranger",
  "description": "blabla"
  }' http://localhost:8080/seasons

# Get
curl http://localhost:8080/seasons

# Get by ID
curl http://localhost:8080/seasons/1

# Patch
curl -X PATCH -H "Content-Type: application/json" -d '{
  "title": "abroba"
  }' http://localhost:8080/seasons/1

# Delete
curl -X DELETE http://localhost:8080/seasons/1
```