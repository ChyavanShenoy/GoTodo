import TodoContainer from "./Components/TodoContainer";
import { useEffect, useState } from "react";

function App() {
  const [todos, setTodos] = useState([]);

  function getTodos() {
    fetch("http://localhost:8080/api/v1/getTodos")
      .then((response) => response.json())
      .then((data) => {
        console.log("Todos fetched successfully");
        setTodos(data);
      })
      .catch((error) => {
        console.error("Error fetching todos:", error);
      });
  }

  useEffect(() => {
    getTodos();
  }, []);

  return (
    <div className="bg-gray-800 min-h-screen flex justify-begin">
      <TodoContainer todos={todos} />
    </div>
  );
}

export default App;
