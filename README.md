# go-bug-repro

Build a Docker image and start a container with all the requirements
```bash
docker build -t go-bug-repro .
docker run --rm -it go-bug-repro /bin/bash
```

Once inside container, you can try the following things:
Because of CGO, compilation might take some time, please be patient.
```bash
# This woks well
go run -mod=mod main.go

go mod vendor
# This will end up in a loop
go run -mod=vendor main.go

# This is the fun part
mkdir /app/dir && cp -R * /app/dir/
cd /app/dir/
# It no longer loops
go run -mod=vendor main.go
```

