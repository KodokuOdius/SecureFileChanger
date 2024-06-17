import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCancel } from "@fortawesome/free-solid-svg-icons";


const CancelBtn = (props) => {
    return (
        <i {...props} title="Отмена" >
            <FontAwesomeIcon icon={faCancel} size="xl" />
        </i>
    )
}
export default CancelBtn