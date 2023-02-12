import { useEffect } from "react";
import { useCookies } from "react-cookie";
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";
import Writer from "./components/Writer";

const Admin = () => {
    const [cookie, setCookie] = useCookies("token");
    const navigate = useNavigate();

    useEffect(() => {
        if (!cookie.token)
            navigate("/login")
    }, [])

    return ( 
        <div className="admin">
            <div className="top-con pt-24 pb-8 px-16 border-b">
                <Link to='/write'>
                        <button className='rounded-full'>Write</button>
                </Link> 
            </div>
            <Writer />
        </div>
     );
}
 
export default Admin;