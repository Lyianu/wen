import { Link } from "react-router-dom";

const Admin = () => {
    return ( 
        <div className="admin">
           <Link to='/write'>
                <button class='rounded-full'>Write</button>
           </Link> 
        </div>
     );
}
 
export default Admin;