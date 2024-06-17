import React, { useContext, useState } from "react";
import { TokenContext } from "../../context";
import FolderService from "../../api/FolderService";


const FolderCreateForm = ({ setShow, addFolder }) => {
    const { token } = useContext(TokenContext)

    const [folderName, setFolderName] = useState("")
    const createFolder = async (e) => {
        e.preventDefault()

        if (folderName === "") {
            return
        }

        try {
            const newFolder = await FolderService.createFolder(token, folderName)
            addFolder(newFolder)

            setFolderName("")
            setShow(false)
        }
        catch (e) {
            return
        }
    }

    return (
        <form className="folder_create_form">
            <input
                value={folderName}
                onChange={e => setFolderName(e.target.value)}
                type="text"
                minLength={1}
                maxLength={50}
                placeholder="Название директории"
            />
            <button onClick={createFolder}>Создать</button>
            <button onClick={() => setShow(false)}>Закрыть</button>
        </form>
    )
}

export default FolderCreateForm