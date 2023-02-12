import { useEffect, useState } from "react";
import { useCookies } from "react-cookie";
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";
import Writer from "./components/Writer";

const Admin = () => {
    const [cookie, setCookie] = useCookies("token");
    const [mode, setMode] = useState("write")
    const navigate = useNavigate();

    useEffect(() => {
        if (!cookie.token)
            navigate("/login")
    }, [])

    return ( 
        <div className="admin">
            <div className="top-con pt-24 pb-8 px-16 border-b justify-items-stretch">
                <button className='rounded-full border p-3 hover:bg-black hover:text-white' onClick={() => setMode("write")}>Write</button>
                <button className='rounded-full border p-3 hover:bg-black hover:text-white' onClick={() => setMode("articles")}>Articles</button>
            </div>
            { mode === "write" && <Writer /> }
        </div>
     );
}
 
export default Admin;