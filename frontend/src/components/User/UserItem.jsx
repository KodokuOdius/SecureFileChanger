import React, { useContext, useState } from "react";
import { useAPI } from "../../hooks/useAPI";
import AdminService from "../../api/AdminService";
import { TokenContext } from "../../context";


// email: "123@qwe.qwe"
// id: 16
// is_admin: false
// is_approved: false
// user_name: null
// user_surname: null
const UserItem = ({ idx, user }) => {
    const { token } = useContext(TokenContext)
    const [isApproved, setApprove] = useState(user.is_approved)

    const [toggleApprove, , errorAPI] = useAPI(async () => {
        const resp = await AdminService.toggleApprove(token, user.id)

        if (resp !== null) {
            console.log(resp)
            setApprove(!isApproved)
        }
    })

    const onToggleApprove = () => {
        toggleApprove()
        if (errorAPI.responce) {
            console.log(errorAPI.message)
        }
    }


    return (
        <div className="user__item">
            <p>{idx}# Email: {user.email}</p>
            <p className="user__approved">
                <label className="user__switch">
                    <input
                        type="checkbox"
                        checked={isApproved}
                        onChange={onToggleApprove}
                    />
                    <span className="user__slider"></span>
                </label>
            </p>
        </div>
    )
}
export default UserItem