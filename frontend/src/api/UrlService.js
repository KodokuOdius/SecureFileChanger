import axios from "axios"
import { APIServer } from "../App"

// url: {
//     create: "/url/create",
//     download: "/url-get/download/",
//     files: "/url-get/files/",
// },
export default class UrlService {
    static async createUrl(token, filesIds, hourLive) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token
            }
        }
        const url = APIServer.serverHost + APIServer.url.create
        const resp = await axios.post(
            url,
            { "file_id": filesIds, "hour_live": hourLive },
            AuthHeader
        )

        if (resp.data === null) {
            return ""
        }

        const resultUrl = window.location.origin + "/d/" + resp.data.url_hex
        return resultUrl
    }

    static async downloadFiles(uuid) {
        const url = APIServer.serverHost + APIServer.url.download
        const resp = await axios.get(
            url + uuid, { responseType: "blob" }
        )

        return resp.data
    }

    static async getFiles(uuid) {
        const url = APIServer.serverHost + APIServer.url.files
        const resp = await axios.get(url + uuid)

        const files = resp.data.data
        if (files === null) {
            return []
        }

        return files
    }
}