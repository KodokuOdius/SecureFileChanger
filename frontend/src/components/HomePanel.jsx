import React, { useState } from "react";
import FolderCreateForm from "./Folder/FolderCreateForm";


const HomePanel = ({ addFolder }) => {
    const [isShowFolderCreate, setIsShowFolderCreate] = useState(false)
    const toggleShowFolderCreate = () => {
        setIsShowFolderCreate(!isShowFolderCreate)
    }

    return (
        <div className="home__panel">
            <button onClick={toggleShowFolderCreate} >Создать Директорию</button>
            {isShowFolderCreate
                ? <FolderCreateForm setShow={setIsShowFolderCreate} addFolder={addFolder} />
                : <></>
            }
            <br />
            <button>Загрузить Документ</button>
        </div>
    )
}

export default HomePanel