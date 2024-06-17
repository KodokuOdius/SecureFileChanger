import React, { useContext, useEffect, useState } from "react";
import FileList from "../components/File/FileList";
import { SharedFilesContext, TokenContext, SelectionFilesContext } from "../context";
import FolderList from "../components/Folder/FolderList";
import HomePanel from "../components/HomePanel";
import FolderService from "../api/FolderService";
import { useAPI } from "../hooks/useAPI";
import FileService from "../api/FileService";
import Loader from "../components/Loader";


const Home = () => {
    const { token } = useContext(TokenContext)

    const [files, setFiles] = useState([])
    const [folders, setFolders] = useState([])

    const [getFiles, isFilesLoading, errorFiles] = useAPI(async () => {
        const filesList = await FileService.getList(token)
        if (filesList.length === 0) {
            return
        }
        setFiles(filesList)
    })

    const [getFolders, isFoldersLoading, errorFolders] = useAPI(async () => {
        const foldersList = await FolderService.getAll(token)
        if (foldersList.length === 0) {
            return
        }
        setFolders(foldersList)
    })

    useEffect(() => {
        getFiles()
        console.log(errorFiles)

        getFolders()
        console.log(errorFolders)
    }, []) // eslint-disable-line react-hooks/exhaustive-deps

    const onDeleteFile = (id) => {
        FileService.deleteFile(token, id)

        setFiles(files.filter(file => file.file_id !== id))
    }

    const addFolder = (folder) => {
        setFolders([...folders, folder])
    }

    const onDeleteFolder = (id) => {
        FolderService.deleteFolder(token, id)

        setFolders(folders.filter(folder => folder.id !== id))
    }

    const addFile = (file) => {
        setFiles([...files, file])
    }

    const [sharedFiles, setSharedFiles] = useState([])
    const [isShowSelect, setIsShowSelect] = useState(false)

    return (
        <SharedFilesContext.Provider value={{ sharedFiles, setSharedFiles }}>
            <SelectionFilesContext.Provider value={{ isShowSelect, setIsShowSelect }}>
                <div className="home">
                    <HomePanel addFolder={addFolder} addFile={addFile} />
                    <div className="home__workspace">
                        <h2 className="home__title" >Главная страница</h2>
                        <div className="home__folders">
                            {folders === null || folders.length === 0
                                ? <></>
                                : <FolderList folders={folders} onDeleteFolder={onDeleteFolder} />
                            }
                        </div>
                        <div className="home__files">
                            {isFilesLoading && isFoldersLoading &&
                                <Loader msg="Идёт загрузка информации" />
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

export default Home