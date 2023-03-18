import { useEffect, useRef, useState } from "react";
import axios from "axios";
import "./App.css";

const baseURL = "https://parakit.fly.dev";

function App() {
  const [summary, setSummary] = useState("");
  const [count, setCount] = useState(0);
  const inputref = useRef(null);
  function createSummarize() {
    axios
      .post(`${baseURL}/summarize/`, {
        text: inputref.current.value,
      })
      .then((res) => {
        //return res.data.paraphrased;
        setSummary(res.data.paraphrased);
        console.log(typeof res.data.paraphrased);
      })
      .catch((error) => {
        console.warn(error);
      });

    console.log(summary);
  }
  return (
    <div className="App">
      <h1>Parakit</h1>
      <div style={{ paddingRight: 40 }}>{count}/500</div>

      <div className="column">
        {/*Input */}

        <textarea
          ref={inputref}
          placeholder="Start typing or paste something here to paraphrase"
          cols={60}
          rows={30}
          maxLength={500}
          onChange={(event) => {
            setCount(inputref.current.value.length);
          }}
        />

        {/*Output */}
        <textarea value={summary} placeholder="" cols={60} rows={30} readOnly />
      </div>
      <button onClick={createSummarize}>Paraphrase</button>
    </div>
  );
}

export default App;
