### Interpreter for Scroopy programming language

Simple interpreter for Scroopy programming language written in Go.

The implementation is based on ["Writing An Interpreter In Go"](https://interpreterbook.com/) by Thorsten Ball.

### Supports:
* Basic data types: integers, booleans, strings, arrays and hashmaps
* Basic math expressions: `+`, `-`, `/`, `*`
* Basic binary expressions: `>`, `<`, `==`, `!=`
* Variable bindings
* Conditionals
* Functions
* Build-in functions
* Higher-order functions
* Closures

Scroopy has a basic REPL. It stands for "read-eval-print loop", it's a simple interactive programming language  
shell that takes single user inputs, evaluates them and prints the result back.

### Using REPL

Build the REPL's binary:
```bash
make build
```

Running REPL example:
```bash
./scroopy-repl

  ______ ___________  ____   ____ ______ ___.__.
 /  ___// ___\_  __ \/  _ \ /  _ \\____ <   |  |
 \___ \\  \___|  | \(  <_> |  <_> )  |_> >___  |
/____  >\___  >__|   \____/ \____/|   __// ____|
     \/     \/                    |__|   \/

Hello $USER! This is the Scroopy programming language!
Feel free to type in commands
>>
>> print("Hello, world!")
"Hello, world!"
>>
>> let greeter = fn(name) { "Hello, " + name + "!"};
>>
>> greeter("Scroopy")
"Hello, Scroopy!"
>>
```

### Scroopy code examples

Define a function to compute a factorial of a number:
```bash
>> let factorial = fn(n) { if (n == 1) { 1 } else { n * factorial(n-1) }};
>>
>> factorial(5)
120
```

Working with arrays:
```bash
>> let arr = [1,2,3,4,5];
>> arr
[1, 2, 3, 4, 5]
>> len(arr)
5
>>
>> first(arr)
1
>>
>> last(arr)
5
>> rest(arr)
[2, 3, 4, 5]
>>
>> push(arr, 6)
[1, 2, 3, 4, 5, 6]
>>
>> arr
[1, 2, 3, 4, 5]
>>
>> let newArr = push(arr, 42)
>>
>> newArr
[1, 2, 3, 4, 5, 42]
```

Working with hashmaps:
```bash
>> let hm = {"key0": 123, "key1": [1,3,1,2], "key2": "hello, world!"}
>>
>> hm
{"key1":[1, 3, 1, 2], "key2":"hello, world!", "key0":123}
>>
>> hm["key1"]
[1, 3, 1, 2]
>>
>> len(hm["key2"])
13
>>
>> hm["unknown-key"]
null
```
