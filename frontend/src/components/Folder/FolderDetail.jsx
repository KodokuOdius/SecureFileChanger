import React, { useContext, useEffect, useState } from "react";
import FileList from "../File/FileList";
import { TokenContext } from "../../context";
import { useParams } from "react-router-dom";
import FolderService from "../../api/FolderService";
import { useAPI } from "../../hooks/useAPI";
import Loader from "../Loader";


const FolderDetail = () => {
    const params = useParams()
    const folderId = params.folderId
    const { token } = useContext(TokenContext)
    const [files, setFiles] = useState([])

    const [getFiles, isLoading, error] = useAPI(async () => {
        const filesList = await FolderService.getFiles(token, folderId)
        if (filesList.length === 0) {
            return
        }
        setFiles(filesList)
    })

    const onDeleteFile = (id) => { }

    useEffect(() => {
        getFiles()
        console.log(error)
    }, []) // eslint-disable-line react-hooks/exhaustive-deps

    return (
        <div className="home__files">
            {isLoading &&
                <Loader msg="Загрузка документов" />
            }
            {files === null || files.length === 0
                ? <h3>Нет Документов</h3>
                : <FileList files={files} onDeleteFile={onDeleteFile} />
            }
        </div>
    )
}

export default FolderDetail