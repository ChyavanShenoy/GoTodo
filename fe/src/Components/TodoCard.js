import React, { useState } from "react";
import { useEffect } from "react";
import { useRef } from "react";

function TodoCard({
  id,
  title,
  description,
  completed,
  createdAt,
  updatedAt,
  deletedAt,
}) {
  const [expanded, setExpanded] = useState(false);
  const [editing, setEditing] = useState(false);
  const [newTitle, setNewTitle] = useState(title);
  const [newDescription, setNewDescription] = useState(description);
  const [newCompleted, setNewCompleted] = useState(completed);
  const editFormRef = useRef(null); // Ref for the edit form

  function onClickExpand(e) {
    e.stopPropagation(); // Prevent triggering when clicking on child elements
    setExpanded(!expanded);
  }

  function onClickEdit(e) {
    e.stopPropagation(); // Prevent triggering onClickExpand
    setEditing(!editing);
  }

  function handleClickOutside(event) {
    if (editFormRef.current && !editFormRef.current.contains(event.target)) {
      setEditing(false);
    }
  }

  function handleClickCancel(event) {
    setEditing(false);
  }

  useEffect(() => {
    // Add event listener
    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      // Remove event listener
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  function updateTodo() {
    fetch(`http://localhost:8080/api/v1/updateTodo`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        id: id,
        title: newTitle,
        description: newDescription,
        completed: newCompleted,
      }),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log("Todo updated successfully");
        setEditing(false);
      })
      .catch((error) => {
        console.error("Error updating todo:", error);
      });
  }

  // Function to format the date for display (you can adjust this as needed)
  const formatDate = (date) =>
    date && new Date(date).toLocaleDateString("en-US");

  return (
    <li
      className="flex flex-col justify-between bg-gray-700 text-white font-bold py-2 px-4 rounded my-2 cursor-pointer"
      onClick={onClickExpand}
    >
      <div className="flex justify-between items-center">
        <input
          type="checkbox"
          className="form-checkbox h-5 w-5"
          checked={completed}
          readOnly
        />
        <span className="text-xl">{title}</span>
        <span className="">
          <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mx-5 w-20"
            onClick={onClickEdit}
          >
            Edit
          </button>
          <button className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded w-20">
            Delete
          </button>
        </span>
      </div>
      {expanded && (
        <div className="mt-4">
          <p className="text-lg font-normal mb-2">Description: {description}</p>
          <p>Created At: {formatDate(createdAt)}</p>
          <p>Updated At: {formatDate(updatedAt)}</p>
        </div>
      )}
      {editing && (
        <div
          className="absolute top-0 left-0 right-0 bottom-0 bg-white bg-opacity-75 flex items-center justify-center z-10"
          onClick={onClickEdit}
        >
          {/* Replace this with your actual edit form */}
          <form
            className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4"
            ref={editFormRef}
            onClick={(e) => e.stopPropagation()} // Prevent triggering onClickEdit
          >
            <div className="mb-4">
              <label
                className="block text-gray-700 text-sm font-bold mb-2"
                for="title"
              >
                Title
              </label>
              <input
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                id="title"
                type="text"
                placeholder="Title"
                value={title}
              />
            </div>
            <div className="mb-6">
              <label
                className="block text-gray-700 text-sm font-bold mb-2"
                for="description"
              >
                Description
              </label>
              <textarea
                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                id="description"
                type="text"
                placeholder="Description"
                value={description}
              />
            </div>
            <div className="flex items-center justify-between my-4">
              <label
                className="block text-gray-700 text-sm font-bold mb-2"
                for="completed"
              >
                Completed
              </label>
              <input
                className="form-checkbox h-5 w-5"
                id="completed"
                type="checkbox"
                checked={completed}
                onClick={(e) => setNewCompleted(e.target.checked)}
              />
            </div>
            <div className="flex items-center justify-between">
              <button
                className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                type="button"
                onClick={updateTodo}
              >
                Save
              </button>
              <button
                className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                type="button"
                onClick={handleClickCancel}
              >
                Cancel
              </button>
            </div>
          </form>
        </div>
      )}
    </li>
  );
}

export default TodoCard;
