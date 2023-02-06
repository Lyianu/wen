import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import { NavBar } from './navbar';
import './index.css'

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <>
        <NavBar />
        <App />
    </>
)