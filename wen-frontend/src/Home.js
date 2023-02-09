import Navbar from "./components/Navbar";
import BlogList from "./components/BlogList";
import Footer from "./components/Footer";
import Hero from "./components/Hero";
import { useState, useEffect } from "react";
import useFetch from "./useFetch";

const Home = ({ title, imgUrl}) => {
    
    const [nav, setNav] = useState(false)
    const toggleNav = () => {
        //alert("clicked!")
        setNav(!nav)
    }




    return ( 
        <div>
            <Hero title_text={ title } image_url={ imgUrl } desc_text="Wen" />
            <BlogList  />
        </div>
     );
}
 
export default Home;