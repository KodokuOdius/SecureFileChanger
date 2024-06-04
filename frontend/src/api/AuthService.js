import axios from "axios";
import { APIServer } from "../App";

// auth: {
//     register: "/auth/register",
//     login: "/auth/login",
// },
export default class AuthService {
    // { email: "", password: "" }
    static async logIn(loginForm) {
        const url = APIServer.serverHost + APIServer.auth.login
        const resp = await axios.post(
            url, loginForm
        )
        const token = resp.data.token
        localStorage.setItem(APIServer.tokenName, token)

        return token
    }

    // { email: "", password: "" }
    static async register(registerForm) {
        const url = APIServer.serverHost + APIServer.auth.register
        const resp = await axios.post(
            url, registerForm
        )

        const newUserId = resp.data

        return newUserId
    }
}