> [!NOTE]  
> It's a work in progress language with working lexer and parser

## AntiLang

> [!CAUTION]
> For your own sanity please don't use this language for more than 15-20mins

All the modern programming languages are very similar with all the if, else, loops looks the same with some minor syntax changes.

I was bored with them as they were not fun/interesting enough, hence AntiLang.

It reverses the structure of all the conditional. loops. functions. everything. while keeping all the operators intact (I'm not that evil🙂).

Below is the syntax of FizzBuzz program in AntiLang, I will work on complete docs with syntax highlighting and a basic website to run your code once I'm done writing the interpereter.

```
{p1; p2; p3} main func [
    ,100 = count let
    ,[$Hello$; $Hell !$] = array let

    ,{array(0)}print

    {i=+1, count > i, 1 = i let} for [
        {i % 3 == 0 && i % 5 == 0} if [
            {$FizzBuzz$}print,
        ] {i % 3 == 0} if else [
            {$Fizz$}print,
        ] {i % 5 == 0} if else [
            {$Buzz$}print,
        ] else [
            {i}print,
        ]
    ]

    5 + 2

    ,0 return
]
```