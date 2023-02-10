import { AiOutlineClose, AiOutlineMenu } from 'react-icons/ai'
import {  useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import useFetch from '../useFetch'

const Navbar = ({toggle, site_name, nav}) => {
  const [pages, setPages] = useState(null)

  const handleNav = () => {
    toggle()
    //setNav(!nav)
  }


  const {data, isPending} = useFetch("http://localhost:8000/api/v1/pages")
  
  useEffect(() => {
    setPages(data && data.data.lists)
  }, [data])

  //console.log(pages)
  return ( 
    <div className='navbar text-xl'>
      {/* Navbar on the left */}
      <div className={nav ? 'z-20 h-screen px-3 fixed left-0 ease-in-out duration-500 border-r-2 bg-gray-100': 'ease-in-out duration-500 h-screen fixed left-[-100%]'}>
        <Link to="/"><h1 className='p-3 font-bold underline underline-offset-4'>
          { site_name }
          </h1></Link>
        <ul>
          { !isPending && pages && (pages.map((page, index) => (
            <li className={ index + 1 === pages.length ? 'p-3 mx-6' : 'p-3 mx-6 border-b'} key={page.id}>
              <Link to={"/page/" + page.id}>{page.title}</Link>
            </li>
          )))
        }
        </ul>
      </div>

      { /* navbar on the top */}
      { /* blur should be applied to parent element so the contents will be blurred too */}
      <div className={nav ? 'z-10 nav-bar blur-sm fixed top-0 right-0 left-0': 'nav-bar fixed top-0 left-0 right-0'}>

        <div className="p-8 flex justify-between">
          <Link to="/">
            <div>
            <h1 className={nav ? "opaque" : "font-bold underline underline-offset-4"}>{ site_name }</h1>
            </div>
          </Link>

          { !isPending && pages && (
              <ul className="flex justify-between invisible md:visible uppercase">
                {pages.map(page => (
                  <li className='px-3 hover:underline' key={page.id}><Link to={"/page/" + page.id}>{page.title}</Link></li>
                ))}
              </ul>
            )

            

          }

        <div onClick={handleNav} className="block md:hidden">
          {nav ? <AiOutlineClose /> : <AiOutlineMenu />}
        </div>
      </div>
    </div>

    </div>
   );
}
 
export default Navbar;