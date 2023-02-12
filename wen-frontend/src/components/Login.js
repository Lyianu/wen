import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useCookies } from 'react-cookie'

const Login = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const [cookie, setCookie] = useCookies("token")

    const navigate = useNavigate();

    const handleSubmit = (e) => {
        e.preventDefault();
        const user = { username, password };

        fetch('/api/v1/user', {
            method: 'POST',
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(user)
        }).then((data) => {
            return data.json()
        }).then((result) => {
            if (result.code === 200) {
                setCookie("token", result.data.token, { path: "/" })
                navigate('/admin')
            }
        })
    }

    return ( 
        <div className="setup flex justify-center m-auto h-screen">
            <div className="form p-10 border m-auto rounded-lg">
               <h1 className="font-bold text-xl">Login</h1>
                <form onSubmit={ handleSubmit } className='flex flex-col p-1 justify-center content-around'>
                    <div className='user p-2'>                        <div className='username-input p-1'>
                            <label className='p-1'>Username</label>
                            <input 
                                type="text"
                                required
                                value={ username }
                                onChange={(e) => setUsername(e.target.value)}
                                className="border rounded-md p-1"
                            />
                        </div>
                        <div className='password-input p-1'>
                            <label className='p-1'>Password</label>
                            <input
                                type="password"
                                required
                                value={ password }
                                onChange={(e) => setPassword(e.target.value)}
                                className="border rounded-md p-1"
                            />
                        </div>
                    </div>
                    <button className='rounded-full border p-3'>Submit</button>
                </form>
            </div>
        </div>
     );
}
 
export default Login;