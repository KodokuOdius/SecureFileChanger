import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSearch } from "@fortawesome/free-solid-svg-icons";


const SearchBtn = (props) => {
    return (
        <i {...props} >
            <FontAwesomeIcon icon={faSearch} size="xl" />
        </i>
    )
}
export default SearchBtn