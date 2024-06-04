import React from "react";
import { Link } from "react-router-dom";

const RouterNavbar = ({ isLoginShow }) => {
    return (
        <nav className="router__navbar">
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