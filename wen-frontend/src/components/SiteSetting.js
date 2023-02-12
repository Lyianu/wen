import { useEffect, useState } from "react";
import useFetch from "../useFetch";

const SiteSetting = () => {
    const [image_url, setImage_url] = useState("")
    const [name, setName] = useState("")
    const [bg_title, setBg_title] = useState("")
    const [desc, setDesc] = useState("")

    const handleSubmit = (e) => {
        e.preventDefault();

        

        console.log(image_url, name, bg_title, desc)
    }

    const {data, isPending} = useFetch("http://localhost:8000/api/v1/site")
    useEffect(() => {
        if (data) {
            setBg_title(data.bg_title);
            setImage_url(data.image_url);
            setName(data.name);
            setDesc(data.desc);
        }
    }, [data])

    return ( 
        <div className="site-setting">
            <form className="flex flex-col h-screen" onSubmit={ handleSubmit }>
                <div className="site-name mx-auto mt-auto p-3">
                    <label>Name</label>
                    <input 
                        className="border m-1 p-1"
                        type="text"
                        required
                        value={ name }
                        onChange={(e) => setName(e.target.value)}
                    />
                </div>
                <div className="desc mx-auto p-3">
                    <label>Description</label>
                    <input
                        className="border m-1 p-1"
                        type="text"
                        value={ desc }
                        onChange={(e) => setDesc(e.target.value)}
                    />
                </div>
                <div className="bg-title mx-auto p-3">
                    <label>Background Title</label>
                    <input
                        className="border m-1 p-1"
                        type="text"
                        value={ bg_title }
                        onChange={(e) => setBg_title(bg_title)}
                    />
                </div>
                <div className="bg-image mx-auto p-3">
                    <label>Background Image URL</label>
                    <input
                        className="border m-1 p-1"
                        type="text"
                        value={ image_url }
                        onChange={(e) => setImage_url(image_url)}
                    />
                </div>

                <button className="rounded-full border p-3 m-auto px-10 hover:bg-black hover:text-white">Submit</button>
            </form>
        </div>
     );
}
 
export default SiteSetting;