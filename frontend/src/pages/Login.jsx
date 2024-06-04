import React, { useContext, useState } from "react";
import { TokenContext } from "../context";
import { Navigate, Link } from "react-router-dom";
import AuthService from "../api/AuthService";
import { useAPI } from "../hooks/useAPI";


const Login = ({ setIsLoginShow }) => {
    const defaultLogin = { email: "", password: "" }
    const [loginForm, setLoginForm] = useState(defaultLogin)

    const { token, setToken } = useContext(TokenContext)

    const [logInApi, , errorLogIn] = useAPI(async () => {
        const newToken = await AuthService.logIn(loginForm)
        setToken(newToken)
    })

    const autherization = async (e) => {
        e.preventDefault()

        logInApi()
        if (errorLogIn !== null) {
            console.log("Неправильный email или пароль")
            return
        }
        setIsLoginShow(false)
    }

    if (token !== "") {
        return <Navigate to="/" />
    }

    return (
        <div className="login__form">
            <h1>Форма авторизации</h1>
            <form method="post">
                <input
                    required
                    value={loginForm.email}
                    id="email"
                    type="email"
                    placeholder="Введите свой email"
                    onChange={e => setLoginForm({ ...loginForm, email: e.target.value })}
                />
                <input
                    required
                    value={loginForm.password}
                    id="password"
                    type="password"
                    minLength={8}
                    placeholder="Введите свой пароль"
                    onChange={e => setLoginForm({ ...loginForm, password: e.target.value })}
                />
            </form>
            <button onClick={autherization}>Войти</button>
            <br />
            <Link to="/register">Зарагестрироваться</Link>
        </div>
    )
}

export default Login