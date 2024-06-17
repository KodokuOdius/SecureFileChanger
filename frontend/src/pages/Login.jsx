import React, { useContext, useEffect, useState } from "react";
import { TokenContext } from "../context";
import { Navigate, Link, useNavigate } from "react-router-dom";
import AuthService from "../api/AuthService";
import { useAPI } from "../hooks/useAPI";
import { APIServer } from "../App";


const Login = ({ setIsLoginShow }) => {
    const defaultLogin = { email: "", password: "" }
    const [loginForm, setLoginForm] = useState(defaultLogin)
    const [errorMsg, setErrorMsg] = useState("")

    const navigate = useNavigate()

    const { token, setToken } = useContext(TokenContext)

    const [logInApi, , errorLogIn] = useAPI(async () => {
        const newToken = await AuthService.logIn(loginForm)
        setToken(newToken)
    })

    useEffect(() => {
        const t = localStorage.getItem(APIServer.tokenName)
        if (t !== "") {
            return () => { navigate("/") }
        }
    }, []) // eslint-disable-line react-hooks/exhaustive-deps

    const autherization = async (e) => {
        e.preventDefault()

        if (defaultLogin === loginForm) {
            return
        }

        logInApi()
        if (errorLogIn !== null) {
            if (!errorLogIn.response) {
                setErrorMsg("Произошла ошибка")
                return
            }
            if (errorLogIn.response.data.message === "user not approved") {
                setErrorMsg("Пользователь не подтрверждён")
            } else {
                setErrorMsg("Неправильный email или пароль")
            }
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
            <div className="register__inp">
                {errorMsg !== ""
                    ? <div className="error_msg">
                        <p>{errorMsg}</p>
                    </div>
                    : <></>
                }
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
            </div>
            <div className="login__btns">
                <button onClick={autherization}>Войти</button>
                <Link to="/register">Зарагестрироваться</Link>
            </div>
        </div>
    )
}

export default Login