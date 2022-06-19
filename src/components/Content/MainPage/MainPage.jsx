import React from "react";
import logo from '../../Header/LogOut/LogOut.png';
import mod from './MainPage.module.css';
import {NavLink} from "react-router-dom";

const MainPage = (props) => {
    return (
        <div className={mod.Main}>
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
            <div className={mod.Filter}>
                <input placeholder="Filter"/>
            </div>
        </div>
    )
}

export default MainPage;