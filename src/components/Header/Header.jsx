import React from "react";

import mod from './Header.module.css';
import {NavLink} from "react-router-dom";
import Settings from "./Settings/Settings";
import LogOut from "./LogOut/LogOut";
import MainLogo from "./MainLogo/MainLogo";


const Header = (props) => {
    return (
        <div className={mod.Head}>
            <NavLink to='/lostor/torrents'>
                <MainLogo/>
            </NavLink>
            <Settings />
            <LogOut />
        </div>
    )
}

export default Header;