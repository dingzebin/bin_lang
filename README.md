# Introduction
A simple Intepreter. Most of the code came from "Writing An Interpreter In Go"

# Variables
```
let age = 2;
age = 3;

```
# For
```
let i = 0;
for (i = 0; i < 10; i=i+1) {
  
}
```
# IF-ELSE
```
let i = 0;
if (i < 10) {

} else {
  if (i > 15) {
    
  }
}
```

# Functions
```
let hello = fn(name) {
    return "Hello " + name;
};
hello(" Bin");

```
# Types
```
bool: true false

string: "str"

int: int64
```
# Array Types

## Declarations
```
let arr = [2, "3", 4];
arr[2]; // 4
arr[2] = 5;
arr[2]; // 5
```

## Builtin Array Functions

```
let arr = [2, 3, 4]
len(arr); // 3
first(arr); // 2
last(arr); // 4
rest(arr); // 3, 4
push(arr, 5); // 2, 3, 4, 5

```


# Map
## Declarations
```
let user = {"name": "Bin", "age": 24};
user[name]; // Bin
user[name] = "bin";
user[name]; // bin 
```
## Builtin Map Functions
```
developing...
```

# Builtin Functions

```
puts("3") // 3

```
