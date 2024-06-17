import React from "react";
import { Link } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import DeleteBtn from "../Buttons/DeleteBtn";

// {
//     "id": 10,
//     "folder_name": "root",
//     "is_root": true
// }
const FolderItem = ({ idx, folder, onDeleteFolder }) => {
    return (
        <div className="folder__item" id={folder.id}>
            <Link to={"/folder/" + folder.id} className="folder__name" >
                <p className="folder__icon">
                    <FontAwesomeIcon icon="fa-solid fa-folder" size="2x" />
                </p>
                <h3>{folder.folder_name}</h3>
            </Link >
            <div className="folder__btns">
                <DeleteBtn onClick={e => onDeleteFolder(folder.id)} />
            </div>
        </div>
    )
}

export default FolderItem