import { useParams } from "react-router-dom";
import { useState, useEffect } from "react";
import useFetch from "./useFetch";
import Hero from "./components/Hero";
import Article from "./components/Article";

const BlogView = ( { img_url } ) => {
    const { id } = useParams();
    const [content, setContent] = useState("")
    const [title, setTitle] = useState("")
    const [subheading, setSubheading] = useState("")
    const {data, isPending} = useFetch("/api/v1/articles/" + id)

    useEffect(() => {
        if (data) {
            setTitle(data.data.title)
            setContent(data.data.content)
            setSubheading("@" + data.data.created_by)
        }
    }, [data]);


    return (
        <div className="page-view">
            <Hero title_text={title} image_url={img_url} is_pending={isPending} desc_text={subheading} />
            <Article content={ content } />
        </div>
     );
}

export default BlogView;