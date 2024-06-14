import React, { useContext, useEffect, useState } from "react";
import UrlService from "../../api/UrlService";
import { SelectionFilesContext, SharedFilesContext, TokenContext } from "../../context";

const hourLiveValues = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23]

const ShareCreateForm = ({ setShow }) => {
    const { token } = useContext(TokenContext)
    const { setIsShowSelect } = useContext(SelectionFilesContext)
    const { sharedFiles, setSharedFiles } = useContext(SharedFilesContext)

    useEffect(() => {
        setIsShowSelect(true)
        return () => {
            setIsShowSelect(false)
            setSharedFiles([])
        }
    }, []) // eslint-disable-line react-hooks/exhaustive-deps

    const [hourLive, setHourLive] = useState(0)
    const [shareURL, setShareURL] = useState("")

    const onCreate = async () => {
        if (sharedFiles.length === 0) {
            return
        }

        const resUrl = await UrlService.createUrl(token, sharedFiles, hourLive)
        setShareURL(resUrl)
    }

    return (
        <div className="share__form">
            <h3>Поделиться документами</h3>
            <p>Время жизни ссылки: </p>
            <select
                name="url_hour_live"
                id="url_hour_live-select"
                onChange={e => setHourLive(Number(e.target.value))}
            >
                {hourLiveValues.map(hour =>
                    <option value={hour} key={hour} >{hour}h</option>
                )}
            </select>
            {shareURL !== "" &&
                <>
                    <p>Ссылка для скачивания: </p>
                    <p>{shareURL}</p>
                </>
            }
            <button onClick={onCreate}>Создать ссылку</button>
            <button onClick={() => setShow(false)}>Закрыть</button>
        </div>
    )
}
export default ShareCreateForm
