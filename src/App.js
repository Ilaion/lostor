import React from "react";
import './App.css';
import Header from "./components/Header/Header";
import {BrowserRouter, Routes, Route} from "react-router-dom";
import MainPage from "./components/Content/MainPage/MainPage";
import SettingsPage from "./components/Content/SettingsPage/SettingsPage";
import LogoutPage from "./components/Content/LogoutPage/LogoutPage";


const App = (props) => {
    return (
        <BrowserRouter>
            <div className="App">
                <Header/>
                <Routes>
                    <Route path='/lostor/login' element={<LogoutPage/>}/>
                    <Route path='/lostor/torrents' element={<MainPage/>}/>
                    <Route path='/lostor/settings' element={<SettingsPage/>}/>
                </Routes>
            </div>
        </BrowserRouter>
    );
}

export default App;
