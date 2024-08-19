# monkey

Self-study of the book "Writing an Interpreter in Go" by Thorsten Ball


## The Monkey Programming Language & Interpreter

Expressed as a list of features, Monkey has the following:
- C-like syntax
- variable bindings
- integers and booleans
- arithmetic expressions
- built-in functions
- first-class and higher-order functions
- closures
- a string data structure
- an array data structure
- a hash data structure


```c
let age = 28;
let name = "Andrey";
let result = 10 * (20/2);

// array and hashes support
let myArray = [1, 2, 3, 4, 5];
let tidras = {
  "name": "Andrey",
  "age": 28,
};

// acessing the elements and hashes is done with index expressions:
myArray[0] // => 1
tidras["name"] // => "Andrey"

// the `let` statements can be used to bind functions to names.
let add = fn (a, b) {
    return a + b;
};
```


## REPL

```
go run main.go
```

