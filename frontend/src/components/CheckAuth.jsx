import React, { useContext, useEffect } from "react";
import { TokenContext } from "../context";
import { useNavigate } from "react-router-dom";
import { APIServer } from "../App";


// qwe1d2311@email.com
const AuthWiddleware = ({ children }) => {
    const { token, setToken } = useContext(TokenContext)
    const navigate = useNavigate()

    useEffect(() => {
        const storageToken = localStorage.getItem(APIServer.tokenName)
        // console.log("AuthWiddleware storageToken = ", storageToken)
        console.log("AuthWiddleware")
        console.log(storageToken)
        console.log(token)
        if (storageToken === null || token === null || token === "") {
            navigate("/login")
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

    // console.log("AuthWiddleware token = ", token)

    return children
}

export default AuthWiddleware