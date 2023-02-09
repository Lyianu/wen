const Hero = ({ title_text, image_url, is_pending }) => {
    const url = "url(\"" + image_url + "\")"

    return ( 
        <div className="hero py-32 bg-slate-600 bg-cover bg-center" style={ { backgroundImage: url, height: "50%" } }>
                <div className="title justify-center align-middle flex">
                    <h1 className="text-5xl">{ is_pending ? "..." : title_text }</h1>
                </div>
        </div>
     );
}
 
export default Hero;