import { useEffect } from "react";
import { useCookies } from "react-cookie";
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";

const Admin = () => {
    const [cookie, setCookie] = useCookies("token");
    const navigate = useNavigate();

    useEffect(() => {
        if (!cookie.token)
            navigate("/login")
    }, [])

    return ( 
        <div className="admin">
            <div className="top-con">
                <Link to='/write'>
                        <button class='rounded-full'>Write</button>
                </Link> 
            </div>
        </div>
     );
}
 
export default Admin;