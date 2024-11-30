default:
  @just --list
  
clean: 
  #!/usr/bin/bash
  if [ -d "./bin" ]; then
    rm -rf ./bin
  fi
  
build: clean
  ./build.sh
  go build
  
test: clean
  go test -v ./...
  

test-short: clean
  go test ./...
