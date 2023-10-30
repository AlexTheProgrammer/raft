const goWasm = new Go()

WebAssembly.instantiateStreaming(fetch("main.wasm"), goWasm.importObject)
.then((result) => {
    goWasm.run(result.instance)

    const body = document.getElementsByTagName("BODY")[0];
    body.innerHTML = getHtml();

})


