import React from "react";

import mod from './Trackers.module.css';
import {NavLink} from "react-router-dom";
import logo from "../../Header/LogOut/LogOut.png";

const Trackers = () =>{
    return(
        <div className={mod.Con}>
            <div className={mod.Last}>
                <NavLink to='/lostor/details'>
                    <h3>Last executed</h3>
                </NavLink>
            </div>
            <div className={mod.Trackers}>
                <NavLink to='/lostor/trackers'>
                    <h3>Trackers</h3>
                </NavLink>
            </div>
            <div className={mod.Logout}>
                <NavLink to='/lostor/login'>
                    <img src={logo} className={mod.logout}/>
                </NavLink>
            </div>
        </div>
    )
}

export default Trackers;