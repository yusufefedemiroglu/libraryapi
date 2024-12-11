import { useState } from "react";
import { addBook } from "../services/api";

// eslint-disable-next-line react/prop-types
export const AddBook = ({setReload}) => {
    const [title, setTitle] = useState("");
    const [author, setAuthor] = useState("");

    const handleAddBook = async (e) => {
        e.preventDefault();
        try {
         const userid = JSON.parse(localStorage.getItem("userdata")).id;
           const res = await addBook({ title, author,userid });
           console.log(res)
           alert("Successfully added book");
            setTitle(""); 
            setAuthor("");
            setReload((prev) => !prev)
        } catch (error) {
            alert("Failed to add book: " + error.message);
        }
    };

    return (
        <form onSubmit={handleAddBook} className="form__group field">
            <h2>Kitap Oluştur</h2>
            <input
                type="text"
                placeholder="Title"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                className="form__field"
                required
            />
            <input
                type="text"
                placeholder="Author"
                value={author}
                onChange={(e) => setAuthor(e.target.value)}
                className="form__field"
                required
            />
            <button type="submit" className="button-2">Add</button>
        </form>
    );
};
