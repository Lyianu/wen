import { Link } from 'react-router-dom'

const Footer = ( {title_text, absolute} ) => {

  return (
    <>
      <footer
        className="footer sticky top-[100vh] -z-10 bg-gray-900"
      >
        <div className="container mx-auto px-4">
          <hr className="mb-6 border-b-1 border-gray-700" />
          <div className="flex flex-wrap items-center md:justify-between justify-center">
            <div className="w-full md:w-4/12 px-4">
              <div className="text-sm text-white font-semibold py-1">
                Copyright © {new Date().getFullYear()}{" "}
                <Link
                  to="/"
                  className="text-white hover:text-gray-400 text-sm font-semibold py-1"
                >
                  {title_text}
                </Link>
              </div>
            </div>
            <div className="w-full md:w-8/12 px-4">
              <ul className="flex flex-wrap list-none md:justify-end  justify-center">
                <li>
                  <a
                    href="https://github.com/Lyianu/wen"
                    className="text-white hover:text-gray-400 text-sm font-semibold block py-1 px-3"
                  >
                    Github
                  </a>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </footer>
    </>
  );
}

 
export default Footer;