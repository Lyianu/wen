const Hero = ({ title_text, image_url, is_pending, desc_text }) => {
    const url = "url(\"" + image_url + "\")"

    return ( 
        <div className="hero py-32 bg-slate-600 bg-cover bg-center relative -z-10" style={ { backgroundImage: url, height: "50%" } }>
                <div className="title justify-center flex">
                    <h1 className="text-5xl">{ is_pending ? "..." : title_text }</h1>
                </div>
                {desc_text && (
                    <div className="desc justify-start flex px-8 italic absolute bottom-10">
                        <h2 className="text-3xl">{ desc_text }</h2>
                    </div>
                )}
        </div>
     );
}
 
export default Hero;