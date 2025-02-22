![Website](https://img.shields.io/website?url=https%3A%2F%2Fsiruscodes.github.io%2FAntiLang%2F&style=flat&label=Online%20interpreter)
![GitHub Release](https://img.shields.io/github/v/release/SirusCodes/AntiLang?style=flat)

# AntiLang - breaking all the conventions

> [!CAUTION]  
> For your own sanity, please don't use this language.

All modern programming languages are very similar. All the ifs, elses, and loops look the same with some minor syntax changes.  
I was bored with them as they were not fun or interesting enough, hence **AntiLang**.

It **reverses** the structure of all the conditionals, loops, functions, everything, while keeping all the operators intact (I'm not that evil🙂). So get ready for some mind-bending coding that will make your brain do the FizzBuzz dance. 💃

Below is the syntax of the **FizzBuzz** program in **AntiLang**. 

> [!WARNING]  
> May cause headaches and/or existential crises.

```
{} main func [
    ,100 = count let
    ,0 = i let

    {count > i} while [
        {i % 3 == 0 && i % 5 == 0} if [
            ,{$It’s a FizzBuzz moment, boys!$}print
        ] {i % 3 == 0} if else [
            ,{$Fizz is life, Buzz is overhyped$}print
        ] {i % 5 == 0} if else [
            ,{$Buzzfeed has nothing on this$}print
        ] else [
            ,{$This is awkward... Why not just $ + i + $?$}print
        ]

        ,1 += i
    ]

    ,0 return
]

,{}main
```

Any self-respecting developer would skip the rest of the documentation and run the code above immediately. Don't worry, you can run AntiLang directly from your browser - [AntiLang](https://siruscodes.github.io/AntiLang/). See you soon, because you'll need it...

## Table of Contents
- [AntiLang - breaking all the conventions](#antilang---breaking-all-the-conventions)
- [How can I try it?](#how-can-i-try-it)
- [Run it](#run-it)
- [AntiLang has a REPL 🙀](#antilang-has-a-repl-)
- [Syntax](#syntax)
  - [Variable Declaration](#variable-declaration)
  - [Operators](#operators)
  - [Data Types](#data-types)
    - [String](#string)
    - [Array](#array)
    - [Map](#map)
  - [Functions](#functions)
  - [Built-in Functions](#built-in-functions)
  - [Conditional Flows](#conditional-flows)
  - [Loops](#loops)
- [Suggestions](#suggestions)
- [All the best](#all-the-best)

## How can I try it?

For those who prefer convenience (and sanity), use the [online interpreter](https://siruscodes.github.io/AntiLang/). (I've spent significant time on this, so please use it!)

I don't know why anyone would want to run this locally, but if you do (I'm not judging you), head over to the [release page](https://github.com/SirusCodes/AntiLang/releases) and get the latest binary for your system. Just remember, with great power comes great responsibility (and possibly a headache).

## Run it

Once you have [downloaded the binary](#how-can-i-try-it) as per your OS and ARCH you can run AntiLang programs by using `antilang run <filename>.al` (use `.al` as file extension).

You can download one of our [samples](https://github.com/SirusCodes/AntiLang/tree/main/sample)

```sh
curl -O https://raw.githubusercontent.com/SirusCodes/AntiLang/refs/heads/main/sample/fizzbuzz.al

./antilang run fizzbuzz.al
```

For Windows -

```pwsh
Invoke-WebRequest "https://raw.githubusercontent.com/SirusCodes/AntiLang/refs/heads/main/sample/fizzbuzz.al" -Outfile fizzbuzz.al

.\antilang.exe run .\fizzbuzz.al
```

## AntiLang has a REPL 🙀

To run REPL just run `antilang repl` and it should start REPL (Read Evaluate Print Loop).

## Syntax

### Variable Declaration

It's a dynamically typed language that supports int, string, map, and array (might add floats as well, but why do we need more types?).

```
,<value> = <var name> let
```

The declaration starts with a **comma** `,` (because commas are underrated) and ends with **let**. It's like a goodbye wave when you leave a party.

```
,10 = ten let
```

### Operators

I thought of keeping these the same as all the other languages, so `a + b` is actually **a + b**, not `a - b` (though I wanted to do that, but I’m not that evil, right?).

Supported operators:

- **+**, **-**, **/**, **\***, and **%** for arithmetic.
- **&&**, **||**, and **!** for logical operators.
- **=**, **+=**, **-=**, **/=**, and **\***= for assignment.
- **<**, **>**, **<=**, **>=**, **==**, and **!=** for comparisons.

PS: Assignment operators are reversed to maintain consistency with `let` statements. So, `1 += i` will **increment** `i` by 1 (yeah, I like to keep things spicy).

### Data Types

#### String

Strings are declared with the `$` sign because it's generally used for **string interpolation**, but I like to think of it as giving your strings a fresh haircut.

```
,$Hello Hell!$ = string let
```

#### Array

Arrays are declared with **`(`** and **`)`**, and values are **separated by `;`**. You know, just because it's cooler that way.

```
,(1; $Hello$) = array let
```

Indices start at **0** (don’t worry, we’re not that cruel). But, you have to specify the index before the variable name!

```
,(1)array
```

#### Map

Maps use **`[`** and **`]`** instead of `{}`. Assignment is done using **`=`** instead of `:`. Why? Because why not.

```
,[$Hello$ = 1; $Hell$ = 6] = map let
,($Hello$)map
```

### Functions

Functions are first-class citizens here in AntiLang. You can pass functions as parameters, like a boss (take that, Java!).

Here's how functions are declared (don’t worry, it gets weird):

```
{a; b}add func [
    ,a + b return
]

,{1; 1}add
```

To return values, we use:

```
,<value> return
```

### Built-in Functions

AntiLang has a small set of built-in functions, and I might add more in the future if you leave me some memes (or suggestions). So far, we support:

- `{array|string}len`: Returns the length of an array or string.
- `{array}first`: Returns the first element of an array.
- `{array}last`: Returns the last element of an array.
- `{array}rest`: Returns the array excluding the first element.
- `{array; element}push`: Adds an element to the end of an array.
- `{array}pop`: Removes the last element from an array.
- `{array; index; element}addAt`: Adds an element at a specified index in an array.
- `{array; index}removeAt`: Removes an element at a specified index in an array.
- `{value}print`: Prints the value to the console.

### Conditional Flows

Yes, we have `if`, `else`, and `else if` just like any normal language. But here, we like to add a little fun.

```
{a < b} if [
    ,{$b is greater$}print
] {a > b} if else [
    ,{$a is greater$}print
] else [
    ,{$They are the same$}print
]
```

### Loops

Right now, AntiLang only supports the **while** loop. I'm too lazy to implement the `for` loop. If this repo gets **100 stars**, I might just go ahead and do it. 🤞

Here’s an example of a **while loop**:

```
,0 = i let
{i <= 10} while [
    ,{i}print
    ,1 += i
]
```

### Suggestions

Do you have a better idea to make this language more interesting? Or just want to send a meme for the fun of it? [Open an issue](https://github.com/SirusCodes/AntiLang/issues/new) and let’s see what we can do to make coding **weirder and funnier**.

Or you can just send a meme. I’ll be fine with that too.

---

### All the best

- AntiLang is all about **reversing logic** but **keeping the operators intact** (well, almost).
- You’re going to be spending more time figuring out the structure than the logic of your program. #TrueCoderPain
- If you think the syntax is confusing, just remember: you're probably just not **AntiLang-ready** yet. It'll get easier (maybe). 😅
  
Go ahead, give it a [try](https://siruscodes.github.io/AntiLang/), and remember to keep your sanity in check. After all, **AntiLang** is not about getting the job done quickly; it's about having fun while losing your mind. 😜
