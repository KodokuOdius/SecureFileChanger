import React, { useContext, useState } from "react";
import { TokenContext } from "../context";
import { Navigate, useNavigate } from "react-router-dom";
import AuthService from "../api/AuthService";
import { useAPI } from "../hooks/useAPI";


const Register = () => {
    const defaultReg = { email: "", password: "" }
    const [regData, setRegData] = useState(defaultReg)

    const navigate = useNavigate()

    const [register, , errorRegister] = useAPI(async () => {
        const resp = await AuthService.register(regData)
        if (resp !== null) {
            console.log("Ожидайте подтверждение доступа")
        }
    })

    const { token } = useContext(TokenContext)
    if (token !== "") {
        return <Navigate to="/" />
    }

    const registration = (e) => {
        e.preventDefault()

        if (regData === defaultReg) {
            return
        }

        register()
        if (errorRegister !== null) {
            return navigate("/login")
        }
    }

    return (
        <div className="register__form">
            <h1>Форма регистрации</h1>
            <form method="post">
                <input
                    required
                    value={regData.email}
                    id="email"
                    type="email"
                    placeholder="Введите свой email"
                    onChange={e => setRegData({ ...regData, email: e.target.value })}
                />
                <input
                    required
                    value={regData.password}
                    id="password"
                    type="password"
                    minLength={8}
                    placeholder="Введите свой пароль"
                    onChange={e => setRegData({ ...regData, password: e.target.value })}
                />
            </form>
            <button onClick={registration}>Зарагестрироваться</button>
        </div>
    )
}

export default Register