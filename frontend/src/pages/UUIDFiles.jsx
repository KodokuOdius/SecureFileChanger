import React, { useEffect, useState } from "react";
import { Navigate, useNavigate, useParams } from "react-router-dom";
import { useAPI } from "../hooks/useAPI";
import UrlService from "../api/UrlService";
import Loader from "../components/Loader";
import FileList from "../components/File/FileList";
import DownloadBtn from "../components/Buttons/DownloadBtn";


const isValidUUID = (uuid) => {
    const uuidRegex = /^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/i;
    return typeof uuid === "string" && uuidRegex.test(uuid);
}

const UUIDFiles = () => {
    const params = useParams()
    const navigate = useNavigate()
    const uuid = params.uuid

    const [isNotFound, setIsNotFound] = useState(false)
    const [files, setFiles] = useState([])

    const [getFiles, isLoading,] = useAPI(async () => {
        try {
            const resp = await UrlService.getFiles(uuid)
            setFiles(resp)
        } catch (e) {
            if (e.response && e.response.status === 404) {
                setIsNotFound(true)
            }
        }
    })

    useEffect(() => {
        if (!isValidUUID(uuid)) {
            return () => navigate("/")
        }
        getFiles()
    }, []) // eslint-disable-line react-hooks/exhaustive-deps

    const onClickLink = async () => {
        let resp
        try {
            resp = await UrlService.downloadFiles(uuid)
        } catch (error) {
            setIsNotFound(true)
            return
        }

        const url = window.URL.createObjectURL(new Blob([resp]))
        const link = document.createElement("a")
        link.setAttribute("href", url)

        if (files.length === 1) {
            const file = files[0]
            link.setAttribute("download", file.file_name + file.file_type)
        } else {
            link.setAttribute("download", "archive.zip")
        }
        document.body.appendChild(link)
        link.click()
        link.parentNode.removeChild(link)
    }

    return (
        <div className="files__uuid">
            <h3 className="uuid__title" >Список документов</h3>
            {!isValidUUID(uuid) &&
                <Navigate to="/" />
            }
            {isLoading
                ? <Loader msg="Загрузка документов" />
                : <>
                    {isNotFound
                        ? <div className="notfound__uuid">
                            <p>Ссылка не действительна</p>
                        </div>
                        : <div className="uuid__workspace">
                            <FileList files={files} />
                            <div className="files__uuid__btns">
                                <button onClick={onClickLink}>
                                    <DownloadBtn />
                                    <span>Скачать</span>
                                    <DownloadBtn />
                                </button>
                            </div>
                        </div>
                    }
                </>
            }
        </div>
    )
}
export default UUIDFiles