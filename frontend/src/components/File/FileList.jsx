import React from "react";
import FileItem from "./FileItem";


const FileList = ({ files, onDeleteFile }) => {
    return (
        <div className="file__list">
            {
                files.map((file, i) =>
                    <FileItem idx={i} file={file} key={file.file_id} onDeleteFile={onDeleteFile} />
                )
            }
        </div>
    )
}
export default FileList