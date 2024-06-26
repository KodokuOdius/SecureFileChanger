import React, { useContext, useEffect, useState } from "react";
import FileList from "../File/FileList";
import { SharedFilesContext, TokenContext, SelectionFilesContext } from "../../context";
import { useParams, Link } from "react-router-dom";
import FolderService from "../../api/FolderService";
import { useAPI } from "../../hooks/useAPI";
import Loader from "../Loader";
import HomePanel from "../HomePanel";
import FileService from "../../api/FileService";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faBackward } from "@fortawesome/free-solid-svg-icons";


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

    const addFile = (file) => {
        setFiles([...files, file])
    }

    const onDeleteFile = (id) => {
        FileService.deleteFile(token, id)

        setFiles(files.filter(file => file.file_id !== id))
    }

    useEffect(() => {
        getFiles()
        console.log(error)
    }, []) // eslint-disable-line react-hooks/exhaustive-deps

    const [sharedFiles, setSharedFiles] = useState([])
    const [isShowSelect, setIsShowSelect] = useState(false)

    return (
        <SharedFilesContext.Provider value={{ sharedFiles, setSharedFiles }}>
            <SelectionFilesContext.Provider value={{ isShowSelect, setIsShowSelect }}>
                <div className="home">
                    <HomePanel addFile={addFile} />
                    <div className="home__workspace">
                        <h2 className="folder__back">
                            <Link to="/">
                                <p className="back__icon">
                                    <FontAwesomeIcon icon={faBackward} size="1x" />
                                </p>
                                <span>Назад</span>
                            </Link>
                        </h2>
                        <div className="home__files">
                            {isLoading &&
                                <Loader msg="Загрузка документов" />
                            }
                            {files === null || files.length === 0
                                ? <h3>Нет Документов</h3>
                                : <FileList files={files} onDeleteFile={onDeleteFile} />
                            }
                        </div>
                    </div>
                </div>
            </SelectionFilesContext.Provider>
        </SharedFilesContext.Provider>
    )
}

export default FolderDetail