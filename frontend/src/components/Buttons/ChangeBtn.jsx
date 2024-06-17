import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faPen } from "@fortawesome/free-solid-svg-icons";


const ChangeBtn = (props) => {
    return (
        <i {...props} title="Изменить" >
            <FontAwesomeIcon icon={faPen} size="xl" />
        </i>
    )
}
export default ChangeBtn