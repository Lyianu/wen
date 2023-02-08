import Home from './Home'
import React, { useState } from 'react'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import PageView from './PageView'
import Navbar from './components/Navbar'

export default function App() {
    const [nav, setNav] = useState(false)
    const toggleNav = () => {
        setNav(!nav)
    }

    return (
       <Router>
            <Navbar toggle={toggleNav} />
            <div className={nav ? "page blur-sm z-0" : "page"}>
                <Routes>
                    <Route exact path="/" element={<Home />} />
                    <Route path="/page/:id" element={<PageView />} />
                </Routes>
            </div>
       </Router>
    )
}