require.config({ paths: { vs: 'https://cdnjs.cloudflare.com/ajax/libs/monaco-editor/0.52.2/min/vs' } });

require(['vs/editor/editor.main'], function () {
    // Register a new language
    monaco.languages.register({ id: "antilang" });

    monaco.languages.setLanguageConfiguration("antilang", {
        brackets: [
            ['{', '}'],
            ['[', ']'],
            ['(', ')']
        ],
        autoClosingPairs: [
            { open: '{', close: '}' },
            { open: '[', close: ']' },
            { open: '(', close: ')' },
            { open: '$', close: '$' },
        ],
        surroundingPairs: [
            { open: '{', close: '}' },
            { open: '[', close: ']' },
            { open: '(', close: ')' },
            { open: '$', close: '$' },
        ]
    });

    // Register a tokens provider for the language
    monaco.languages.setMonarchTokensProvider("antilang", {
        keywords: [
            'let', 'func', 'while', 'return', 'null', 'if', 'else', 'true', 'false'
        ],

        operators: [
            '+',
            '-',
            '*',
            '/',
            '%',
            '+=',
            '-=',
            '*=',
            '/=',
            '&&',
            '||',
            '==',
            '<',
            '>',
            '=',
            '!',
            '!=',
            '<=',
            '>=',
            '(',
            ')',
            ']',
            '{',
            '}',
            ',',
            ';',
        ],

        // we include these common regular expressions
        symbols: /[=><!~?:&|+\-*\/\^%]+/,
        escapes: /\\(?:[abfnrtv\\"']|x[0-9A-Fa-f]{1,4}|u[0-9A-Fa-f]{4}|U[0-9A-Fa-f]{8})/,

        // The main tokenizer for our languages
        tokenizer: {
            root: [
                // identifiers and keywords
                [
                    /[a-zA-Z_]\w*/,
                    {
                        cases: {
                            '@keywords': { token: 'keyword.$0' },
                            '@default': 'identifier'
                        }
                    }
                ],

                // whitespace
                { include: '@whitespace' },

                // delimiters and operators
                [/[{}()\[\]]/, '@brackets'],
                [/[<>](?!@symbols)/, '@brackets'],
                [
                    /@symbols/,
                    {
                        cases: {
                            '@operators': 'delimiter',
                            '@default': ''
                        }
                    }
                ],

                // numbers
                [/\d*\d+[eE]([\-+]?\d+)?/, 'number.float'],
                [/\d*\.\d+([eE][\-+]?\d+)?/, 'number.float'],
                [/\d[\d']*/, 'number'],
                [/\d/, 'number'],

                // delimiter: after number because of .\d floats
                [/[;,.]/, 'delimiter'],

                // strings
                [/\$([^\$\\]|\\.)*$/, 'string.invalid'], // non-teminated string
                [/\$/, 'string', '@string'],
            ],

            whitespace: [
                [/[ \t\r\n]+/, ''],
                [/\/\*\*(?!\/)/, 'comment.doc', '@doccomment'],
                [/\/\*/, 'comment', '@comment'],
                [/\/\/.*$/, 'comment']
            ],

            comment: [
                [/[^\/*]+/, 'comment'],
                // [/\/\*/, 'comment', '@push' ],    // nested comment not allowed :-(
                // [/\/\*/,    'comment.invalid' ],    // this breaks block comments in the shape of /* //*/
                [/\*\//, 'comment', '@pop'],
                [/[\/*]/, 'comment']
            ],
            //Identical copy of comment above, except for the addition of .doc
            doccomment: [
                [/[^\/*]+/, 'comment.doc'],
                // [/\/\*/, 'comment.doc', '@push' ],    // nested comment not allowed :-(
                [/\/\*/, 'comment.doc.invalid'],
                [/\*\//, 'comment.doc', '@pop'],
                [/[\/*]/, 'comment.doc']
            ],

            string: [
                [/[^\\\$]+/, 'string'],
                [/@escapes/, 'string.escape'],
                [/\\./, 'string.escape.invalid'],
                [/\$/, 'string', '@pop']
            ],
        }
    });

    // Register a completion item provider for the new language
    monaco.languages.registerCompletionItemProvider("antilang", {
        provideCompletionItems: (model, position) => {
            var word = model.getWordUntilPosition(position);
            var range = {
                startLineNumber: position.lineNumber,
                endLineNumber: position.lineNumber,
                startColumn: word.startColumn,
                endColumn: word.endColumn,
            };
            var suggestions = [
                {
                    label: "simpleText",
                    kind: monaco.languages.CompletionItemKind.Text,
                    insertText: "simpleText",
                    range: range,
                },
                {
                    label: "testing",
                    kind: monaco.languages.CompletionItemKind.Keyword,
                    insertText: "testing(${1:condition})",
                    insertTextRules:
                        monaco.languages.CompletionItemInsertTextRule
                            .InsertAsSnippet,
                    range: range,
                },
                {
                    label: "ifelse",
                    kind: monaco.languages.CompletionItemKind.Snippet,
                    insertText: [
                        "if (${1:condition}) {",
                        "\t$0",
                        "} else {",
                        "\t",
                        "}",
                    ].join("\n"),
                    insertTextRules:
                        monaco.languages.CompletionItemInsertTextRule
                            .InsertAsSnippet,
                    documentation: "If-Else Statement",
                    range: range,
                },
            ];
            return { suggestions: suggestions };
        },
    });

    monaco.editor.defineTheme('antilangTheme', {
        base: 'vs-dark',
        inherit: true,
        rules: [
            { token: 'comment', foreground: '6A9955' },
            { token: 'keyword', foreground: '569CD6' },
            { token: 'identifier', foreground: '9CDCFE' },
            { token: 'number', foreground: 'B5CEA8' },
            { token: 'string', foreground: 'CE9178' },
            { token: 'delimiter', foreground: 'D4D4D4' },
            { token: 'delimiter.bracket', foreground: 'D4D4D4' },
        ],
        colors: {
            'editor.foreground': '#E0E0E0',
            'editor.background': '#1E1E1E',
            'editorCursor.foreground': '#AEAFAD',
            'editor.lineHighlightBackground': '#2C2C2C',
            'editorLineNumber.foreground': '#858585',
            'editor.selectionBackground': '#264F78',
            'editor.inactiveSelectionBackground': '#3A3D41',
        }
    });

    window.editor = monaco.editor.create(document.getElementById("codeArea"), {
        value: getCode(),
        language: "antilang",
        theme: "antilangTheme"
    });
});

function getCode() {
    return `{} main func [
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

,{}main`
}
