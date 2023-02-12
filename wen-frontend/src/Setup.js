import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'

const Setup = () => {
    const [name, setName] = useState("Wen");
    const [desc, setDesc] = useState("Another Wen Blog");
    const [user, setUser] = useState("admin");
    const [pass, setPass] = useState("");
    const navigate = useNavigate();

    useEffect(() => {
        fetch("http://localhost:8000/api/v1/site")
        .then((data) => data.json())
        .then((data) => {
            if(data.name !== "")
                navigate("/404")
                
        });
    }, [])

    const handleSubmit = (e) => {
        e.preventDefault();
        const settings = { name, desc, user, pass };

        fetch('http://localhost:8000/api/v1/site', {
            method: 'POST',
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(settings)
        }).then(() => {
            navigate('/');
        })
    }

    return ( 
        <div className="setup flex justify-center m-auto h-screen">
            <div className="form p-10 border m-auto rounded-lg">
               <h1 className="font-bold text-xl">Wen Setup</h1>
                <form onSubmit={ handleSubmit } className='flex flex-col p-1 justify-center content-around'>
                    <div className='website-settings p-1 border-b content-evenly'>
                        <h1 className='p-2'>Website Settings:</h1>
                        <div className='name p-1'>
                            <label className='p-1'>Website Name</label>
                            <input
                                type="text"
                                required
                                value={ name }
                                onChange={(e) => setName(e.target.value)}
                                className="border rounded-md p-1"
                            />
                        </div>
                        <div className='description p-1'>
                            <label className='p-1'>Website Description</label>
                            <input
                                type="text"
                                required
                                value={ desc }
                                onChange={(e) => setDesc(e.target.value)}
                                className="border rounded-md p-1"
                            />
                        </div>
                    </div>
                    <div className='admin_settings p-2'>
                        <h1 className='p-2'>Admin User:</h1>
                        <div className='admin_user p-1'>
                            <label className='p-1'>Username</label>
                            <input 
                                type="text"
                                required
                                value={ user }
                                onChange={(e) => setUser(e.target.value)}
                                className="border rounded-md p-1"
                            />
                        </div>
                        <div className='admin_pass p-1'>
                            <label className='p-1'>Password</label>
                            <input
                                type="password"
                                required
                                value={ pass }
                                onChange={(e) => setPass(e.target.value)}
                                className="border rounded-md p-1"
                            />
                        </div>
                    </div>
                    <button className='rounded-full border p-3'>Create</button>
                </form>
            </div>
        </div>
     );
}
 
export default Setup;