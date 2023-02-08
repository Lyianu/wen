const Hero = ({ title_text, image_url }) => {
    const url = "url(\"" + image_url + "\")"

    return ( 
        <div className="hero py-48 bg-slate-600" style={ { backgroundImage: url } }>
                <div className="title justify-center align-middle flex">
                    <h1 className="text-3xl">{ title_text }</h1>
                </div>
        </div>
     );
}
 
export default Hero;