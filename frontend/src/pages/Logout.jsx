import React, { useContext } from "react";
import { TokenContext } from "../context";
import { Navigate } from "react-router-dom";
import { APIServer } from "../App";


const Logout = ({ setIsLoginShow }) => {
    const { token, setToken } = useContext(TokenContext)
    if (token !== "") {
        setToken("")
        localStorage.setItem(APIServer.tokenName, "")
        setIsLoginShow(true)
        return <Navigate to="/login" />
    }

    return <Navigate to="/" />
}

export default Logout