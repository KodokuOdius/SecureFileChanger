import React, { useContext, useEffect, useState } from "react";
import { TokenContext } from "../context";
import { useAPI } from "../hooks/useAPI";
import Loader from "../components/Loader";
import UserList from "../components/User/UserList";
import AdminService from "../api/AdminService";
import { useNavigate } from "react-router-dom";


const AdminPanel = () => {
    const { token } = useContext(TokenContext)
    const navigate = useNavigate()

    const [users, setUsers] = useState([])
    const [getUserList, isLoading, errorList] = useAPI(async () => {
        const users = await AdminService.userList(token)

        if (users === null) {
            return
        }

        setUsers(users)
    })

    useEffect(() => {
        getUserList()

        if (errorList.responce && errorList.responce.status === 403) {
            return () => navigate("/")
        }
    }, []) // eslint-disable-line react-hooks/exhaustive-deps

    const [searchEmailChange, setSearchEmailChange] = useState("")
    const [searchEmail, setSearchEmail] = useState("")

    const onSearch = (e) => {
        e.preventDefault()
        setSearchEmail(searchEmailChange)
    }

    return (
        <div className="users__page">
            {!isLoading &&
                <div className="list__search">
                    <h2>Поисковая строка</h2>
                    <input
                        type="text"
                        value={searchEmailChange}
                        placeholder="Введите email Сотрудника"
                        onChange={e => setSearchEmailChange(e.target.value)}
                    />
                    <button onClick={onSearch} >Поиск</button>
                </div>
            }
            {isLoading &&
                <Loader msg="Загрузка информации о Сотрудниках" />
            }
            {users === null || users.length === 0
                ? <></>
                : <UserList users={users} likeEmail={searchEmail} />
            }
        </div>
    )
}
export default AdminPanel