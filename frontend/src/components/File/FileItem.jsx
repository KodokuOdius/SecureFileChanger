import React, { useContext } from "react";
import { humanSizeBytes } from "../../App";
import { TokenContext } from "../../context";
import FileService from "../../api/FileService";


// defaultFileInput = {
//     file_id: 0, file_name: "",
//     size_bytes: 0, file_type: ""
// }
const FileItem = ({ idx, file, onDeleteFile }) => {
    const { token } = useContext(TokenContext)

    const onDownloadFile = async (e) => {
        const resp = await FileService.downloadFile(token, file.file_id)

        const url = window.URL.createObjectURL(new Blob([resp]))
        const link = document.createElement("a")
        link.setAttribute("href", url)
        link.setAttribute("download", file.file_name + file.file_type)
        document.body.appendChild(link)
        link.click()
        link.parentNode.removeChild(link)
    }

    return (
        <div className="file_item" id={file.file_id}>
            <h4>{idx}# {file.file_name + file.file_type}</h4>
            <p>{humanSizeBytes(file.size_bytes)}</p>
            {onDeleteFile &&
                <>
                    <button onClick={() => onDeleteFile(file.file_id)}>Удалить</button>
                    <button onClick={onDownloadFile}>Скачать</button>
                </>
            }
        </div>
    )
}
export default FileItem