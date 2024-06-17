import React, { useContext, useState } from "react";
import { TokenContext } from "../context";
import { Navigate, Link } from "react-router-dom";
import AuthService from "../api/AuthService";
import { useAPI } from "../hooks/useAPI";


const Register = () => {
    const defaultReg = { email: "", password: "" }
    const [regData, setRegData] = useState(defaultReg)
    const [errorMsg, setErrorMsg] = useState("")
    const [isSuccsess, setIsSuccsess] = useState(false)

    // const navigate = useNavigate()

    const [register, , errorRegister] = useAPI(async () => {
        const resp = await AuthService.register(regData)
        if (resp !== null) {
            setIsSuccsess(true)
            setErrorMsg("")
        }
    })

    const { token } = useContext(TokenContext)
    if (token !== "") {
        return <Navigate to="/" />
    }

    const isValidEmail = (email) => {
        return /\S+@\S+\.\S+/.test(email)
    }

    const registration = (e) => {
        e.preventDefault()

        if (regData === defaultReg) {
            return
        }

        if (regData.password.length < 8) {
            setErrorMsg("Пароль долзжен быть более 8 символов")
            return
        }

        if (!isValidEmail(regData.email)) {
            setErrorMsg("Некорректный email адрес")
            return
        }

        register()
        if (errorRegister !== null) {
            setErrorMsg("Произошла ошибка во время регистрации")
            return
        }
        setErrorMsg("")
        setIsSuccsess(true)
    }

    return (
        <div className="register__form">
            <h1>Форма регистрации</h1>
            <div className="register__inp">
                {errorMsg !== "" &&
                    <div className="error__msg">
                        <p>{errorMsg}</p>
                    </div>
                }
                {isSuccsess &&
                    <div className="succsess__msg">
                        <p>Ожидайте подтверждения Администратора</p>
                    </div>
                }
                <form method="post">
                    <input
                        required
                        value={regData.email}
                        id="email"
                        type="email"
                        placeholder="Введите свой email"
                        onChange={e => setRegData({ ...regData, email: e.target.value })}
                        disabled={isSuccsess}
                    />
                    <input
                        required
                        value={regData.password}
                        id="password"
                        type="password"
                        minLength={8}
                        placeholder="Введите свой пароль"
                        onChange={e => setRegData({ ...regData, password: e.target.value })}
                        disabled={isSuccsess}
                    />
                </form>
            </div>
            <div className="register__form">
                {!isSuccsess &&
                    <button onClick={registration}>Зарагестрироваться</button>
                }
                <Link to="/login">Войти</Link>
            </div>
        </div>
    )
}

export default Register