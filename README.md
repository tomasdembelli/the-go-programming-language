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

`String literal`: A sequence of bytes enclosed in double quotes.

We can include Unicode code points in string literals because Go source files are always encoded in UTF-8.

`Escape sequence`s can be used to insert arbitrary byte values into string. They start with backslash, `\n` (newline) , `\t` (tab).

A `raw string literal` is written between backquotes, ``some raw text``.
Escape sequences will be ignored in raw string literals. The backslashes will be represented as they are, not as an escape character.

Unicode assigns a standard number called a `Unicode code point`, in Go terminology a `rune`, to every characters in every language in the world.

Representing a Unicode code point requires an int32. But, since most of the time the meaningful part can still fit in 1 byte (which is for the ASCII characters), it would be waste of memory if we use int32 every time. To fix that issue, UTF-8 has been developed, which is a variable-length encoding that can be 1 to 4 bytes depending on the character encoded. UTF-8 is ASCII backward compatible.

[`Rune literal`](https://go.dev/ref/spec#Rune_literals)s are written in single quote.

Looping over a string:
- character by character (rune == character in Go):
```Go
	for i, r := range "€"  {  // r is a rune (int32)
		fmt.Printf("%d, %q, %d\n",i, r, r)
	}
    //outputs (only 1 iteration, since there is 1 characters)
    0, '€', 8364
```
- byte by byte:
```Go
    euroSign := "€"
	for i:=0; i < len(euroSign); i++  {
        // r is a byte, because string is a sequence of bytes
        r := euroSign[i]  
		fmt.Printf("%d, %q, %d\n",i, r, r)
	}
    // outputs (3 iterations due to 3 bytes)
    0, 'â', 226
    1, '\u0082', 130
    2, '¬', 172
```

See page 70 for an illustration.

Because strings are immutable, building up strings incrementally can involve a lot of allocation and copying.
In such cases, it is more efficient to use `bytes.Buffer` type.

The `bytes` package has many functions that the `strings` package have. `Contains, Count, Fields, HasPrefix, Index, Join`. 
This allows us to convert a string into a byte slice and do manipulation on it more efficiently.
Because a slice is not immutable and it won't create a copy of the data at every manipulation as the `strings` package would do for the immutable strings.

`bytes.Buffer` can be used for constructing strings. It does not require initialization, because its zero value is usable.

`strconv` can be used to convert a string in to number and vice a versa. `Itoa` => Integer to ASCII

`Constants` are `expressions` whose value is known to the `compiler` and whose evaluation is guaranteed to occur at compile time. The underlying type of every constant is a `basic type: boolean, string, or number`.

`Iota`, `constant generator`, begins at zero and increments by one for each item in the sequence.

Many constants are not committed to a particular type.

By deferring this commitment, `untyped constants` not only retain their `higher precision` until later, but they can participate in many more expressions than committed constants `without requiring conversions`.

## Composite Types

Basic types are the atoms of our (Go) universe.

Composite types are the molecules created by combining the basic types in various ways.

Arrays and structs are aggregate types and their size is fixed. Also, `arrays are homogeneous`, whereas `structs are heterogeneous`.

`Slices and maps are dynamic data structures` that grow as values are added.

Array contains one particular type only. Its elements can be accessed via `subscript notation`, where subscripts run from zero to one less than the array length.

We can use an `array literal` to initialize an array with a list of values.

Usage of `ellipsis`:
`a := [...]int{1,2}`, The length of the array `a` is determined by the number of initializers.

The size of an array is part of its type.
```Go
a := [2]int{1,2}
a = [3]int{1,2,3}
````
This returns an [incompatible assign](https://pkg.go.dev/golang.org/x/tools/internaltypesinternal#IncompatibleAssign)

It is possible to specify the index of elements while initializing an array,

`a := [...]int{99: 5}` This will initiate an array in size of 100. Except the last element, all other elements values will be `0`.

Arrays can be compared if their element types are comparable and their sizes are the same.

When a function is called, a copy of each argument vale is assigned to the corresponding parameter variable, so he function receives a copy, not the original.

`Slices` represent variable-length sequences whose elements all have the same type.

A slice is a lightweight data structure that gives access to its underlying array.

A slice has 3 components; a `pointer`, a `length`, and a `capacity`. Pointer points to the first element of the slice (not necessarily the array's first element).The length is the number of slice elements, it can't exceed the capacity. The capacity is usually the number of elements between the start of the slice and the end of the underlying array.

Unlike arrays, slices are not comparable. So, we cannot use == operator on slices. We need to write a custom comparison function depending on the use case. 

The zero value of a slice type is nil.

```Go
var s []int // len(s) == 0, s == nil
s := nil // len(s) == 0, s == nil
s = []int(nil) // **conversion expression** len(s) == 0, s == nil
s = []int{} // len(s) == 0, s != nil
```

The `make` functions creates an unnamed array variable and returns a slice of it; the array is accessible only through the returned slice.

We can not know that whether a given call to `append` will cause a reallocation, 
so we can't assume that the original slice refers to the same array as the resulting slice, nor that it refers to a different one.

Slices are not `pure` reference type but resemble an aggregate type such as this struct:
```Go
type IntSlice struct {
    ptr *int
    len, cap int
}
```

`Variadic function`s accept any number of final arguments.

A hash table is an unordered collection of `key:value` pairs.

In Go, a `map` isa reference to a hash table.

The `key` type must be a `comparable` with `==` operator.

A map element is nota variable, so we canot take its address.
```Go
_ = &ages["bob"] // compile error: cannot take address of map element
```

The order of map iteration is unspecified.
We can access the map elements by sorting the keys beforehand (store them in a slice),
and accessing the map values via sorted keys.

Storing to a `nil` map causes a panic.
Other operations, `len`, `delete`, lookup, `range` are safe to be applied on `nil` maps.

Go does not provide a set type.
But, `map`s can be used instead. `map[T]bool`.
Map keys have to be unique.

`Struct` is an aggregate data type that groups together zeao or more named values of arbitrary types as a single entity.
Each value is called a field.

Field order is `significant` to the type identity.

A struct type may contain a mixture of exported and unexported fields.

A named struct type `S` can't declare a field of the same type `S`: an aggregate value cannot contain itself.

The zero value for a struct is composed of the zero values of each of its fields.

Empty struct has zero size.

Go is a `call-by-value` language.

If all the fields ofa struct are comparable, the struct itself is comparable.
It can also be used as a map key.

Anonymous fields are the fields in structs with type but no names.
The type must be a named type or a pointer to a named type.

In nested structs, the outer struct type gains methods of embedded types.
`Composition` is the central to object-oriented programming in Go.

`marshaling`: Converting a Go data structure to JSON is called `marshaling`.
`json.MarchalIndent` can be used to pretty print.

Only exported fields are marshalled.

A field tay is a string of metadata associated at compile time with the field of a struct.

`omitempty`: No JSON output should be produced if the field has the zero value for its type.

`unmarshaling`: Decoding JSON to Go data structure.

The matching process that associates JSON names with Go struct names during unmarshaling is case-insensitive.

A `template` is a string or a file containing one or more portions enclosed in double braces `{{..}}`, called `action.`

The pipe operator `|` can be used in actions.

## Functions

A function lets us wrap up a sequence of statements in a unit that can be called from elsewhere in a program, perhaps multiple times.

A function hides it's implementation details from its users.
```go
func name(parameter-list)  (result-list) {
	body
}
```
The parameter list specifies the names and types of the function's parameters, which care the local variables whose values or arguments are supplied by the caller.

Leaving off the result list entirely declares a function that does not return any value and is called only for its effects.

Like parameters, results may be named. In that case, each name declares a local variable initialized to the zero value for its type.

Two functions have the same type or signature if they have the same sequence of parameter types and the same sequence of result types.

Go has no concept of default parameter values, nor any way to specify arguments by name, so the names of parameters and results don't matter to the caller except as documentation.

Arguments are passed by values to the functions. So, they receive a copy of each argument. If the argument contains a reference type variable like slice, map, any modification to the argument will have an impact on the referenced value.

A function is recursive if it is calling itself either directly or indirectly.

The `golang.ord/x/...` repositories are not in the standard library either because they are still under development or they are not needed by the majority of Go programmers. They are maintained by the Go team.

Many programming language implementations use a fixed-size function call stack. Fixed-size stack impose a limit on the depth of recursion.
In contract, typical Go implementations use variable-size stacks that start small and grow as needed up to a limit on the order ofa gigabyte.

Go's garbage collector recycles unused memory, but do not assume it will release unused operating system resources like open files and network connections. 
They should be closed explicitly.

Conventionally, a final bool result from a function indicates success; an error result often needs no explanation.

In a function with named results, the operands of a return statement may be omitted. This is called `bare return`.
Bare returns are not encouraged to use due to difficulty of making sense out of the code while reading by a human.

A function for which failure is an expected behavior returns an additional result, conventionally the last one. 
If the failure has only one possible cause, the result is a boolean.
If the failure may have a variety of causes for which the caller will need an explanation, the type of the additional result is error.

The built-in type `error` is an `interface` type. If it is `nil`, that means `success`.

Usually when a function returns a non-nil error, its other results are undefined and should be ignored. However, some functions may return partial result, if this is the case, this should be well documented.

Go programs use ordinary control-flow mechanism like if and return to respond to errors. This style undeniably demands that more attention be paid to error-handling logic, which is precisely the point.

When a function call returns an error, it's the caller's responsibility to check it and take appropriate action.

Error propagation means making a failure in a subroutine a failure of the calling routine.

Because error messages are frequently chained together, message strings should not be capitalized and newlines should be avoided.

In general, the call `f(x)` is responsible for reporting the attempted operation `f` and the argument value `x` as they relate to the context of the error.

We can `retry` the failed operation, possibly with a delay between tries, and perhaps with a limit on the number of attempts or the time spent trying before giving up entirely.

Stopping the entire program should be reserved to main package. And, it should only happen when it is impossible to recover from a failure.

Generally, after checking an error, failure is usually dealt with before success. So that success section does not handled in an else statement with an indentation.

Functions are `first-class values` in Go. They have types, and they may be assigned to variables or passed to or returned from functions.

The zero value of a function type is nil. Calling a nil function value causes a panic.

`strings.Map` applies a function to each character of a string, joining the results to make another string

The `*` adverb prints a string padded with a variable number of spaces when using `fmt.Printf`.
```go
fmt.Printf("%*s</%s>\n", 2, "", n.Data)
```

Named functions can be declared only at the package level. We cannot use nested functions.

A `function literal` is written like a function declaration, but without a name following the func keyword. It is an expression, and its value is called an anonymous function.

When an anonymous function requires recursion, we must first declare a variable, and then assign the anonymous function to that variable.

Declare a local variable in for loops if the captured values have to be frozen and used after exiting the for loop.

A `variadic function` is one that can be called with varying numbers of arguments.

Slice elements can be flatten with a preceding ellipsis when calling a variadic function.
```go
values := []{1,2,3,4}
fmt.Println(values...)
```

Defer functions are ordinary functions. 
Their arguments are evaluated when they are declared, but they are run when the outer function finishes. 
This includes panic, too. Multiple defer statements run in LIFO.

The right place for a defer statement that releases a resource is immediately after the resource has been successfully acquired.

Do not use defer statement in for loop or be very careful. It could result in resource leak, i.e., file descriptors.

When the Go runtime detects mistakes like out-of-bounds array access or nil pointer dereference, it panics.

During a panic, normal execution stops, all deferred function calls in that gorutine are executed, and the program crashes with a log message.

The panic function can be used when some impossible situation happens, for instance, execution reaches a case that logically can't happen.
Intentional panic should be only used for logical inconsistencies. 
In a robust program, expected error, the kind that arise from incorrect input, misconfiguration, or failing I/O should be handled gracefully; they are best dealt with using error values.

When a panic occurs, all deferred functions are run in reverse order.

The recover function must be used in deferred call. It will return the value/cause of the panic.

Entire stack can be written out with `runtime.Stack`.

We should not attempt to recover from another package's panic.

We can recover from a panic selectively. We can check the panic reason, if it is a known/expected issue then we can return an error, otherwise continue panicking.

From some conditions there is no recovery. Running out of memory, for example, causes the Go runtime to terminate the program  with a fatal error.








