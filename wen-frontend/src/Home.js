import Navbar from "./components/Navbar";
import Blog from "./components/Blog";
import Footer from "./components/Footer";
import Hero from "./components/Hero";
import { useState, useEffect } from "react";
import useFetch from "./useFetch";

const Home = () => {
    
    const [nav, setNav] = useState(false)
    const [title, setTitle] = useState("")
    const [imgUrl, setImgUrl] = useState("")
    const toggleNav = () => {
        //alert("clicked!")
        setNav(!nav)
    }
    const { data, isPending } = useFetch("http://localhost:8000/api/v1/site")

    useEffect(() => {
        if (data) {
            setTitle(data.name);
            setImgUrl(data.image_url); 
        }
    }, [data])



    return ( 
        <div>
            <Hero title_text={ title } image_url={ imgUrl } />
            <Blog />
            <Footer title_text={ title } />
        </div>
     );
}
 
export default Home;