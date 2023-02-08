import { AiOutlineClose, AiOutlineMenu } from 'react-icons/ai'
import {  useState } from 'react'

const Navbar = ({toggle, pages}) => {
  const [nav, setNav] = useState(false)

  const handleNav = () => {
    toggle()
    setNav(!nav)
  }

  return ( 
    <div>
      {/* Navbar on the left */}
      <div className={nav ? 'z-20 h-screen px-3 fixed left-0 ease-in-out duration-500 border-r-2 bg-gray-100': 'ease-in-out duration-500 h-screen fixed left-[-100%]'}>
        <h1 className='p-3 font-bold underline underline-offset-4'>Wen</h1>

        <ul>
          {pages.map(page => (
            <li className='p-3 mx-6 border-b' key={page.id}>
              <a href={"/read/" + page.id}>{page.title}</a>
            </li>
          ))}
        </ul>
      </div>

      { /* navbar on the top */}
      { /* FIXME: after side-memu is called, if reduce the width of the screen, it can't be flipped */}
      { /* blur should be applied to parent element so the contents will be blurred too */}
      <div className={nav ? 'z-10 nav-bar blur-sm fixed top-0 right-0 left-0 ': 'nav-bar fixed top-0  left-0 right-0'}>

        <div className="p-3 flex justify-between">
          <h1 className={nav ? "opaque" : "font-bold underline underline-offset-4"}>Wen</h1>

          <ul className="flex justify-between invisible md:visible uppercase">
            {pages.map(page => (
              <li className='px-3 hover:underline'><a href={"/read/" + page.id}>{page.title}</a></li>
            ))}
          </ul>

        <div onClick={handleNav} className="block md:hidden">
          {nav ? <AiOutlineClose /> : <AiOutlineMenu />}
        </div>
      </div>
    </div>

    </div>
   );
}
 
export default Navbar;