,1 = i let

{i < 16} while [
    {i % 3 == 0 && i % 5 == 0} if [
        ,{$FizzBuzz$}print
    ] else [
        {i % 3 == 0} if [
            ,{$Fizz$}print
        ] else [
            {i % 5 == 0} if [
                ,{$Buzz$}print
            ] else [
                ,{i}print
            ]
        ]
    ]

    ,1 += i
]