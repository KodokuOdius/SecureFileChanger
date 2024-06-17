import React, { useState } from "react";
import { BrowserRouter, Route, Routes, Navigate } from "react-router-dom";
import Home from "../../pages/Home";
import Login from "../../pages/Login";
import Logout from "../../pages/Logout";
import Profile from "../../pages/Profile";
import AuthWiddleware from "../CheckAuth";
import Register from "../../pages/Register";
import RouterNavbar from "./RouterNavbar";
import AdminPanel from "../../pages/AdminPanel";
import UUIDFiles from "../../pages/UUIDFiles";
import FolderDetail from "../Folder/FolderDetail";


const Rounter = () => {
    const [isLoginShow, setIsLoginShow] = useState(false)

    return (
        <BrowserRouter>
            <div className="nav__roter">
                <RouterNavbar isLoginShow={isLoginShow} />
                <Routes>
                    <Route path="/login" element={<Login setIsLoginShow={setIsLoginShow} />} />
                    <Route path="/logout" element={<Logout setIsLoginShow={setIsLoginShow} />} />
                    <Route path="/register" element={<Register />} />
                    <Route path="/folder/:folderId" element={
                        <AuthWiddleware>
                            <FolderDetail />
                        </AuthWiddleware>
                    } />
                    <Route path="/profile"
                        redirect="/"
                        validator={Profile}
                        element={
                            <AuthWiddleware>
                                <Profile />
                            </AuthWiddleware>
                        } />
                    <Route path="/admin-panel" element={
                        <AuthWiddleware>
                            <AdminPanel />
                        </AuthWiddleware>
                    } />
                    <Route path="/" element={
                        <AuthWiddleware>
                            <Home />
                        </AuthWiddleware>
                    } />
                    <Route path="/d/:uuid" element={
                        <UUIDFiles />
                    } />
                    <Route path="/*" element={<Navigate to="/" />} />
                </Routes>
            </div>
        </BrowserRouter >
    )
}


export default Rounter