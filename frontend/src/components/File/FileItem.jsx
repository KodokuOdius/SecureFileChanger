import React, { useContext, useState } from "react";
import { humanSizeBytes } from "../../App";
import { SelectionFilesContext, SharedFilesContext, TokenContext } from "../../context";
import FileService from "../../api/FileService";
import { useAPI } from "../../hooks/useAPI";


// defaultFileInput = {
//     file_id: 0, file_name: "",
//     size_bytes: 0, file_type: ""
// }
const FileItem = ({ idx, file, onDeleteFile }) => {
    const { token } = useContext(TokenContext)
    const { isShowSelect } = useContext(SelectionFilesContext)
    const { sharedFiles, setSharedFiles } = useContext(SharedFilesContext)

    const [fileData, setFileData] = useState({ ...file, old_name: file.file_name })
    const [isEdit, setIsEdit] = useState(false)

    const onDownloadFile = async (e) => {
        const resp = await FileService.downloadFile(token, fileData.file_id)

        const url = window.URL.createObjectURL(new Blob([resp]))
        const link = document.createElement("a")
        link.setAttribute("href", url)
        link.setAttribute("download", fileData.file_name + fileData.file_type)
        document.body.appendChild(link)
        link.click()
        link.parentNode.removeChild(link)
    }

    const onToggleEditFile = (e) => {
        if (isEdit) {
            setFileData({ ...fileData, file_name: fileData.old_name })
        }
        setIsEdit(!isEdit)
    }

    const [updateFileName, ,] = useAPI(async () => {
        const res = await FileService.updateName(token, file.file_id, fileData.file_name)
        if (res === null) {
            setFileData({ ...fileData, file_name: fileData.old_name })
        }
    })

    const onSaveName = (e) => {
        if (fileData.file_name.trim().length === 0) {
            setFileData({ ...fileData, file_name: fileData.old_name })
            setIsEdit(false)
            return
        }

        updateFileName()
        setFileData({ ...fileData, old_name: fileData.file_name })
        setIsEdit(false)
    }

    const onCheckBox = (e) => {
        if (e.target.checked) {
            setSharedFiles([...sharedFiles, file.file_id])
            return
        }

        setSharedFiles(sharedFiles.filter(fileId => fileId !== file.file_id))
    }

    return (
        <div className="file__item" id={file.file_id}>
            <h4>
                <input
                    className="file__item__inp"
                    type="text"
                    disabled={!isEdit}
                    value={fileData.file_name}
                    onChange={e => setFileData({ ...fileData, file_name: e.target.value })}
                />
                <span>{file.file_type}</span>
            </h4>
            <p>{humanSizeBytes(file.size_bytes)}</p>
            {onDeleteFile &&
                <>
                    <button onClick={() => onDeleteFile(file.file_id)}>Удалить</button>
                    <button onClick={onDownloadFile}>Скачать</button>
                    <button onClick={onToggleEditFile}>
                        {!isEdit ? "Изменить" : "Отмена"}
                    </button>
                    {isEdit
                        && <button onClick={onSaveName}>Сохранить</button>
                    }
                </>
            }
            {isShowSelect &&
                <div className="file__checking">
                    <input
                        type="checkbox"
                        onChange={onCheckBox}
                    />
                </div>
            }
        </div>
    )
}
export default FileItem