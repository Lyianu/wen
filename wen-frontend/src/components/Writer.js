import { useState } from "react";
import { useCookies } from "react-cookie";
import PostAuth from "../PostAuth";

const Writer = () => {
    const [title, setTitle] = useState("");
    const [content, setContent] = useState("");

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
            <form onSubmit={ handleSubmit }>
                <input
                    type="text"
                    value={ title }
                    onChange={(e) => setTitle(e.target.value)}
                    required
                />
                <input 
                    type="text"
                    value={ content }
                    onChange={(e) => setContent(e.target.value)}
                />
                <button className="rounded-full">Submit</button>
            </form>
        </div>
     );
}
 
export default Writer;