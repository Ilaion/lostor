import React from "react";

import mod from './LogoutPage.module.css';

const LogoutPage = () => {
    return(
        <div className={mod.loginform}>
            <div className={mod.header}>
                <div>
                    Enter Lostor
                </div>
            </div>
            <div>
                <input placeholder='Password'/>
            </div>
            <button className={mod.button}>ENTER</button>
        </div>
    )
}

export default LogoutPage;