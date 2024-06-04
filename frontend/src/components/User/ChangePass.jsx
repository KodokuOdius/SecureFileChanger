import React, { useContext, useState } from "react";
import UserService from "../../api/UserService";
import { TokenContext } from "../../context";
import { useAPI } from "../../hooks/useAPI";


const ChangePass = ({ setIsPassShow }) => {
    const { token } = useContext(TokenContext)

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

    return (
        <form method="post" className="change__pass">
            <input
                type="password"
                placeholder="Старый пароль"
                value={passwords.old_password}
                onChange={e => setPasswords({ ...passwords, old_password: e.target.value })}
            />

            <input
                type="password"
                placeholder="Новый пароль"
                value={passwords.new_pasword}
                onChange={e => setPasswords({ ...passwords, new_pasword: e.target.value })}
            />

            <div className="form__btns">
                <button onClick={change}>Изменить</button>
                <button onClick={cancel}>Отменить</button>
            </div>
        </form>
    )
}
export default ChangePass