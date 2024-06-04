import React, { useContext, useEffect, useState } from "react";
import { TokenContext } from "../context";
import { useAPI } from "../hooks/useAPI";
import UserService from "../api/UserService";
import { APIServer, humanSizeBytes } from "../App";
import Loader from "../components/Loader";
import ChangePass from "../components/User/ChangePass";
import { useNavigate } from "react-router-dom";


const Profile = () => {
    const { token, setToken } = useContext(TokenContext)
    const [isUpdateShow, setIsUpdateShow] = useState(false)
    const [isPassShow, setIsPassShow] = useState(false)
    const navigate = useNavigate()

    const defaultUserInfo = {
        email: "", name: "", surname: "",
        is_admin: false, is_approved: false,
        used_bytes: 0
    }

    const [userInfo, setUserInfo] = useState(defaultUserInfo)

    const [getUserInfo, isLoading, errorApi] = useAPI(async () => {
        const user = await UserService.getInfo(token)
        if (user === null) {
            return
        }

        if (user.name === null) {
            user.name = ""
        }

        if (user.surname === null) {
            user.surname = ""
        }

        user.old_name = user.name
        user.old_surname = user.surname

        setUserInfo(user)
    })

    const [deleteUser, , errorDelete] = useAPI(async () => {
        const res = await UserService.deleteUser(token)
        console.log(res)
    })

    useEffect(() => {
        getUserInfo()
        console.log(errorApi)
    }, []) // eslint-disable-line react-hooks/exhaustive-deps

    const onDeleteUser = (e) => {
        deleteUser()
        if (errorDelete !== null) {
            setToken("")
            localStorage.setItem(APIServer.tokenName, "")
        }

        return navigate("/login")
    }


    const toggleShowUpdatePanel = () => {
        if (isUpdateShow) {
            setUserInfo({ ...userInfo, name: userInfo.old_name, surname: userInfo.old_surname })
        }
        setIsUpdateShow(!isUpdateShow)
    }

    const onSaveUpdate = async (e) => {
        if (userInfo.name.trim().length === 0 || userInfo.name.trim().length === 0) {
            setUserInfo({ ...userInfo, name: userInfo.old_name, surname: userInfo.old_surname })
            setIsUpdateShow(false)
            return
        }

        const res = await UserService.updateUser(token, userInfo.name, userInfo.surname)
        if (res === null) {
            setUserInfo({ ...userInfo, name: userInfo.old_name, surname: userInfo.old_surname })
        }

        setUserInfo({ ...userInfo, old_name: userInfo.name, old_surname: userInfo.surname })
        setIsUpdateShow(false)
    }

    const onChangePassword = () => setIsPassShow(true)
    const onAdminPanel = () => navigate("/admin-panel")

    return (
        <div className="user__profile">
            <h2>Профиль пользователя</h2>
            <div className="user__info">
                {isLoading
                    ? <Loader msg="Загрузка ифнормации о профиле" />
                    : <div className="info__detail">
                        {userInfo.is_admin &&
                            <p className="info__value">Администратор</p>
                        }
                        <p className="info__value">
                            Email: {userInfo.email}
                        </p>
                        <p className="info__value">
                            Имя:
                            <input
                                type="text"
                                value={userInfo.name}
                                disabled={!isUpdateShow}
                                onChange={e => setUserInfo({ ...userInfo, name: e.target.value })}
                            />
                        </p>
                        <p className="info__value">
                            Фамилия:
                            <input
                                type="text"
                                value={userInfo.surname}
                                disabled={!isUpdateShow}
                                onChange={e => { setUserInfo({ ...userInfo, surname: e.target.value }) }}
                            />
                        </p>
                        <p className="info__value">
                            Учётная запись {!userInfo.is_approved && "не"} подтверждена
                        </p>
                        <p className="info__value">
                            Использовано памяти: {humanSizeBytes(userInfo.used_bytes)}
                        </p>
                    </div>
                }
            </div>

            <h3>Кнопки для провиля</h3>
            <div className="profile__btns">
                <button onClick={toggleShowUpdatePanel}>
                    {isUpdateShow ? "Отменить" : "Редактировать"}
                </button>
                {isUpdateShow
                    ? <button onClick={onSaveUpdate}>Сохранить</button>
                    : <>
                        <button onClick={onChangePassword}>Сменить пароль</button>
                        {isPassShow &&
                            <ChangePass setIsPassShow={setIsPassShow} />
                        }
                        {!userInfo.is_admin
                            ? <button onClick={onDeleteUser}>Удалить</button>
                            : <button onClick={onAdminPanel}>Список сотрудников</button>
                        }
                    </>
                }
            </div>
        </div>
    )
}
export default Profile