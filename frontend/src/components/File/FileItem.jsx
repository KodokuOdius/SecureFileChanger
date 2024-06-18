import React, { useContext, useRef, useState } from "react";
import { FilesIcons, humanSizeBytes } from "../../App";
import { SelectionFilesContext, SharedFilesContext, TokenContext } from "../../context";
import FileService from "../../api/FileService";
import { useAPI } from "../../hooks/useAPI";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import DeleteBtn from "../Buttons/DeleteBtn";
import DownloadBtn from "../Buttons/DownloadBtn";
import CancelBtn from "../Buttons/CancelBtn";
import SaveBtn from "../Buttons/SaveBtn";
import ChangeBtn from "../Buttons/ChangeBtn";
import DeleteModal from "../Modal/DeleteModal";


// defaultFileInput = {
//     file_id: 0, file_name: "",
//     size_bytes: 0, file_type: ""
// }
const FileItem = ({ idx, file, onDeleteFile }) => {
    const { token } = useContext(TokenContext)
    const { isShowSelect } = useContext(SelectionFilesContext)
    const { sharedFiles, setSharedFiles } = useContext(SharedFilesContext)

    const fileName = useRef()

    const [fileData, setFileData] = useState({ ...file, old_name: file.file_name })
    const [isEdit, setIsEdit] = useState(false)

    const [isShowModal, setIsShowModal] = useState(false)

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
        fileName.current.focus()

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
            setSharedFiles([...sharedFiles, file])
            return
        }

        setSharedFiles(sharedFiles.filter(f => f.file_id !== file.file_id))
    }

    const onChangeName = (e) => {
        setFileData({ ...fileData, file_name: e.target.value })
    }

    const onShowDelete = () => setIsShowModal(true)

    return (
        <div className="file__item" id={file.file_id}>
            {isShowModal &&
                <DeleteModal
                    msg="Вы точно хотите удалить этот документ?"
                    onClose={() => setIsShowModal(false)}
                    onDelete={() => onDeleteFile(file.file_id)}
                />
            }
            <div className="item__name">
                <div className="file__name">
                    <p className="file__icon">
                        {FilesIcons[file.file_type.replace(".", "")] !== ""
                            ? <FontAwesomeIcon icon={FilesIcons[file.file_type.replace(".", "")]} size="xl" />
                            : <FontAwesomeIcon icon={FilesIcons["file"]} size="xl" />
                        }
                    </p>
                    <input autoFocus
                        ref={fileName}
                        className="item__inp"
                        type="text"
                        disabled={!isEdit}
                        value={fileData.file_name}
                        onChange={onChangeName}
                        minLength={1}
                        maxLength={100}
                        size={fileData.file_name.length}
                    />
                </div>
                <div className="item__info">
                    <p>Type: {file.file_type.replace(".", "")}</p>
                    <p>{humanSizeBytes(file.size_bytes)}</p>
                </div>
            </div>
            {onDeleteFile &&
                <div className="item__btns">
                    {isShowSelect &&
                        <div className="item__checking">
                            <input
                                type="checkbox"
                                onChange={onCheckBox}
                            />
                        </div>
                    }
                    <DownloadBtn onClick={onDownloadFile} />
                    {isEdit &&
                        <SaveBtn onClick={onSaveName} />
                    }
                    {!isEdit
                        ? <ChangeBtn onClick={onToggleEditFile} />
                        : <CancelBtn onClick={onToggleEditFile} />
                    }
                    <DeleteBtn onClick={onShowDelete} />
                </div>
            }
        </div>
    )
}
export default FileItem