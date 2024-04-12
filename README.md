# golang

### installation guide

https://go.dev/doc/install

### run 
` 
 go run hello-world/hello.go
`

### build Linux/MacOS
` 
 go build hello-word/hello.go
` 

### build Windows
`
GOOS=windows go build hello-word/hello.go
`

### run builded file
`
./hello
`

### create a module
`
  go mod init [module_name]     
`

### build module
`
  go build     
`

### run
`
 ./[module_name]   
`

### install packages
`
go get github.com/google/uuid
`

---

### Variables

```
  * declare type 
  var name string = "Miqueias"

  * infer type 
  
  var middleName = "Santos"
  lastName := "Sousa"
```