import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCheck } from "@fortawesome/free-solid-svg-icons";


const SaveBtn = (props) => {
    return (
        <i {...props} title="Сохранить" >
            <FontAwesomeIcon icon={faCheck} size="xl" />
        </i>
    )
}
export default SaveBtn