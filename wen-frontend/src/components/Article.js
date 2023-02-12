import Parser from 'html-react-parser'

const Article = ( { content } ) => {
    
    return ( 
        <div className="article">
            <div className="p-3 m-8 md:m-16">
                { Parser(content) }
            </div>
        </div>
     );
}
 
export default Article;