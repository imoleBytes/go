# Go Fundamentals

## Here the basics of go is learnt
    - variables
    - conditional statement
    - loops
    - etc
    
## The Entry Point of a Go Program

every go file mut start with a package name
the entry point of the program is the main function of the  package main

file: first.go
```go
        package main

        func main () {
            print("Hello world")
        }
```
it 's then run with `go run first.go`

go run will compile and run the source file for you at a go. 

but if you want to compile yourself, run `go build first.go` you will get an exxecutable file first. (a binary)


## Package 
In go, Package is like a folder to group together diferent go files

create a folder, and in the folder have several go files and make sure you start the files with package name_of_package like this
```go
package name_of_package

func SayHello(){
    ...
}
```
name_of_package does not necesary have to be name of folder. but the package has to be unique in the module

all functions and variables decalred in all the files of a package are accessible within the package without importing
this occurs because go see all the files within a package as one giant file. 

The most required package iis the main package, where main function reside

### Public and Private items

Everything created in the package remain and can only be useed in the package (across files in the package). But if you want a variable or a function tobe used in another package, make the first letter of the resource capital letter

then import the package in the file that you want to use it.

another important point is there can only be one pakage in folder



## Module

Creating a Module is like having packages in python and a way of managing dependencies like package.json in javascript

to create a module, run `go mod init name_of_module`
this will create a go.mod file which serves as the package.jon file

with the module now, you can run the program with `go run .` this will run the main function of the package main. 


in some kind of folder structure

workspace > modules > packages > files

## Variables
 create variable with the var keyword and the data type after the name of the variable. once a variableis decalred, its default value is assigned depending on the data type. for istance string > ''
 int > 0
 etc

aother way to decalreand initialize a variable is to use the shortcut :=, this will infer the datatype from the value given.
const for decalring a costastant that will not chnge thoughout the progrma

 ```go
 var name string
 name = "imole"
 year := 2014

 const number int = 120
 ```


### data types:
- string must be double quotes
- int
- float
- bool
- 

## Collections
- arrays
    fixed number of items grouped together
    the data type is [num of elements]data typeof the elements
    e.g  
    ```go
    var names [2]string
    names[0] = "imole"
    names[1] = "kay"
    or 
    use the shrtcut
    names := [2]string{"michael", "kola"}
    ```

- slices
    its same as array except that it's not fixed.a dynamic array
    the number of elements is not specified
    e.g
    ```go
        var countries []string
        countries = {"Nigeria", "Ghana"}

        append(countries, "Mali") // this will add "Mali" to the slice countries, the lenght becomes 3
        numOfElemets := len(countries)
        fmt.Println(numOfElements)  // 3 
    ```

- map
    map is like dictioaries in python, or hashmap in other languages. Collection of key-value pairs. The data type is map[key_data_type]value_data_type
    ```go
        var countryCodes map[string]int

        countryCodes = {"nigeria": 234, "benin": 221}
    ```


## FUNCTIONS
codes, statements grouped together as one block of code that can be called just like functions in other language

arguments passsed are copies of the original value.
```go
    func add(a int, b int) int {
        return a + b
    }

    func addAndSubstract(a, b int) (int, int) {
        return a + b, a - b
    }
```
return type if specifeied willapply. you can return multiple values
the data types of the parameters must be given

## Pointers
    just like the pointers in C, it is the addressof a variable
    prefix the variablewith & and get the addresss of the variable

    also * to dereference the value in an address
```go
    var name = "mike"
    addreessofname := &name
    value := *addressofname


    func birthage( age *int){
        *age++
    }```
 
