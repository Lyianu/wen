import Navbar from "./components/Navbar";
import Blog from "./components/Blog";
import Footer from "./components/Footer";
import Hero from "./components/Hero";
import { useState, useEffect } from "react";

const Home = () => {
    
    const [nav, setNav] = useState(false)
    const toggleNav = () => {
        //alert("clicked!")
        setNav(!nav)
    }
    const [pages, setPages] = useState(null)

    useEffect(() => {
    fetch("http://localhost:8000/api/v1/pages")
        .then((response) => {
          return response.json()
        })
        .then((data) => {
          //console.log(data.data.lists)
          setPages(data.data.lists)
        })
    }, [])

    return ( 
        <div>
            {pages && <Navbar toggle={toggleNav} pages={pages} />}
            <div className={nav ? "blur-sm z-0" : ""}>
                <Hero />
                <Blog />
                <Footer />
            </div>
        </div>
     );
}
 
export default Home;