import React, { useRef } from "react";


const DeleteModal = ({ msg, onDelete, onClose }) => {
    const mainDiv = useRef()

    const onModalClick = (e) => {
        if (e.target === mainDiv.current) {
            onClose()
        }
    }

    return (
        <div className="delete__modal" ref={mainDiv} onClick={onModalClick}>
            <div className="modal__body">
                <p className="modal__title">{msg}</p>
                <div className="modal__btns">
                    <button onClick={onDelete}>Удалить</button>
                    <button onClick={onClose}>Отмена</button>
                </div>
            </div>
        </div>
    )
}
export default DeleteModal