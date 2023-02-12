import { useState } from "react";
import { useCookies } from "react-cookie";
import PostAuth from "../PostAuth";

const Writer = ( {_title, _content }) => {
    const [title, setTitle] = useState(_title ? _title : "");
    const [content, setContent] = useState(_content ? _content : "");

    const [cookie] = useCookies("token")

    const handleSubmit = (e) => {
        e.preventDefault()     
            
        PostAuth("http://localhost:8000/api/v1/articles", {
            "title": title,
            "content": content
        }, cookie) 
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
                <input 
                    className="rounded-sm border border-black h-[80%] mx-10 my-3"
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