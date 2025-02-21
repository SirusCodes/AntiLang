const go = new Go();
WebAssembly.instantiateStreaming(fetch("antilang.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});

const terminal = document.getElementById("terminal");
const terminalContent = document.querySelector(".terminal-content");

document.getElementById("runButton").addEventListener("click", () => {
    clearTerminal();
    terminal.classList.add("visible");

    const code = window.editor.getValue();
    window.execute(code);
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