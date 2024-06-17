import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import React from "react";
import { faTrashCan } from "@fortawesome/free-solid-svg-icons";


const DeleteBtn = (props) => {
    return (
        <i {...props} title="Удалить" >
            <FontAwesomeIcon icon={faTrashCan} size="xl" />
        </i>
    )
}
export default DeleteBtn