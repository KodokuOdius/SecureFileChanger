import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";


const NotFound = () => {
    const navigate = useNavigate()

    useEffect(() => {
        navigate("/")
    }, []) // eslint-disable-line react-hooks/exhaustive-deps

    return (
        <div className="notfounc__page">
            <p className="notfound__title">Ничего не найдено</p>
        </div>
    )
}
export default NotFound;