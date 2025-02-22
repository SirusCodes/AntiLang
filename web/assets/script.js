const go = new Go();
WebAssembly.instantiateStreaming(fetch("antilang.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});

const terminal = document.getElementById("terminal");
const terminalContent = document.querySelector(".terminal-content");
const sampleSelector = document.getElementById("sampleSelector");

document.getElementById("runButton").addEventListener("click", () => {
    clearTerminal();
    terminal.classList.add("visible");

    const code = window.editor.getValue();
    window.execute(code);
});

window.samples = {
    "fizzbuzz.al": `{} main func [
    ,20 = count let
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

,{}main`,
    "loop.al": `{} main func [
    ,10 = count let
    ,0 = i let

    {count > i} while [
        ,{$Looping: $ + i}print
        ,1 += i
    ]

    ,0 return
]

,{}main`,
    "hello.al": `{} main func [
    ,{$Hello, World!$}print
    ,0 return
]

,{}main`,
    "add.al": `{a; b} add func [
    ,a + b return
]

,{{1; 1}add}print`,
    "blank.al": `{} main func [
    ,0 return
]

,{}main`,
};

for (const sample in samples) {
    const option = document.createElement("option");
    option.textContent = sample;
    option.value = sample;
    sampleSelector.appendChild(option);
}

sampleSelector.addEventListener("change", (event) => {
    const sample = event.target.value;
    const code = samples[sample];
    window.editor.setValue(code);
});

document.addEventListener("stdout", (e) => {
    const span = document.createElement("span");
    span.textContent = e.detail;
    terminalContent.appendChild(span);
    terminalContent.scrollTop = terminalContent.scrollHeight;
});

function closeTerminal() {
    terminal.classList.remove("visible");
}

function clearTerminal() {
    terminalContent.innerHTML = "";
}