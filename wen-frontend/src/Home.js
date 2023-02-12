import BlogList from "./components/BlogList";
import Hero from "./components/Hero";
import { useState, useEffect } from "react";
import Pagination from "./components/Pagination";
import useFetch from "./useFetch";

const Home = ({ title, imgUrl, desc_text}) => {
    const [page, setPage] = useState(1)
    const [total, setTotal] = useState(0)
    
    const [nav, setNav] = useState(false)
    const toggleNav = () => {
        //alert("clicked!")
        setNav(!nav)
    }

    const {data} = useFetch("http://localhost:8000/api/v1/articles")

    useEffect(() => {
        if (data) {
            setTotal(data.data.total)
        }
    }, [data])

    const changePage = (offset) => {
        let targetPage = page + offset;
        if (targetPage < 1)
            targetPage = 1;
        if (targetPage > Math.trunc((total - 1) / 10 + 1))
            targetPage = Math.trunc((total - 1) / 10 + 1);
        setPage(targetPage)
    }




    return ( 
        <div>
            <Hero title_text={ title } image_url={ imgUrl } desc_text={ desc_text } />
            <BlogList page_num={page}  />
            <Pagination change={changePage} page_num={page} total={total} />
        </div>
     );
}
 
export default Home;