import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { useCookies } from "react-cookie";
import Writer from "./Writer";
import PutAuth from "../PutAuth";

const ArticleEdit = () => {
    const { id } = useParams();
    const [title, setTitle] = useState("");
    const [content, setContent] = useState("");

    const [cookie] = useCookies();
    useEffect(() => {
        if (id) 
            fetch("/api/v1/articles/" + id + "/md", {
                headers: new Headers({
                    "Authorization": "Bearer " + cookie.token
                })
            }).then((data) => data.json())
            .then((data) => {
                if (data) {
                    setContent(data.data.content);
                    setTitle(data.data.title);
                }
            })
    }, [id])

    return ( 
        <div className="writer p-24">
            <h1 className="text-2xl">Editing Article: {title}</h1>
            <Writer _title={title} _content={content} _endpoint={"/api/v1/articles/" + id} _func={PutAuth} />
        </div>
     );
}
 
export default ArticleEdit;