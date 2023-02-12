import { useState, useEffect } from "react";
import { useCookies } from "react-cookie";
import { useNavigate } from "react-router-dom";
import PostAuth from "../PostAuth";

const Writer = ( {_title, _content, _endpoint, _func }) => {
    const [title, setTitle] = useState(_title ? _title : "");
    const [content, setContent] = useState(_content ? _content : "");

    const [cookie] = useCookies("token")

    const navigate = useNavigate()

    useEffect(() => {
        setTitle(_title);
        setContent(_content);
    }, [_title, _content]);

    if (!_endpoint) {
        _endpoint = "http://localhost:8000/api/v1/articles";
    }

    if (!_func) {
        _func = PostAuth;
    }

    const handleSubmit = (e) => {
        e.preventDefault()     
            
        _func(_endpoint, {
            "title": title,
            "content": content
        }, cookie) 

        navigate("/")
    }


    return ( 
        <div className="writer">
            <form className="m-auto flex flex-col h-screen" onSubmit={ handleSubmit }>
                <input
                    className="rounded-sm border border-black h-[5%] mx-10 my-3"
                    type="text"
                    value={ title }
                    onChange={(e) => setTitle(e.target.value)}
                    required
                />
                <textarea
                    className="rounded-sm border border-black h-[80%] mx-10 my-3 p-2"
                    type="text"
                    value={ content }
                    onChange={(e) => setContent(e.target.value)}
                />
                <button className="rounded-full border p-3 m-auto px-10 hover:bg-black hover:text-white">Submit</button>
            </form>
        </div>
     );
}
 
export default Writer;