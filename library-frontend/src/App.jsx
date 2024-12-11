import { useState } from "react";
import { AddBook } from "./components/AddBook";
import Login from "./components/Login";
import { BookList } from "./components/BookList";
export function App() {

    const [isAuthenticated,setAuth] = useState(!!localStorage.getItem("userdata"));
    const [reload,setReload] = useState(false)


  return (
    <div>
     {!isAuthenticated ? (
       <Login setAuth={setAuth}/>
      ) : (
        <>
      <BookList reload={reload}/>
      <AddBook setReload={setReload}/>
      </>
     )}
    </div>
  )
}


