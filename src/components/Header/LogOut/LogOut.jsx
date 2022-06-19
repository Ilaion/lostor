import React from "react";
import logo from'./LogOut.png';
import mod from './LogOut.module.css';
import {NavLink} from "react-router-dom";

const LogOut = () => {
    return(
        <div className={mod.LogOut}>
            <NavLink to='/lostor/login'>
                <img src={logo} alt='LogOut'/>
            </NavLink>
        </div>
    )
}

export default LogOut;