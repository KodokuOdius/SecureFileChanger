import axios from "axios"
import { APIServer } from "../App"


// user: {
//     info: "/user/info",
//     delete: "/user/delete/",
//     update: "/user/update/",
//     newPassword: "/user/new-password",
//     limit: "/user/limit",
// },
export default class UserService {
    static async getInfo(token) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token
            }
        }

        const url = APIServer.serverHost + APIServer.user.info
        const resp = await axios.get(url, AuthHeader)

        return resp.data
    }

    static async deleteUser(token) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token
            }
        }

        const url = APIServer.serverHost + APIServer.user.delete
        const resp = await axios.delete(url, AuthHeader)

        return resp
    }

    static async updateUser(token, name, surname) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token,
            }
        }

        const url = APIServer.serverHost + APIServer.user.update
        const resp = await axios.put(
            url,
            { "user_name": name, "user_surname": surname },
            AuthHeader
        )

        return resp.data
    }

    static async usedBytes(token) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token
            }
        }

        const url = APIServer.serverHost + APIServer.user.limit
        const resp = await axios.get(url, AuthHeader)

        return resp.data.used_bytes
    }

    static async newPassword(token, oldPass, newPass) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token
            }
        }

        const url = APIServer.serverHost + APIServer.user.newPassword
        const resp = await axios.post(
            url,
            { "old_password": oldPass, "new_password": newPass },
            AuthHeader
        )

        return resp.data
    }

}