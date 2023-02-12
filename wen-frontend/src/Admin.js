import { useEffect, useState } from "react";
import { useCookies } from "react-cookie";
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";
import useFetch from "./useFetch";

import Pagination from "./components/Pagination";
import SiteSetting from "./components/SiteSetting";
import Writer from "./components/Writer";
import BlogEditList from "./components/BlogEditList";

const Admin = () => {
    const [page, setPage] = useState(1)
    const [total, setTotal] = useState(0)

    const [cookie, setCookie] = useCookies("token");
    const [mode, setMode] = useState("write")
    const navigate = useNavigate();

    useEffect(() => {
        if (!cookie.token)
            navigate("/login")
    }, [])

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
        if (targetPage > Math.trunc((total - 1) / 10) + 1)
            targetPage = Math.trunc((total - 1) / 10) + 1;
        setPage(targetPage)
    }

    return ( 
        <div className="admin">
            <div className="top-con pt-24 pb-8 px-16 border-b justify-items-stretch">
                <button className='rounded-full border p-3 hover:bg-black hover:text-white' onClick={() => setMode("write")}>Write</button>
                <button className='rounded-full border p-3 hover:bg-black hover:text-white' onClick={() => setMode("articles")}>Articles</button>
                <button className='rounded-full border p-3 hover:bg-black hover:text-white' onClick={() => setMode("site")}>Site</button>
            </div>
            { mode === "write" && (
                <>
                    <h1 className="text-2xl px-10 pt-3">Add a post</h1>
                    <Writer />
                </>
            )}

            { mode == "articles" && (
                <>
                    <h1 className="text-2xl px-10 pt-3">Manage Articles</h1>
                    <BlogEditList page_num={page} />
                    {total > 10 && <Pagination total={total} page_num={page} change={changePage} />}
                </>
            )}

            { mode === "site" && (
                <>
                    <h1 className="text-2xl px-10 pt-3">Site settings</h1>
                    <SiteSetting />
                </>
            )}
        </div>
     );
}
 
export default Admin;