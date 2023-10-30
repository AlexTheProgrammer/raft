# gowasm

## build

```sh
    cd cmd/wasm
    GOOS=js GOARCH=wasm go build -o ../../assets/main.wasm
```

## run

```sh
    cd cmd/server
    go run main.go
    
# note that once running changes to the wasm don't require a restart of the server
# I intend to move the server out into a separate tool.
```

## todo
- build a very lightweight tool that essentially runs like npm or yarn
- it should enable you to do a one liner like `raft dev` and make a watcher for your site.
- it should be a standalone tool. that has it's own repo.
- it should be installed with go install (maybe go get)
- it should also support `raft build` and that should generate a dist/ folder with a classic minified html,css,js but also main.wasm  (can't get much smaller than binary)
- it should get rid of the need for the developer to have the assets folder (that is the developer should just be writting go code)
- issue in gopls tooling errors out on syscall/js this is a known issue and apparently there is some work around
