import React from "react";
import logo from '../../Header/LogOut/LogOut.png';
import mod from './MainPage.module.css';
import {NavLink} from "react-router-dom";
import plus from './plus.png';

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
                    <NavLink to='/lostor/trackers' >
                        <h3>Trackers</h3>
                    </NavLink>
                </div>
                <div className={mod.Logout}>
                    <NavLink to='/lostor/login'>
                        <img src={logo} className={mod.logout}/>
                    </NavLink>
                </div>
            </div>
            <div className={mod.sort}>
                <label>Sort</label>
            </div>
            <div className={mod.Filter}>
                <input placeholder="Filter"/>
                <select className={mod.droper}>
                    <option>Last Update</option>
                    <option>Name</option>
                </select>
                <NavLink to='/lostor/torrents'>
                    <img src={plus} className={mod.plus}/>
                </NavLink>
            </div>

        </div>
    )
}

export default MainPage;