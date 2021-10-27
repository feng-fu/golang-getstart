### 1. go project terminal.

1. create a new go project:

    `go mod init example.com/demo`

2. use local deps
    
    `go mod edit --replace example.com/demo=../demo`

3. run project

    `go run .` or `go run hello.go`

4. test project

    `go test`

5. build project

    `go build`


### 2. syntax

  1. interface variable
  ```go
    var a string
    var a string = "hello"
    var a = "hello"
    a := "hello" // just in func
  ```
  2. define an any type
  ```go
    var a interface{}
    var i int = 5
    var b = "Hello"
    a = i
    a = b
  ```

