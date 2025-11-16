# Go-concepts

## 01-hello-golang
This module contains a minimal executable entry point to verify the Go toolchain and basic project structure.

Run with ```go run```
```sh
cd 01-hello-golang
go run .
```

> ```go run .``` compiles and executes all ```.go``` files in the current directory that belong to the same package.

> It automatically finds and runs the package containing the ```main()``` entry point (i.e., ```package main```).

Build & Execute Binary
```sh
cd 01-hello-golang
go build .
./01-hello-golang
```

## Debug Configuration (VS Code + Delve)
To enable debugging for this module, update the `.vscode/launch.json` file and replace the `program` path with the target directory (`01-hello-golang` in this case):
```sh
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch Package",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "${workspaceFolder}/01-hello-golang",
      "buildFlags": ["-gcflags=all=-N -l"],
    }
  ]
}
```
> The flags `-gcflags=all=-N -l` ensure that compiler optimizations are disabled, making breakpoints, stepping, and variable inspection work as expected with Delve.

Result:
<img width="1440" height="821" alt="Image" src="https://github.com/user-attachments/assets/2a46207e-a372-47d6-9080-8572011b6282" />
