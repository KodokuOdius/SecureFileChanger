import React from "react";
import { Link } from "react-router-dom";

// {
//     "id": 10,
//     "folder_name": "root",
//     "is_root": true
// }
const FolderItem = ({ idx, folder, onDeleteFolder }) => {

    return (
        <div className="folder__item" id={folder.id}>
            <h3>
                Folder <Link to={"/folder/" + folder.id} >
                    {folder.folder_name}
                </Link >
            </h3>
            <button onClick={e => onDeleteFolder(folder.id)}>Удалить</button>
        </div>
    )
}

export default FolderItem