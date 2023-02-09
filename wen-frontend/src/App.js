import Home from './Home'
import Footer from './components/Footer'
import React, { useState, useEffect } from 'react'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import PageView from './PageView'
import Navbar from './components/Navbar'
import useFetch from './useFetch'

export default function App() {
    const [nav, setNav] = useState(false)
    const [name, setName] = useState("")
    const [title, setTitle] = useState("")
    const [imgUrl, setImgUrl] = useState("")
    const toggleNav = () => {
        setNav(!nav)
    }


    const { data, isPending } = useFetch("http://localhost:8000/api/v1/site")

    useEffect(() => {
        if (data) {
            setTitle(data.bg_title);
            setImgUrl(data.image_url);
            setName(data.name) 
        }
    }, [data])


    return (
       <Router>
        <div className='wen flex flex-col h-screen'>
            <Navbar toggle={toggleNav} site_name={ name } nav={nav} />
            <div className={nav ? "page blur-sm z-0 flex-grow" : "page flex-grow"} onClick={() => setNav(false)}>
                <Routes>
                    <Route exact path="/" element={<Home title={title} imgUrl={imgUrl} />} />
                    <Route path="/page/:id" element={<PageView img_url={imgUrl} />} />
                </Routes>
                <Footer title_text={ name } />
            </div>
            </div>
       </Router>
    )
}