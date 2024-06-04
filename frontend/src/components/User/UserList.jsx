import React from "react";
import UserItem from "./UserItem";


const UserList = ({ users, likeEmail }) => {

    return (
        <div className="user__list">
            {users.map((user, idx) => {
                if (likeEmail === "") {
                    return <UserItem idx={idx} key={user.id} user={user} />
                }

                if (user.email.indexOf(likeEmail) >= 0) {
                    return <UserItem idx={idx} key={user.id} user={user} />
                }
                return <></>
            })}
        </div>
    )
}
export default UserList