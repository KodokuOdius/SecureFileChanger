import axios from "axios"
import { APIServer } from "../App"


// file: {
//     list: "/file/list",
//     upload: "/file/upload",
//     downloadMany: "/file/download-many",
//     download: "/file/download/",
//     delete: "/file/delete/",
//     update: "/file/update/",
// },
export default class FileService {

    static async getList(token) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token
            }
        }

        const url = APIServer.serverHost + APIServer.file.list
        const resp = await axios.get(url, AuthHeader)

        const files = resp.data.data
        if (files === null) {
            return []
        }

        return files
    }

    static async deleteFile(token, fileId) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token
            }
        }

        const url = APIServer.serverHost + APIServer.file.delete
        const resp = await axios.delete(
            url + fileId, AuthHeader
        )

        return resp.data
    }

    static async updateName(token, fileId, newName) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token
            }
        }

        const url = APIServer.serverHost + APIServer.file.update
        const resp = await axios.put(
            url + fileId,
            { "file_name": newName },
            AuthHeader
        )

        return resp.data
    }

    static async uploadFile(token, file) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token
            }
        }
        const url = APIServer.serverHost + APIServer.file.upload
        const resp = await axios.post(
            url, file, AuthHeader
        )

        return resp.data
    }

    static async downloadFile(token, fileId) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token
            },
            responseType: "blob"
        }
        const url = APIServer.serverHost + APIServer.file.download
        const resp = await axios.get(
            url + fileId, AuthHeader
        )

        return resp.data
    }

    static async downloadFiles(token, fileIds) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token
            },
            responseType: "blob"
        }
        const url = APIServer.serverHost + APIServer.file.downloadMany
        const resp = await axios.post(
            url, { "file_id": fileIds }, AuthHeader
        )

        return resp.data
    }

}