import React from "react";
import FolderItem from "./FolderItem";


const FolderList = ({ folders, onDeleteFolder }) => {

    return (
        <div className="folder__list">
            {
                folders.map((folder, i) => {
                    return <FolderItem
                        idx={i}
                        folder={folder}
                        key={folder.id}
                        onDeleteFolder={onDeleteFolder}
                    />
                })
            }
        </div>
    )
}

export default FolderList