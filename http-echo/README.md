# WAX Sample - Echo

## Running

If you want to try livereloading views from FS use:

```
go run ./cmd/main.go -isDev
```

When you will skip ```-isDev``` views will be loaded from embedded views.

***I personally use wgo to run GO apps during development for go livereload.***

On server startup WAX will generate ```/views/model.generated.d.ts``` with types exposed for views.

This sample also starts WebSocket for views live-reloading.

## Editing view files

Install package ```wax-jsx``` for intellisense in TSX/JSX (see package.json).
