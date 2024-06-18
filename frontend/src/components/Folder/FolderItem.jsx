import React, { useState } from "react";
import { Link } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import DeleteBtn from "../Buttons/DeleteBtn";
import DeleteModal from "../Modal/DeleteModal";

// {
//     "id": 10,
//     "folder_name": "root",
//     "is_root": true
// }
const FolderItem = ({ idx, folder, onDeleteFolder }) => {
    const [isShowModal, setIsShowModal] = useState(false)
    const onShowDelete = () => setIsShowModal(true)

    return (
        <div className="folder__item" id={folder.id}>
            {isShowModal &&
                <DeleteModal
                    msg="Вы точно хотите удалить эту директорию?"
                    onClose={() => setIsShowModal(false)}
                    onDelete={() => onDeleteFolder(folder.id)}
                />
            }
            <Link to={"/folder/" + folder.id} className="folder__name" >
                <p className="folder__icon">
                    <FontAwesomeIcon icon="fa-solid fa-folder" size="2x" />
                </p>
                <h3>{folder.folder_name}</h3>
            </Link >
            <div className="folder__btns">
                <DeleteBtn onClick={onShowDelete} />
            </div>
        </div>
    )
}

export default FolderItem