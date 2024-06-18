import React, { useState } from 'react';
import './styles/App.css'
import Rounter from './components/Router/Router';
import { TokenContext } from './context';
import { library } from '@fortawesome/fontawesome-svg-core';
import { faFolder, faFileExcel, faFilePdf, faFilePowerpoint, faFileWord, faFileText, faFile } from "@fortawesome/free-solid-svg-icons";

library.add(faFolder, faFileExcel, faFilePdf, faFilePowerpoint, faFileWord, faFileText, faFile)

export const FilesIcons = {
    "pdf": faFilePdf,
    "xlsx": faFileExcel,
    "doc": faFileWord,
    "docx": faFileWord,
    "pptx": faFilePowerpoint,
    "txt": faFileText,
    "file": faFile
}

export const APIServer = {
    serverHost: "http://82.97.248.164:8080/api",
    tokenName: "COMPANYCLOUD_TOKEN",
    auth: {
        register: "/auth/register",
        login: "/auth/login",
    },
    user: {
        info: "/user/info",
        delete: "/user/delete",
        update: "/user/update",
        newPassword: "/user/new-password",
        limit: "/user/limit",
    },
    file: {
        list: "/file/list",
        upload: "/file/upload",
        downloadMany: "/file/download-many",
        download: "/file/download/",
        delete: "/file/",
        update: "/file/update/",
    },
    folder: {
        create: "/folder/create",
        list: "/folder/all",
        getFiles: "/folder/",
        update: "/folder/update/",
        delete: "/folder/",
    },
    admin: {
        userList: "/admin/user-list",
        userListSearch: "/admin/user-list/search",
        toggleApprove: "/admin/toggle-approve/",
    },
    url: {
        create: "/url/create",
        download: "/url-get/download/",
        files: "/url-get/files/",
    },
}

export const humanSizeBytes = function (bytes) {
    if (bytes === 0) {
        return "0.00 B";
    }

    const e = Math.floor(Math.log(bytes) / Math.log(1024));
    return (bytes / Math.pow(1024, e)).toFixed(2) + ' ' + ' KMGTP'.charAt(e) + 'B';
}


const App = () => {
    const [token, setToken] = useState("")

    return (
        <TokenContext.Provider value={{
            token, setToken
        }}>
            <>
                <Rounter />
            </>
        </TokenContext.Provider>
    );
}

export default App;
