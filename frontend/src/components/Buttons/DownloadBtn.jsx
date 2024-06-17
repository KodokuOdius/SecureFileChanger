import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faDownload } from "@fortawesome/free-solid-svg-icons";


const DownloadBtn = (props) => {
    return (
        <i {...props} title="Скачать" >
            <FontAwesomeIcon icon={faDownload} size="xl" />
        </i>
    )
}
export default DownloadBtn;