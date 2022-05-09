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
- Simple expressions* are combined into larger ones with operations like addition and subtraction.
- Basic types are collected into aggregates like arrays and structs.
- Expressions are used in statements whose execution order is determined by `control-flow` statements like `if` and `for`.
- Statements are grouped into functions for isolation and reuse.
- Functions are gathered into source files and packages.

`Operand` is an `expression*` that yields a value that an `operator` operates on.


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

The `new` function creates an `unnamed variable` of type T, initializes it to the zero value of T, and returns its address, which is a value of type *T.

Followings are identical:
```Go
func newInt() *int {
    var i int
    return &i
}

func newInt() *int {
    return new(int)
}
```

#### Lifetime of a variable

The lifetime of package-level variable is the entire execution of the program.

Local variables live on from the declaration until it becomes unreachable, at which point its storage may be `recycled`.

Pointers and other kind of references that ultimately lead to a variable keeps a local variable alive. When no such path exists anymore, then the variable becomes unreachable. It can no longer affect the rest of the computation.

The Go compiler decides where to store a variable, `stack` or `heap`. But, if a local variable is referenced by a global pointer, then this local variable escapes/outlives the function it has been created. It keeps living on, thanks to the global pointer, on the heap memory. Stack is used for short lived variables.

`Tip`: Do not keep pointers to short-lived objects in long-lived objects if it is not needed. This will prevent the garbage collector from reclaiming the short-lived objects.

#### Assignment

Each of the arithmetic and bitwise binary operators has a corresponding `assignment operator`, i.e., `*=`.

Assignment operators save us from having to repeat and re-evaluate the expression for the variable.
```Go
count[x] = count[x]*5 // re-evaluate the `count[x]` expression

count[x] *= 5  // no re-evaluation of the expression
```
`Tuple assignment` is used when assigning value to multiple variables in one line, i.e., reading 2 values (1 is error) from a function. Other use case is swapping the values of two variables. This is possible due to evaluating the right hand side before updating the left hand side.
```Go
a, b = b, a // their values are swapped without needing a temp variable
```

Assignment is `legal` only if the value is assignable to the type of the variable.

An `implicit assignment` happens when a variable gets assigned a value without using the `=` operator. For example, when a function is called, its parameters get assigned with the argument values implicitly.

`nil` may be assigned to any variable of interface or reference type.

A `type declaration` defines a new named type that has the same underlying type as an existing type.
```Go
type name underlying-type

type Celsius float64
type Fahrenheit float64
```

Type declaration prevents inadvertent errors like combining incompatible values (Celsius and Fahrenheit) in an arithmetic operation. Even though they have the same underlying-type, `float64`, they can not be compared or combined in arithmetic operations.

A conversion `T(x)` - converting x into T- is allowed when they both have the same underlying-type. 

This conversion does not change the representation of the value: `Celsius(34.5) => 34.5`

A conversion never fails at run time. Converting float to int results in loosing the fraction part. Converting string to slice of byte allocates a copy of the string data.

Creating a `String()` method on the named type will enable `%v` and `%s` conversion characters to represent the named value as defined in the `String()` method.

```Go
type Celsius float64

func(c Celsius) String() string {return fmt.Sprintf("%g°C", c)}

c := 100

fmt.Println(c) // 100°C
fmt.Printf("%v", c) // 100°C
```

#### Packages and Files

Packages in Go supports modularity, encapsulation, separate compilation, and reuse.

To refer to a function from outside its package, we must `qualify` the identifier to make it explicit.

Exported identifiers start with an upper-case letter.

Only one file in each package should have a package doc comment. Extensive doc comments are often placed in a file of their own, conventionally called `doc.go`.

Every package is identified by a unique string called its `import path`. The Go language does not define where these strings come from or what they mean; it's up to the tools to interpret them. For `go` tool, it denotes a directory containing one or more Go source files that together make up the package.

By convention, a package's name matches the last segment of its import path.

