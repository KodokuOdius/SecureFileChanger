import axios from "axios"
import { APIServer } from "../App"


// admin: {
//     userList: "/admin/user-list",
//     userListSearch: "/admin/user-list/search",
//     toggleApprove: "/admin/toggle-approve/",
// },
export default class AdminService {

    static async userList(token) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token
            }
        }
        const url = APIServer.serverHost + APIServer.admin.userList
        const resp = await axios.get(url, AuthHeader)

        const users = resp.data.data
        if (users === null) {
            return []
        }
        return users
    }

    static async toggleApprove(token, userId) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token
            }
        }

        const url = APIServer.serverHost + APIServer.admin.toggleApprove
        const resp = await axios.put(
            url + userId, {},
            AuthHeader
        )

        return resp.data
    }
}
