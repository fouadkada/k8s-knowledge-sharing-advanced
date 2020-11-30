import React from "react";

import "./Menu.css";

const Menu = () => {
    return (
        <div className="ui fixed inverted menu">
            <div className="ui container">
                <a href="/" className="header item">
                    <img className="logo" src="logo192.png" alt="logo"/>Movies
                </a>
            </div>
        </div>
    );
}

export default Menu;