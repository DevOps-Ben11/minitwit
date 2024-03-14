import { useEffect, useState } from "react";
import "../App.css";
function App() {
  const [count, setCount] = useState(0);
  const [latest, setLatest] = useState(null);

  /* I want to receive the list of messages but it doesn't work for now */
//   const [messages, setMessages] = useState([]);

//   useEffect(() => {
//     fetch('/sim/msgs', {
//         headers: {
//         'Authorization': 'Basic c2ltdWxhdG9yOnN1cGVyX3NhZmUh'
//       }})
//         .then((response) => response.json())
//         .then((data) => setMessages(data));
//   }, []);

  useEffect(() => {
    fetch("/sim/latest")
      .then((response) => response.json())
      .then((data) => setLatest(data.latest));
  }, []);

  return (
    <>
        <h1>Hello from Home</h1>

        <div>messages: {/*messages doesn't work for now, it doesn't receive the list of messages.*/}</div>
        <div>Latest: {latest}</div>

        <div className="card">
            <button onClick={() => setCount((count) => count + 1)}>
            count is {count}
            </button>
            <p>
            Edit <code>src/App.tsx</code> and save to test HMR
            </p>
        </div>
        <p className="read-the-docs">
            Click on the Vite and React logos to learn more
        </p>
    </>
  );
}

export default App;
