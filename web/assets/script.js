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