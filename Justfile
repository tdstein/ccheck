application := './cmd/ccheck'

build:
    ./scripts/build.bash {{ application }}

clean:
    rm -r ./bin/

run *args:
    go run {{ application }} {{ args }}

test:
    go test -cover -timeout 30s ./... 