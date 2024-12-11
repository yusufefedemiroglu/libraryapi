import { useEffect, useState } from "react";
import { fetchBooks } from "../services/api";
import "../../styles/bl.scss"

// eslint-disable-next-line react/prop-types
export const BookList = ({reload}) => {
    const [books, setBooks] = useState([]);
    const [errorMessage, setErrorMessage] = useState('');
    const [isLoading, setIsLoading] = useState(true);
    const userID = JSON.parse(localStorage.getItem("userdata")).id;// Örnek: Kullanıcının ID'si localStorage'dan alınıyor
    const token = JSON.parse(localStorage.getItem("userdata")).token;    // Token da aynı şekilde alınıyor
     async () => {
        try {
            const { data } = await fetchBooks({userID,token});
            setBooks(data);
            console.log(books)
        } catch (error) {
            alert(error);
        }
    };

    useEffect(() => {
        fetchBooks(userID, token)
        .then(response => {
          setIsLoading(false);
          if (response.data === null) {
            setErrorMessage('Henüz kitabınız yok. Kitap eklemek için aşağıdaki formu kullanabilirsiniz.');
          } else {
            setBooks(response.data);
          }
        })
        .catch(error => {
            console.log(error)
          setIsLoading(false);
          setErrorMessage('Kitaplar alınırken bir hata oluştu.');
        }); 
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [reload]);

    return (
        <div>
        {isLoading ? (
          <p>Yükleniyor...</p> // Yükleme sırasında gösterilecek mesaj
        ) : (
          <>
            {errorMessage ? (
              <p>{errorMessage}</p> // Kitap yoksa gösterilecek mesaj
            ) : (
              <ul>
        {books.map(book => (
        <li key={book._id}>
        {book.title} - {book.author} {/* Kitap başlığı ve yazarı */}
        </li>
        ))}
              </ul>
            )}
          </>
        )}
      </div>
    );
};
