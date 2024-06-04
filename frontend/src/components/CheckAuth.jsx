import React, { useContext, useEffect } from "react";
import { TokenContext } from "../context";
import { Navigate } from "react-router-dom";
import { APIServer } from "../App";


// qwe1d2311@email.com
const AuthWiddleware = ({ children }) => {
    const { token, setToken } = useContext(TokenContext)

    useEffect(() => {
        const storageToken = localStorage.getItem(APIServer.tokenName)
        console.log("AuthWiddleware storageToken = ", storageToken)

        if (storageToken === null || token === null) {
            return () => { }
        }

        if (storageToken !== "" && token === "") {
            setToken(storageToken)
            return () => { }
        }

        if (storageToken !== token) {
            localStorage.setItem(APIServer.tokenName, token)
            return () => { }
        }

    }, []) // eslint-disable-line react-hooks/exhaustive-deps

    console.log("AuthWiddleware token = ", token)

    if (token === null || token === "") {
        return <Navigate to="/login" />
    }

    return children
}

export default AuthWiddleware