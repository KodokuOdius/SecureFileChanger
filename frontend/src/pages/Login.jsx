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
            <h2 className="login__title" >Форма авторизации</h2>
            <div className="login__inp">
                <form method="post">
                    <div className="inp__item">
                        <p className="item__title">Введите email</p>
                        <input
                            required
                            value={loginForm.email}
                            id="email"
                            type="email"
                            minLength={1}
                            maxLength={100}
                            placeholder="Email"
                            onChange={e => setLoginForm({ ...loginForm, email: e.target.value })}
                        />
                    </div>
                    <div className="inp__item">
                        <p className="item__title">Введите пароль</p>
                        <input
                            required
                            value={loginForm.password}
                            id="password"
                            type="password"
                            minLength={8}
                            maxLength={100}
                            placeholder="Пароль"
                            onChange={e => setLoginForm({ ...loginForm, password: e.target.value })}
                        />
                    </div>
                </form>
                <div className="error__msg">
                    <p>{errorMsg}</p>
                </div>
            </div>
            <div className="login__btns">
                <button onClick={autherization}>Войти</button>
                <p>
                    <Link to="/register">Зарагистрироваться</Link>
                </p>
            </div>
        </div>
    )
}

export default Login