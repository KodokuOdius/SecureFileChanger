import React from "react";
import { Link, useLocation } from "react-router-dom";

const RouterNavbar = ({ isLoginShow }) => {
    const location = useLocation()

    const noNavBarPath = ["/login", "/register", "/login/", "/register/"]

    if (noNavBarPath.some(path => path === location.pathname)) {
        return <></>
    }

    return (
        <nav className="router__navbar">
            <h2 className="router__logo">CloudCompany</h2>
            <ul className="navbar__links">
                <li className="links__item">
                    <Link to="/">Домашнаяя страница</Link>
                </li>
                <li className="links__item">
                    <Link to="/profile">Профиль</Link>
                </li>
                {isLoginShow &&
                    <li className="links__item">
                        <Link to="/login">Войти</Link>
                    </li>
                }
                {!isLoginShow &&
                    <li className="links__item">
                        <Link to="/logout">Выйти</Link>
                    </li>
                }
            </ul>
        </nav>
    )
}

export default RouterNavbar