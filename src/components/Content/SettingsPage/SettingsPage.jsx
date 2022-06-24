import React from "react";
import mod from './SettingsPage.module.css';
import {NavLink} from "react-router-dom";
import arrowLeft from './arrowLeft.png';
import general from './general.png';
import trackers from './trackers.png';
import notification from './notification.png';
import clients from './clients.png';
import authentication from './authentication.png';
import about from './about.png';

const SettingsPage = () => {
    return (
        <div className={mod.main}>

            <div className={mod.reols}>
                <NavLink to='lostor/torrents'>
                    <img src={arrowLeft} alt='back to main page'/>
                </NavLink>
            </div>
            <div>Settnigs</div>
            <div className={mod.reols}>
            <NavLink to='lostor/torrents'>
                <img src={general} alt='general settings'/>
            </NavLink>
            </div>
            <div>General</div>
            <div className={mod.reols}>
            <NavLink to='lostor/torrents'>
                <img src={trackers} alt='back to main page'/>
            </NavLink>
            </div>
            <div>Trackers</div>
            <div className={mod.reols}>
            <NavLink to='lostor/torrents'>
                <img src={notification} alt='back to main page'/>
            </NavLink>
            </div>
            <div>Notifiers</div>
            <div className={mod.reols}>
            <NavLink to='lostor/torrents'>
                <img src={clients} alt='back to main page'/>
            </NavLink>
            </div>
            <div>Clients</div>
            <div className={mod.reols}>
            <NavLink to='lostor/torrents'>
                <img src={authentication} alt='back to main page'/>
            </NavLink>
            </div>
            <div>Authentication</div>
            <div className={mod.reols}>
            <NavLink to='lostor/torrents'>
                <img src={about} alt='back to main page'/>
            </NavLink>
            </div>
            <div>About</div>

        </div>
    )
}

export default SettingsPage;