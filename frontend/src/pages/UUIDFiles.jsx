import React, { useEffect, useState } from "react";
import { Link, Navigate, useNavigate, useParams } from "react-router-dom";
import { useAPI } from "../hooks/useAPI";
import UrlService from "../api/UrlService";
import Loader from "../components/Loader";
import FileList from "../components/File/FileList";
import { APIServer } from "../App";


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


    return (
        <div className="files__uuid">
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
                        : <FileList files={files} />
                    }
                    <div className="files__uuid__btn">
                        <Link to={APIServer.serverHost + APIServer.url.download + uuid} download >Скачать</Link>
                    </div>
                </>
            }
        </div>
    )
}
export default UUIDFiles