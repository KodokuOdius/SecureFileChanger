import React, { useContext, useRef, useState } from "react";
import UserService from "../../api/UserService";
import { TokenContext } from "../../context";
import { useAPI } from "../../hooks/useAPI";


const ChangePass = ({ setIsPassShow }) => {
    const { token } = useContext(TokenContext)

    const mainDiv = useRef()

    const defaultPasswords = { old_password: "", new_pasword: "" }
    const [passwords, setPasswords] = useState(defaultPasswords)

    const [changePass, , errorChange] = useAPI(async () => {
        const res = await UserService.newPassword(token, passwords.old_password, passwords.new_pasword)
        console.log(res)
    })

    const cancel = (e) => {
        e.preventDefault()
        setPasswords(defaultPasswords)
        setIsPassShow(false)
    }

    const change = async (e) => {
        e.preventDefault()
        if (passwords.new_pasword === passwords.old_password) {
            console.log("same passwords")
            return
        }

        changePass()
        if (errorChange !== null) {
            console.log("Неправильный старый пароль или слишком слабый новый пароль")
        } else {
            console.log("Пароль изменён")
            setIsPassShow(false)
        }
    }

    const onModalClick = (e) => {
        if (e.target === mainDiv.current) {
            cancel(e)
        }
    }

    return (
        <div className="change__password" ref={mainDiv} onClick={onModalClick} >
            <form method="post" className="change__pass">
                <div className="pass__inp">
                    <div className="inp__item">
                        <p>Введите старый пароль</p>
                        <input
                            type="password"
                            placeholder="Старый пароль"
                            value={passwords.old_password}
                            minLength={1}
                            maxLength={100}
                            onChange={e => setPasswords({ ...passwords, old_password: e.target.value })}
                        />
                    </div>
                    <div className="inp__item">
                        <p>Введите новый пароль</p>
                        <input
                            type="password"
                            placeholder="Новый пароль"
                            minLength={1}
                            maxLength={100}
                            value={passwords.new_pasword}
                            onChange={e => setPasswords({ ...passwords, new_pasword: e.target.value })}
                        />
                    </div>
                </div>
                <div className="form__btns">
                    <button onClick={change}>Изменить</button>
                    <button onClick={cancel}>Отменить</button>
                </div>
            </form>
        </div>
    )
}
export default ChangePass