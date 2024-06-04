import axios from "axios"
import { APIServer } from "../App"


// folder: {
//     create: "/folder/create",
//     list: "/folder/all",
//     getFiles: "/folder/",
//     update: "/folder/update/",
//     delete: "/folder/",
// },
export default class FolderService {

    static async getAll(token) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token
            }
        }
        const resp = await axios.get(
            APIServer.serverHost + APIServer.folder.list,
            AuthHeader
        )
        const folders = resp.data.data
        if (folders === null) {
            return []
        }

        return folders
    }

    static async deleteFolder(token, folderId) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token
            }
        }
        const resp = await axios.delete(
            APIServer.serverHost + APIServer.folder.delete + folderId,
            AuthHeader
        )
        return resp.data
    }


    static async createFolder(token, folderName) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token
            }
        }
        const url = APIServer.serverHost + APIServer.folder.create
        const resp = await axios.post(
            url,
            { "folder_name": folderName },
            AuthHeader
        )
        return resp.data
    }

    static async getFiles(token, folderId) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token
            }
        }
        const url = APIServer.serverHost + APIServer.folder.getFiles
        const resp = await axios.get(
            url + folderId, AuthHeader
        )

        const files = resp.data.data
        if (files === null) {
            return []
        }

        return files
    }

    static async updateFolder(token, folderId, folderName) {
        const AuthHeader = {
            headers: {
                "Authorization": "Bearer " + token
            }
        }

        const url = APIServer.serverHost + APIServer.folder.update
        const resp = await axios.put(
            url + folderId,
            { "folder_name": folderName },
            AuthHeader
        )

        return resp.data
    }
}