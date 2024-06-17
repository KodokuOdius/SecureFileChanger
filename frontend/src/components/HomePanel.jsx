import React, { useContext, useRef, useState } from "react";
import FolderCreateForm from "./Folder/FolderCreateForm";
import FileService from "../api/FileService";
import { TokenContext } from "../context";
import ShareCreateForm from "./Share/ShareCreateForm";
import { useParams } from "react-router-dom";
import DownloadDeleteForm from "./File/DownloadDeleteForm";


const HomePanel = ({ addFolder, addFile }) => {
    const { token } = useContext(TokenContext)
    const params = useParams()

    const [isShowFolderCreate, setIsShowFolderCreate] = useState(false)
    const toggleShowFolderCreate = () => {
        setIsShowFolderCreate(!isShowFolderCreate)
    }

    const [isShowShareForm, setIsShowShareForm] = useState(false)
    const [isShowDownloadDelete, setIsShowDownloadDelete] = useState(false)

    const fileInputRef = useRef()

    const onClickUpload = (e) => {
        fileInputRef.current.value = ""
        e.preventDefault()
        return fileInputRef.current.click()
    }

    const onUploadFile = async (e) => {
        try {
            const file = e.target.files[0]
            const resp = await FileService.uploadFile(token, file, params.folderId)
            addFile(resp)
        }
        catch (e) {
            if (e.response) {
                console.log(e.response.data.message)
            }
            return
        }
    }

    const onToggleShare = () => setIsShowShareForm(!isShowShareForm)

    const onSelect = () => setIsShowDownloadDelete(!isShowDownloadDelete)

    return (
        <div className="home__panel">
            <div className="panel__item">
                <button onClick={onSelect}>Выбрать документы</button>
                {isShowDownloadDelete &&
                    <DownloadDeleteForm setShow={setIsShowDownloadDelete} />
                }
            </div>
            {!params.folderId &&
                <div className="panel__item">
                    <button onClick={toggleShowFolderCreate}>Создать Директорию</button>
                    {isShowFolderCreate &&
                        <FolderCreateForm setShow={setIsShowFolderCreate} addFolder={addFolder} />
                    }
                </div>
            }
            <div className="panel__item">
                <form method="post" encType="multipart/form-data">
                    <input onChange={onUploadFile} ref={fileInputRef} type="file" hidden multiple={true} />
                    <button onClick={onClickUpload}>Загрузить Документ</button>
                </form>
            </div>
            <div className="panel__item">
                <button onClick={onToggleShare}>Поделиться</button>
                {isShowShareForm &&
                    <ShareCreateForm setShow={setIsShowShareForm} />
                }
            </div>
        </div>
    )
}

export default HomePanel