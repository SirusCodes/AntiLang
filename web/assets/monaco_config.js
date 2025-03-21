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
        value: window.samples["fizzbuzz.al"],
        language: "antilang",
        theme: "antilangTheme"
    });

    // Hide the loading spinner once the editor is loaded
    document.getElementById("loadingSpinner").style.display = "none";
});
