import useFetch from "../useFetch";
import { useState, useEffect } from "react";

const BlogList = ( { page_num } ) => {
    if (!page_num)
        page_num = 1;

    const [blogs, setBlogs] = useState(null)
    const [is_pending, setIs_pending] = useState(true)
    const {data, isPending} = useFetch("http://127.0.0.1:8000/api/v1/articles")

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
                        blogs.map(blog => (
                            <div className="blog-thumbnail mx-16 my-5 border-b pb-3" key={blog.id}>
                                <h2 className="blog text-xl">{blog.title}</h2>
                                <p className="desc text-base text-slate-500">{blog.desc === "" ? "This article has no description" : blog.desc}</p>
                            </div>
                        ))
                    )}
                </>
            )}
        </div>
     );
}
 
export default BlogList;