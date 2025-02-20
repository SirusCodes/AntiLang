const go = new Go();
WebAssembly.instantiateStreaming(fetch("antilang.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});

document.getElementById("runButton").addEventListener("click", () => {
    clearTerminal();

    const code = document.getElementById("codeArea").value;
    window.execute(code);
});

document.addEventListener("stdout", (e) => {
    const terminal = document.getElementById("terminal");
    const span = document.createElement("span");
    span.textContent = e.detail;
    terminal.appendChild(span);
    terminal.scrollTop = terminal.scrollHeight;
});

function clearTerminal() {
    const terminal = document.getElementById("terminal");
    terminal.innerHTML = "";
}

document.getElementById("codeArea").value = `{} main func [
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