import useFetch from "../useFetch";
import { useState, useEffect } from "react";
import { Link } from "react-router-dom";

const BlogEditList = ( { page_num } ) => {
    if (!page_num)
        page_num = 1;

    const [blogs, setBlogs] = useState(null)
    const [is_pending, setIs_pending] = useState(true)
    const {data, isPending} = useFetch("http://127.0.0.1:8000/api/v1/articles?page=" + page_num)

    useEffect(() => {
        if (data) {
            setBlogs(data.data.lists)
            setIs_pending(isPending)
        }
    }, [data])

    return ( 
        <div className="-z-10 mx-3 blog">
            {!is_pending && blogs && (
                <>
                    { is_pending ? "Loading" : (
                        blogs.map((blog, index) => (
                            <div className={index + 1 === blogs.length ? "blog-thumbnail mx-16 my-5 pb-3 flex flex-row justify-around" : "blog-thumbnail mx-16 my-5 border-b pb-3 flex flex-row justify-around"} key={blog.id}>
                                <div classname="article-preview">
                                    <div className="preview">
                                        <Link to={"/read/" + blog.id}><h2 className="blog text-xl">{blog.title}</h2></Link>
                                        <p className="desc text-base text-slate-500">{blog.desc === "" ? "This article has no description" : blog.desc}</p>
                                    </div>
                                </div>
                                <div className="article-options">
                                    <button className="rounded-full border py-3 px-7 mx-1">EDIT</button>
                                    <button className="rounded-full border py-3 px-7 mx-1">DELETE</button>
                                </div>
                            </div>
                        ))
                    )}
                </>
            )}
        </div>
     );
}
 
export default BlogEditList;