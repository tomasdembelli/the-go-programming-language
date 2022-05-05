# the-go-programming-language
Notes from the book [`The Go Programming Language`](https://www.gopl.io/) Alan A. A. Donovan and Brian W. Kernighan. 

[Source code](https://github.com/adonovan/gopl.io) for the code examples.

## Intro

Go is a general purpose language.

Go has automatic memory management or garbage collection.

Go is a replacement for untyped scripting languages. It balances expressiveness with safety. 

From C: expression syntax, control-flow statements, basic data types, call-by-value parameter passing, pointers. Also, efficient machine code and cooperate well with the OS abstraction.

Simplicity is the key to good software.

Simplicity of design ensures stability, security and coherence.

There is no class and class hierarchies. Complex objects behaviors are created from simpler ones by composition, not inheritance.

Relationship between concrete and abstract types (interfaces) is implicit, so a concrete type may satisfy an interface that the type's author was unaware of.

`go` tool is used for compilation, testing, benchmarking, linting, documentation, etc.


## Tutorial

Go is a compiled language. The Go toolchain converts a source program and the things it depends on into instructions in the native machine language of a computer.

Go natively handles Unicode.

Go code is organized into packages. A package consists of one or more .go source files in a single directory.

Package `main` is special. It defines a `standalone executable program`, not a library. The `main` function is special, it is where `execution of the program begins`. 

**Whatever `main` does is what the program does**.

A program will not compile if there are missing imports or if there are unnecessary ones.

The order of declarations (type, constants, variables) does not matter.

Multiple statements must be separated via semicolon if they are on the same line.

Comments begin with `//`.

Short variable declaration `:=` declares variables and gives them appropriate types based on the initializer values.

For loop is the only loop statement. All three parts are _optional_.
```
for initialization; condition; post {
// zero or more statements
}
```
Initialization must be a _simple statement_: short variable declaration, an increment or assignment statement, or a function call.

Go does not permit unused local variables.

Blank identifier: `_` underscore. It is used whenever syntax requires a variable name but program logic does not.
```
for _, thing := range sliceOfThings {
// some statements
}
```

Variable declaration:
```
s := ""         // cannot be used outside of a function (at package level)
var s string
var s = "" 
var s string = ""
```

A map provides a constant-time operations to store, retrieve for an item in the set.
The map key can be any type whose values can be compared with `==`. The map value can be any type at all.
The order of map iteration is random _practically_.

Quoted string conversion character is `%q`.

A map is a reference to the data structure created by make. When a map is passed as an argument to a function, the function receives a copy of the reference. Any change in the function will be visible on the actual map. Example [code](./map/main.go).

The value of a constant **must** be a string, number or boolean.

Composite literals are a compact notation for instantiating any of Go's composite types, i.e., slice, struct.

A `goroutine` is a concurrent function execution.

A `channel` is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine.

When one goroutine attempts to send or receive on an un-buffered channel, it blocks until another goroutine attempts the corresponding receive or send operation. When the value is transferred, both goroutines proceed.

HTTP server runs the handler functions for each incoming request in a separate goroutine so that it can serve multiple requests simultaneously.

`io.Writer` interface is commonly used across the standard library. From stdout to http response writer. Investigate further and use it to take advantage of its common usage.

Tagless switch is equivalent to switch true.
The `default` statement can be placed in between case statements.
Case statements do not fall through from one to next. But there is a fallthrough statement that overrides this behavior. An [example](https://golangbyexample.com/fallthrough-keyword-golang/) for fallthrough statement.

The `&` operator yields the address of a variable, and the `*` operator retrieves the variable that the pointer refers to.

## Program Structure

**What is a programming language?**
- One builds large programs from small set of constructs.
- Variables store values.
- Simple expressions are combined into larger ones with operations like addition and subtraction.
- Basic types are collected into aggregates like arrays and structs.
- Expressions are used in statements whose execution order is determined by `control-flow` statements like `if` and `for`.
- Statements are grouped into functions for isolation and reuse.
- Functions are gathered into source files and packages.

### Names

There are 25 keywords that are reserved, like `if`, `for`, etc.

There are around 30 predeclared names like `int`, `true`, `nil`. They are not reserved. But, reusing them requires special cases and care.

Package names are always in lower case. If an entity start with an upper case then it is exported and available outside of the package.

Names should be short. However, if it is used outside of its own package, more expressive names can be used. The larger the scope of a name, the longer and more meaningful it should be.

Stylistically, camelCase is used, but not enforced. Acronyms and initialisms like ASCII and HTML are always used in the same case: **correct** `escapeHTML`, **incorrect** `escapeHtml`. 


### Declarations

A `declaration` names a `program entity` and `specifies` some or all of its `properties`.

There are 4 major type of declarations: `var`, `const`, `type`, `func`.

Local declarations are `visible` only within the function in which they are declared.

At the end of running a function, `control` and any results are then returned to the caller.


### Variables

The general variable declaration form: 

`var name type = expression` : The expression is the initial value.

Other options:

`var name = expression`: The type is determined by the initializer expression.

Or,

`var name type` : The initial value is the `zero value` for the type.

Zero values:
- `0` for numbers
- `false` for booleans
- `""`, empty string, for strings
- `nil` for interfaces and `reference types` (slice, pointer, map, channel, function)

The zero value of an aggregate type like an `array` or a `struct` has the zero value of all its elements or fields.

Zero value mechanism ensures that a variable always holds a well-defined value of its type.
There is no such thing as an uninitialized variable in Go, which prevents unpredictable errors.

Short variable declaration can only be used in functions.

`a := 5` equals to `var a int = 5`

If we wanted `a` to be a `float64` type with the same initial value we had to use the long declaration. See [TestTypes](./tests/misc_test.go).

`var a float64 = 5`

#### Pointers

A `variable` is a storage containing a value.

A `pointer value` is the address of a variable.

```
a := 5
p := &a  // get the address of a and assign it to p
*p = 6  // update the value of the variable pointed by p, which is a
a == 6  // true  - now the value is 6
```
See [TestPointer](./tests/misc_test.go).

Each component of aggregate type, a field of a struct or an element of an array, is also a variable and thus has an address too.

Expressions that denote variables are the only expressions to which the address-of operator `&` may be applied.

Two pointers are equal `==` if and only if they point to the same variable or both are `nil`.

A function can update a value of a passed argument as a side effect when the argument is a pointer to the this variable.

Creating a pointer to a variable is called `aliasing`. Aliasing also occurs when we copy values of reference types like slices, maps, and channels, and even structs, arrays, and interfaces that contain these types. For example if we pass a slice to a function as a variable (not a pointer), the function will/can manipulate the original slice. 
