import { Editor} from "./apps/editingArea/Editor";
import "./App.css"

const send = (sendingData)=>{
  useEffect(() => {
     fetch("http://localhost:2000/oo", {
        method: 'POST',
        body: JSON.stringify(sendingData),
        headers: {
           'Content-type': 'application/json; charset=UTF-8',
        },
        mode: "no-cors"
     })
        .then((response) => response.json())
        .then((data) => {
           console.log(data);
           // Handle data
        })
        .catch((err) => {
           console.log(err.message);
        });
  }, []);
}

function App() {

  return (
    <>
      <div id="mainWindow">
        <Editor/>
      </div>
    </>
  )
}

export default App
