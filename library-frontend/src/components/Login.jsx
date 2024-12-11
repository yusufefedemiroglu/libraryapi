import  {useState} from "react";
import {login} from "../services/api"
import {createUser} from "../services/api"
import "../../styles/login.scss"

// eslint-disable-next-line react/prop-types
const Login = ({ setAuth }) => {
    const [username,setUsername] = useState("");
    const [password,setPassword] = useState("");
    const switchers = [...document.querySelectorAll('.switcher')]

    switchers.forEach(item => {
	item.addEventListener('click', function() {
		switchers.forEach(item => item.parentElement.classList.remove('is-active'))
		this.parentElement.classList.add('is-active')
	})
    })


    const handleLogin = async (e) => {
        e.preventDefault();
            try {
                const {data} = await login(username,password);
                const userdata = {
                    token : data.token,
                    id:data.id
                }
                localStorage.setItem("userdata",JSON.stringify(userdata));
                setAuth(true);
                alert("success");
            } catch (error) {
                alert(error)
            }
        
    }
    const handleCreateUser = async (e) => {
        e.preventDefault();
        try {
            // API'ye kullanıcı oluşturma isteği gönderiliyor
            const data = await createUser(username, password);
            alert("User created successfully!");
            console.log("New user data:", data); // API'den dönen verileri loglayın
        } catch (error) {
            // Hata durumunda kullanıcıya bildirim gösteriliyor
            console.error("Error creating user:", error);
            alert("Failed to create user. Please try again.");
        }
    };
    

    return (
     <section className="forms-section">
  <h1 className="section-title">Giriş Ekranı</h1>
  <div className="forms">
    {/* Login Form */}
    <div className="form-wrapper is-active">
      <button type="button" className="switcher switcher-login">
        Giriş Yap
        <span className="underline"></span>
      </button>
      <form className="form form-login" onSubmit={handleLogin}>
        <fieldset>
          <legend>Please, enter your username and password for login.</legend>
          <div className="input-block">
            <label htmlFor="login-username">Username</label>
            <input
              id="login-username"
              type="text"
              placeholder="Username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
          </div>
          <div className="input-block">
            <label htmlFor="login-password">Password</label>
            <input
              id="login-password"
              type="password"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
        </fieldset>
        <button type="submit" className="btn-login">Login</button>
      </form>
    </div>
    {/* Sign Up Form */}
    <div className="form-wrapper">
      <button type="button" className="switcher switcher-signup">
        Kayıt Ol
        <span className="underline"></span>
      </button>
      <form className="form form-signup" onSubmit={handleCreateUser}>
        <fieldset>
          <legend>
            Please, enter your email, password and password confirmation for
            sign up.
          </legend>
          <div className="input-block">
            <label htmlFor="signup-email">Username</label>
            <input
              id="Username"
              type="username"
              placeholder="Username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
          </div>
          <div className="input-block">
            <label htmlFor="signup-password">Password</label>
            <input
              id="signup-password"
              type="password"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
        </fieldset>
        <button type="submit" className="btn-signup">Kayıt ol</button>
      </form>
    </div>
  </div>
</section>

    
    );
};
export default Login;
