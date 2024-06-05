import React, { useContext, useRef, useState } from "react";
import FolderCreateForm from "./Folder/FolderCreateForm";
import FileService from "../api/FileService";
import { TokenContext } from "../context";


const HomePanel = ({ addFolder, addFile }) => {
    const { token } = useContext(TokenContext)

    const [isShowFolderCreate, setIsShowFolderCreate] = useState(false)
    const toggleShowFolderCreate = () => {
        setIsShowFolderCreate(!isShowFolderCreate)
    }

    const fileInputRef = useRef()

    const onClickUpload = (e) => {
        e.preventDefault()
        return fileInputRef.current.click()
    }

    const onUploadFile = async (e) => {
        const file = e.target.files[0]
        const resp = await FileService.uploadFile(token, file)
        console.log(resp)
        addFile(resp)
    }

    return (
        <div className="home__panel">
            <button onClick={toggleShowFolderCreate} >Создать Директорию</button>
            {isShowFolderCreate
                ? <FolderCreateForm setShow={setIsShowFolderCreate} addFolder={addFolder} />
                : <></>
            }
            <br />
            <form method="post" enctype="multipart/form-data" >
                <input onChange={onUploadFile} ref={fileInputRef} type="file" hidden multiple={false} />
                <button onClick={onClickUpload}>Загрузить Документ</button>
            </form>
        </div>
    )
}

export default HomePanel