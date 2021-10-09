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

### TODO
- [ ] Support floats
- [ ] Add actual Scroopy code examples
- [ ] Add basic benchmarks: lexing, parsing, evaluating
- [ ] Eliminate all the existing TODOs in code