Package initialization begins with by initializing package-level variables in the order in which they are declared, except that dependencies are resolved first.

`go` tool sorts files in a package by their names before giving it to the compiler.

If there is a complex logic at package initialization, the `init function` can be used. A init function can not be called, otherwise it is a normal function. A package can have multiple init functions. They will be executed in the order in which they are declared.

Packages are initialized in the order of import, dependencies first. So, if package p imports q, then q will be fully initialized before p. The main package is the last to initialize.

#### Scope

The scope of a declaration is the part of the source code where a use of the declared name refers to that declaration.

The `scope` should not be confused with the lifetime. Scope is a `compile-time property` that represents the region of the program text. The `lifetime` of a variable is a `run-time property` which is the duration that a variable can be referred to by other parts of the program.

A `syntactic block` is a sequence of enclosed statements enclosed in braces like those that surround the body of a function or loop.

A name declared inside a syntactic block is not visible outside that block.

A `lexical block` is not bound by braces. The lexical block for the entire source code is called the `universe block`. 

The built-in types, functions are in the universe block and can be referred to throughout the entire program.

Declaration outside functions, that is, at `package level`, can be referred to from any file in the same package.

Imported packages are only available in the file that they are imported to, `file level`.

`Shadowing`: When the compiler encounters a reference to a name, it looks for a declaration, starting from the innermost enclosing lexical block and working up to the universe block. If there are 2 variables with the same name where 1 is local and the other one is global, or lives in a larger scope, the local one wins/shadows/hides the other one.

Normal practice in Go is do deal with the error in the if block and then return, so that the successful execution path is not indented.
```Go
// Good
f, err := os.Open(fileName)
if err != nil {
    return err
}
// do stuff with f

//Not great
if f, err := os.Open(fileName); err!= nil {
    return err
} else {
    //do stuff with f
}
```


## Basic Data Types

Computers operate fundamentally on fixed-size numbers called `word`s.

There are four categories of types: 
- Basic: numbers, strings, and booleans
- Aggregate: array and struct 
- Reference: map, slice, channel, pointer, functions 
- Interface

`int` and `uint` have the same size, either `32` or `64`.

`rune` holds a Unicode code point and it is a synonym for `uint32`.
`byte` is a raw data and it is a synonym for `uint8`.
They may be used interchangeably. 

Unsigned numbers tend to be used only when their bitwise operators or peculiar arithmetic operators are required. 
They are typically not used for merely non-negative quantities.

`float64` should be used instead of `float32`, unless there is a good reason.

Digits maybe omitted when writing a float number, i.e., `.4` or `34.`.

The `complex(real, imag)` function can be used to construct a complex number.
There are two types: `complex64` and `complex128`.

Short-circuit behavior: If the answer is already determined by the value of the left operand (expression), `<left_operand> && <right_operand>`, the right operand is not evaluated.

`&&` has higher precedence than `||`, because `&&` is Boolean multiplication, `||` is Boolean addition.

### String

A `string` is an `immutable` sequence of bytes.

A string usually contains a `human readable text`.

Text strings are conventionally interpreted as UTF-8 encoded sequences of Unicode code points (runes).

The `len` function returns the `number of bytes` in a string.

`s[i]` returns the `i-th` byte in the string `s`. But this does not mean that it will return the `i-th` character.
Because, the [UTF-8](https://en.wikipedia.org/wiki/UTF-8) encoding of non-ASCII code point requires two or more bytes. See the [TestStrings](./tests/misc_test.go) test.

The substring operation `s[2:5]` yields a `new` string including the second byte but not including the fifth byte.

The `+` operator makes a `new` string by concatenating two strings.

String values are `immutable`. But a string variable can be updated. In the background a new string is created instead of mutating its value.

Because strings are immutable, this is not allowed: `s[3] = "L"`.

Immutability enables sharing the same underlying string between multiple variables that have the same or similar (partially same - substring) values. See page 66 for the illustration.










