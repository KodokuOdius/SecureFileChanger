import React, { useContext, useEffect, useState } from "react";
import FileList from "../components/File/FileList";
import { TokenContext } from "../context";
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

    const onDeleteFile = (id) => { }

    const addFolder = (folder) => {
        setFolders([...folders, folder])
    }

    const onDeleteFolder = (id) => {
        FolderService.deleteFolder(token, id)

        setFolders(folders.filter(folder => folder.id !== id))
    }

    return (
        <div className="home">
            <h1>Главная страница</h1>
            <div className="home__workspace">
                <HomePanel addFolder={addFolder} />
                <div className="home__files">
                    {isFilesLoading && isFoldersLoading &&
                        <Loader msg="Идёт загрузка информации" />
                    }
                    {files === null || files.length === 0
                        ? <h3>Нет Документов</h3>
                        : <FileList files={files} onDeleteFile={onDeleteFile} />
                    }
                    <br />
                    {folders === null || folders.length === 0
                        ? <></>
                        : <FolderList folders={folders} onDeleteFolder={onDeleteFolder} />
                    }
                </div>
            </div>
        </div>
    )
}

export default Home