const Footer = ( {title_text} ) => {
    let date = new Date().getFullYear();

    return ( 
        <div className="footer sticky top-[100vh] -z-10">
            <div className="flex flex-col items-center p-3 bg-[#000030] bottom-0 justify-center">
                <div className="black text-lg">
                    <p>&copy;{date} {title_text}</p>
                    <p>Powered by Wen</p>
                </div>
            </div>
        </div>
     );
}
 
export default Footer;