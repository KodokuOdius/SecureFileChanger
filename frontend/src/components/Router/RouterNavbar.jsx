import React from "react";
import { Link, useLocation } from "react-router-dom";

const RouterNavbar = ({ isLoginShow }) => {
    const location = useLocation()

    const noNavBarPath = ["/login", "/register", "/login/", "/register/"]

    if (noNavBarPath.some(path => path === location.pathname)) {
        return <></>
    }

    if (location.pathname.includes("/d/")) {
        return <></>
    }

    return (
        <nav className="router__navbar">
            <div className="navbar__list">
                <h2 className="router__logo">
                    <Link to="/">CloudCompany</Link>
                </h2>
                <ul className="navbar__links">
                    <li className="links__item">
                        <Link to="/">Домашнаяя страница</Link>
                    </li>
                    <li className="links__item">
                        <Link to="/profile">Профиль</Link>
                    </li>
                    <li className="links__item">
                        <Link to="/logout">Выйти</Link>
                    </li>
                </ul>
            </div>
        </nav>
    )
}

export default RouterNavbar