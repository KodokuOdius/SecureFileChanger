import React, { useContext, useEffect } from "react";
import { SelectionFilesContext, SharedFilesContext, TokenContext } from "../../context";
import FileService from "../../api/FileService";


const DownloadDeleteForm = ({ setShow }) => {
    const { token } = useContext(TokenContext)
    const { setIsShowSelect } = useContext(SelectionFilesContext)
    const { sharedFiles, setSharedFiles } = useContext(SharedFilesContext)

    useEffect(() => {
        setIsShowSelect(true)
        return () => {
            setIsShowSelect(false)
            setSharedFiles([])
        }
    }, []) // eslint-disable-line react-hooks/exhaustive-deps

    const onDownload = async () => {
        if (sharedFiles.length === 0) {
            return
        }

        if (sharedFiles.length === 1) {
            console.log(sharedFiles)
            const curFile = sharedFiles[0]
            const resp = await FileService.downloadFile(token, curFile.file_id)
            console.log(resp)
            const url = window.URL.createObjectURL(new Blob([resp]))
            const link = document.createElement("a")
            link.setAttribute("href", url)
            link.setAttribute("download", curFile.file_name + curFile.file_type)
            document.body.appendChild(link)
            link.click()
            link.parentNode.removeChild(link)

        } else {
            const fileIds = sharedFiles.map(f => f.file_id)
            const resp = await FileService.downloadFiles(token, fileIds)
            const url = window.URL.createObjectURL(new Blob([resp]))
            const link = document.createElement("a")
            link.setAttribute("href", url)
            link.setAttribute("download", "archive.zip")
            document.body.appendChild(link)
            link.click()
            link.parentNode.removeChild(link)
        }
    }

    return (
        <div className="download__form">
            <button onClick={onDownload}>Скачать</button>
            <button onClick={() => setShow(false)}>Закрыть</button>
        </div>
    )
}
export default DownloadDeleteForm;