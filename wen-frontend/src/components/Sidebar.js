import { Link } from "react-router-dom";

const Sidebar = ({ list }) => {
    return ( 
        <div className="absolute right-0 sidebar p-10 border h-max">
            {list && (
                list.map((tag) => (
                    <div className="tag" key={tag.id}>
                       <Link to={`/tag/${tag.id}`}>{tag.name}</Link> 
                    </div>
                ))
            )}

        </div>
     );
}
 
export default Sidebar;