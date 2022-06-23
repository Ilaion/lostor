import React from "react";
import mod from './SettingsPage.module.css';
import {NavLink} from "react-router-dom";
import arrowLeft from './arrowLeft.png';

const SettingsPage = () => {
    return (
        <div className={mod.main}>


            <NavLink to='lostor/torrents'>
                <img src={arrowLeft} alt='back to main page'/>
            </NavLink>
            <div>Settnigs</div>
            <NavLink to='lostor/torrents'>
                <img src={arrowLeft} alt='back to main page'/>
            </NavLink>
            <div>General</div>
            <NavLink to='lostor/torrents'>
                <img src={arrowLeft} alt='back to main page'/>
            </NavLink>
            <div>Trackers</div>
            <NavLink to='lostor/torrents'>
                <img src={arrowLeft} alt='back to main page'/>
            </NavLink>
            <div>Notifiers</div>
            <NavLink to='lostor/torrents'>
                <img src={arrowLeft} alt='back to main page'/>
            </NavLink>
            <div>Clients</div>
            <NavLink to='lostor/torrents'>
                <img src={arrowLeft} alt='back to main page'/>
            </NavLink>
            <div>Authentication</div>
            <NavLink to='lostor/torrents'>
                <img src={arrowLeft} alt='back to main page'/>
            </NavLink>
            <div>About</div>

        </div>
    )
}

export default SettingsPage;