# AntiLang

> [!CAUTION]
> For your own sanity please don't use this language.

All the modern programming languages are very similar with all the if, else, loops looks the same with some minor syntax changes.

I was bored with them as they were not fun/interesting enough, hence AntiLang.

It reverses the structure of all the conditional. loops. functions. everything. while keeping all the operators intact (I'm not that evil🙂).

Below is the syntax of FizzBuzz program in AntiLang, I will work on complete docs with syntax highlighting and a basic website to run your code once I'm done writing the interpereter.

```
{} main func [
    ,100 = count let
    ,0 = i let

    {count > i} while [
        {i % 3 == 0 && i % 5 == 0} if [
            ,{$FizzBuzz$}print
        ] {i % 3 == 0} if else [
            ,{$Fizz$}print
        ] {i % 5 == 0} if else [
            ,{$Buzz$}print
        ] else [
            ,{i}print
        ]

        ,1 += i
    ]

    ,0 return
]
```

## Syntax

### Variable declaration

It's a dynamically typed language with supports int, string, map and array (might add floats well)

```
,<value> = <var name> let
```

The declaration should start with a `,` and end with `let`.

```
,10 = ten let
```

### Operators

I thought of keeping this same as all the other languages like hence it does exactly what is looks like. So `a + b` is actually `a + b` not `a - b`.

PS: Doing otherwise was the plan.

AntiLang Supports:

- +, -, /, *, and % arithmetic operators
- &&, ||, and ! logical operators
- =, +=, -=, /=, and *= assignment operators
- <, >, <=, >=, ==, and != comparison operators

> [!Note]
> I lied, assignment operators are reversed to maintain consistency with let statements. `1 += i` will be increment i by 1 and similarly for others. 

### Data types

#### String

Strings are declared using `$` sign (since it's generally used for string interpolation).

```
,$Hello Hell!$ = string let
```

#### Array

I hope you might have assumed it that array would be declared with a `(` and `)`, and separated by `;`.

```
,(1; $Hello$) = array let
```

Index starts with **0** (come on I'm not that evil) but you have to specify the index before the variable name.

```
,(1)array
```

#### Map

Maps use `[` and `]` instead of `{` and `}` and assignment is done using `=` instead of `:`.

Again I have kept mapping to be **key=value** pairs and similar to array you have to pass the keys before the variable. 

```
,[$Hello$ = 1; $Hell$ = 6] = map let
,($Hello$)map
```

### Functions

In AntiLang Functions are first class citizen, means you can pass functions as parameters. Surprised? such a stupid langugage can do it and Java couldn't (which runs on more than 3 billion devices)

As we know syntax will be a weird.

```
{a; b}add func [
    ,a + b return
]

,{1; 1}add
```

and obviously you return like that...

```
,<value> return
```

### Builtin Functions

AntiLang have a support for a bunch of builtin function. I might add a few more - just [open an issue](https://github.com/SirusCodes/AntiLang/issues/new) for it.

Currently it supports:

- `{array|string}len`: Returns the length of an array or string.
- `{array}first`: Returns the first element of an array.
- `{array}last`: Returns the last element of an array.
- `{array}rest`: Returns the array excluding the first element.
- `{array; element}push`: Adds an element to the end of an array.
- `{array}pop`: Removes the last element from an array.
- `{array; index; element}addAt`: Adds an element at a specified index in an array.
- `{array; index}removeAt`: Removes an element at a specified index in an array.
- `{value}print`: Prints the value to the console.

### Conditonal Flows

We support if, else and else if conditional flows.

```
{a < b} if [
    ,{$b is greater$}print
] {a > b} if else [
    ,{$a is greater$}print
] else [
    ,{$They are same$}print
]
```

### Loops

Currently AntiLang only supports `while` loop (I'm lazy to implement `for` loop). If this repo gets 100 stars I will do it 🤞

```
,0 = i let
{i <= 10} while [
    ,{i}print
    ,1 += i
]
```

## Suggestions

Do let me know if you think there is a way to make this language a bit more interesting and fun by raising a [suggestion issue](https://github.com/SirusCodes/AntiLang/issues/new).
