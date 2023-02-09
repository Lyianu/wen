import Hero from "./components/Hero";
import useFetch from "./useFetch";
import { useParams } from "react-router-dom";
import { useState, useEffect } from "react";
import Article from "./components/Article";

const PageView = ( {img_url} ) => {
    const { id } = useParams();
    const [content, setContent] = useState("")
    const [title, setTitle] = useState("")
    const {data, isPending} = useFetch("http://127.0.0.1:8000/api/v1/pages/" + id)

    useEffect(() => {
        if (data) {
            setTitle(data.data.title)
            setContent(data.data.content)
        }
    }, [data]);


    return (
        <div className="page-view">
            <Hero title_text={title} image_url={img_url} is_pending={isPending} />
            <Article content={ content } />
        </div>
     );
}
 
export default PageView;