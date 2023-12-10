import React, { useState } from "react";

const AddTodo = () => {
  const [todo, setTodo] = useState("");

  const handleInputChange = (e) => {
    setTodo(e.target.value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // Add your logic here to handle the submission of the todo data
    console.log("Todo:", todo);
    setTodo("");

    var dataToSend = {
      id: 1,
      title: todo,
      completed: false,
      createdAt: new Date(),
      updatedAt: new Date(),
    };

    // send json data to backend
    fetch("http://localhost:8080/api/v1/addTodo", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(dataToSend),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log("Success:", data);
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={todo}
          onChange={handleInputChange}
          style={{
            padding: "10px",
            fontSize: "16px",
            borderRadius: "5px",
            border: "1px solid #ccc",
            marginRight: "10px",
          }}
        />
        <button
          type="submit"
          style={{
            padding: "10px 20px",
            fontSize: "16px",
            borderRadius: "5px",
            backgroundColor: "#4CAF50",
            color: "white",
            border: "none",
            cursor: "pointer",
          }}
        >
          Add Todo
        </button>
      </form>
    </div>
  );
};

export default AddTodo;
