require(['vs/editor/editor.main'], function () {
  monaco.languages.register({id: 'alogl60v2'});
  monaco.languages.setMonarchTokensProvider('alogl60v2', {% include './website/monarch.js' %});
  brackets = [["[","]"], ["(",")"], ["'BEGIN'","'END'"]],
  monaco.languages.setLanguageConfiguration('alogl60v2', {
    surroundingPairs: [{open: '"', close: '"'}, {open: "'", close: "'"}],
    brackets: brackets,
    autoClosingPairs: [{open:':',close: '='},
    ...brackets.map(([o,c]) => ({open: o,close: c}))],
    folding: {markers: {
    // 'BEGIN'
    start: /\'[ \t\r\n]*B[ \t\r\n]*E[ \t\r\n]*G[ \t\r\n]*I[ \t\r\n]*N[ \t\r\n]*\'/,
    // 'END'
    end: /\'[ \t\r\n]*E[ \t\r\n]*N[ \t\r\n]*D[ \t\r\n]*\'/}},
    indentationRules: {
    // 'BEGIN'
    increaseIndentPattern: /\'[ \t\r\n]*B[ \t\r\n]*E[ \t\r\n]*G[ \t\r\n]*I[ \t\r\n]*N[ \t\r\n]*\'[ \t\r\n]*$/,
    // 'END'
    decreaseIndentPattern: /\'[ \t\r\n]*E[ \t\r\n]*N[ \t\r\n]*D[ \t\r\n]*\'[ \t\r\n]*;?[ \t\r\n]*$/,
    // 'DO' | 'THEN'
    indentNextLinePattern: /\'[ \t\r\n]*(D[ \t\r\n]*O)|(T[ \t\r\n]*H[ \t\r\n]*E[ \t\r\n]*N)[ \t\r\n]*\'$/}
  });

  const editor = monaco.editor.create(document.getElementById('editor'), {
    lineNumbers: 'on',
    scrollBeyondLastLine: false,
    theme: 'vs-dark',
    automaticLayout: true,
    fixedOverflowWidgets: true,
    minimap: {enabled: false},
    autoIndent: "full",
    unicodeHighlight: {allowedCharacters: {"×": true, "∨": true}},
    
    language: 'alogl60v2',
  });

  let prevDecorations = [];
  function onCompileErr(line, desc) {
    if (!(line >= 0)) {
      prevDecorations = editor.deltaDecorations(prevDecorations, []);
      return;
    }
    prevDecorations = editor.deltaDecorations(prevDecorations,[{
      range: new monaco.Range(line, 1, line, 1),
      options: {
        isWholeLine: true,
        className: 'errorContentClass',
        overviewRuler: {color: "#ff5959", darkColor: "#bc1616", position: monaco.editor.OverviewRulerLane.Right}
      }
    }]);
  }

  const exampleSources = ({{ example_sources_json }});

  let prevSelected;
  function onChangeSource() {
    if (prevSelected)
      exampleSources[prevSelected] = editor.getValue();
    const selected = document.querySelector("#example-source-select").value;
    editor.setValue(exampleSources[selected])
    prevSelected = selected;
  }
  document.querySelector("#example-source-select").onchange = onChangeSource;
  onChangeSource();

  document.querySelector('#target-ip').onclick = () => document.querySelector('#team-num').focus();

  async function compile() {
    const resultNode = document.querySelector("#result");

    const ip = "fd66:666:" + document.querySelector("#team-num").value + "::2";

    const compileResp = await fetch("http://[" + ip + "]:6061/compile", {
      method: 'POST',
      headers: {
        "Content-Type": "text/plain",
      },
      body: editor.getValue()
    });

    const resultText = await compileResp.text();
    const compileResult = JSON.parse(resultText);
    
    if ('error' in compileResult) {
      onCompileErr(compileResult.line, compileResult.error);
      resultNode.classList.add("failed");
      resultNode.innerText = "Compilation failed:\n" + compileResult.error;
      return [null, null, null];
    } else {
      onCompileErr(-1, "");
      resultNode.classList.remove("failed");
      resultNode.innerText = "Compiled successfully\n";
    }
    return [ip, resultText, compileResult];
  }

  document.querySelector('#execute').onclick = async () => {
    const resultNode = document.querySelector("#result");
    try {
      const [ip, resultText] = await compile();
      if (!resultText) return;

      const executeResp = await fetch("http://[" + ip + "]:6062/execute", {
        method: 'POST',
        headers: {
          "Content-Type": "application/json;charset=UTF-8",
        },
        body: resultText
      });

      const execResult = await executeResp.json();

      resultNode.innerText += "Execution Result: " + execResult.result + "\n";
      if (!execResult.stdout) {
        return;
      }
      // base64 to utf-8
      const stdout = decodeURIComponent(atob(execResult.stdout).split('').map(function(c) {
        return '%' + c.charCodeAt(0).toString(16).padStart(2, '0');
      }).join(''));
      resultNode.innerText += "Stdout: >>>\n" + stdout + "\n<<<";
    } catch(error) {
      resultNode.classList.add("failed");
      resultNode.innerText = "Failed: " + error;
    }
  };
  document.querySelector("#download-elf").onclick = async () => {
    const resultNode = document.querySelector("#result");
    try {
      const [_1, _2, compileResult] = await compile();
      if (!compileResult) return;

      console.log(compileResult.binary);

      var a = document.createElement("a");
      a.href = "data:application/octet-stream;base64," + compileResult.binary;
      a.download = "a.out";
      a.click();
    } catch(error) {
      resultNode.classList.add("failed");
      resultNode.innerText = "Failed: " + error;
    }
  };

  document.querySelector("#open-spec").onclick = () => {
    document.querySelector('#md-overlay').classList.remove("hidden");
    document.querySelector('#close-spec').classList.remove("hidden");
  };

  document.querySelector("#close-spec").onclick = () => {
    document.querySelector('#md-overlay').classList.add("hidden");
    document.querySelector('#close-spec').classList.add("hidden");
  };
});
