,1 = i let

{i <= 15} while [
    {i % 3 == 0 && i % 5 == 0} if [
        ,{$FizzBuzz$}print
    ] {i % 3 == 0} if else  [
        ,{$Fizz$}print
    ] {i % 5 == 0} if else [
        ,{$Buzz$}print
    ] else [
        ,{i}print
    ]

    ,1 += i
]