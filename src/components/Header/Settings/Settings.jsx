import React from "react";
import logo from './Settings.png';
import mod from './Settings.module.css';
import {NavLink} from "react-router-dom";

const Settings = (props) => {
    return (
        <div className={mod.settings}>
            <NavLink to='/lostor/settings'>
                <img src={logo} alt='Settings'/>
            </NavLink>
        </div>
    )
}

export default Settings;