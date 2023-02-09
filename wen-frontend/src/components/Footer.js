const Footer = ( {title_text} ) => {
    let date = new Date().getFullYear();

    return ( 
        <div className="footer sticky top-[100vh]">
            <div className="flex flex-col items-center p-3 bg-[#000030] bottom-0 justify-center">
                <div style={{color: "black", textShadow: "0px 0px 3px white"}}>
                    <p>&copy;{date} {title_text}</p>
                    <p>Powered by Wen</p>
                </div>
            </div>
        </div>
     );
}
 
export default Footer;