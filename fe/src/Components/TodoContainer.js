import React from "react";
import TodoCard from "./TodoCard";

function TodoContainer(props) {
  var todos = props.todos;
  return (
    <>
      <div className="container mx-auto mt-5 px-4">
        <h1 className="text-center text-4xl">Todo Manager</h1>
        {todos.length === 0 ? (
          <h2 className="text-center text-2xl">No todos found</h2>
        ) : (
          todos.map((todo) => {
            return (
              <TodoCard
                completed={todo.completed}
                createdAt={todo.createdAt}
                deletedAt={todo.deletedAt}
                description={todo.description}
                id={todo.id}
                title={todo.title}
                updatedAt={todo.updatedAt}
                key={todo.id}
              />
            );
          })
        )}
      </div>
    </>
  );
}

export default TodoContainer;
