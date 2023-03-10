import Home from './Home'
import Footer from './components/Footer'
import React, { useState, useEffect } from 'react'
import { BrowserRouter as Router, Navigate, Route, Routes } from 'react-router-dom'
import PageView from './PageView'
import BlogView from './BlogView'
import Navbar from './components/Navbar'
import useFetch from './useFetch'
import Setup from './Setup'
import Login from './components/Login'
import NotFound from './components/NotFound'
import Admin from './Admin'
import ArticleEdit from './components/ArticleEdit'
import PageEdit from './components/PageEdit'

export default function App() {
    const [nav, setNav] = useState(false)
    const [name, setName] = useState("")
    const [title, setTitle] = useState("")
    const [imgUrl, setImgUrl] = useState("")
    const [desc, setDesc] = useState("")
    const toggleNav = () => {
        setNav(!nav)
    }


    const { data, isPending } = useFetch("/api/v1/site")

    useEffect(() => {
        if (data) {
            setTitle(data.bg_title);
            setImgUrl(data.image_url);
            setName(data.name);
            setDesc(data.desc)
        }
    }, [data])


    return (
       <Router>
        <div className='wen flex flex-col h-screen'>
            <Navbar toggle={toggleNav} site_name={ name } nav={nav} />
            <div className={nav ? "page blur-sm z-0 flex-grow" : "page flex-grow"} onClick={() => setNav(false)}>
                <Routes>
                    <Route exact path="/" element={<Home title={title} imgUrl={imgUrl} desc_text={desc} />} />
                    <Route path="/page/:id" element={<PageView img_url={imgUrl} />} />
                    <Route path="/read/:id" element={<BlogView img_url={imgUrl} />} />
                    <Route path="/setup" element={<Setup />} />
                    <Route path="/login" element={<Login />} />
                    <Route path="/admin" element={<Admin />} />
                    <Route path="/edit/:id" element={<ArticleEdit />} />
                    <Route path="/editpage/:id" element={<PageEdit />} />

                    <Route path='*' element={ <NotFound /> } />
                </Routes>
                <Footer title_text={ name } />
            </div>
            </div>
       </Router>
    )
}