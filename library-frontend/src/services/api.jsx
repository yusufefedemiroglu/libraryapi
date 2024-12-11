import API from "./apiConfig";

export const login = (username, password) =>
  API.post("/login", { username, password });

export const fetchBooks = (userID,token) =>  {
  token = JSON.parse(localStorage.getItem("userdata")).token; // Kullanıcı token'ını localStorage'dan alın (veya uygun bir yerden).
  userID= JSON.parse(localStorage.getItem("userdata")).id; 
  console.log(userID)// Kullanıcı token'ını localStorage'dan alın (veya uygun bir yerden).
    return API.get("/books", {
        headers: {
            Authorization: `Bearer ${token}`,
            "x-user-id": userID, // Authorization başlığına token ekleniyor.
        },
    });

}


export const addBook = async (bookData) => {
    const response =  API.post("/books", bookData);
    return response.data;
};
export const createUser = async (username,password) => {
  const res = API.post("/create",{username,password})
  return res.data
};

export const fetchUser = () => API.get("/user");
