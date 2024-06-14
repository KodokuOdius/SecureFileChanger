import React, { useContext, useRef, useState } from "react";
import FolderCreateForm from "./Folder/FolderCreateForm";
import FileService from "../api/FileService";
import { TokenContext } from "../context";
import ShareCreateForm from "./Share/ShareCreateForm";
import { useParams } from "react-router-dom";


const HomePanel = ({ addFolder, addFile }) => {
    const { token } = useContext(TokenContext)
    const params = useParams()

    const [isShowFolderCreate, setIsShowFolderCreate] = useState(false)
    const toggleShowFolderCreate = () => {
        setIsShowFolderCreate(!isShowFolderCreate)
    }

    const [isShowShareForm, setIsShowShareForm] = useState(false)

    const fileInputRef = useRef()

    const onClickUpload = (e) => {
        e.preventDefault()
        return fileInputRef.current.click()
    }

    const onUploadFile = async (e) => {
        const file = e.target.files[0]
        const resp = await FileService.uploadFile(token, file, params.folderId)
        addFile(resp)
    }

    const onToggleShare = () => setIsShowShareForm(!isShowShareForm)

    return (
        <div className="home__panel">
            {!params.folderId &&
                <button onClick={toggleShowFolderCreate} >Создать Директорию</button>
            }
            {isShowFolderCreate &&
                <FolderCreateForm setShow={setIsShowFolderCreate} addFolder={addFolder} />
            }
            <form method="post" encType="multipart/form-data" >
                <input onChange={onUploadFile} ref={fileInputRef} type="file" hidden multiple={false} />
                <button onClick={onClickUpload}>Загрузить Документ</button>
            </form>
            <button onClick={onToggleShare}>Поделиться</button>
            {isShowShareForm &&
                <ShareCreateForm setShow={setIsShowShareForm} />
            }
        </div>
    )
}

export default HomePanel