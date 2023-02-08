import React from "react";
import { ReactDOM } from "react";

export class Pagelink extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <div>
                <a href={this.props.url}>{this.props.title}</a>
            </div>
        )
    }
}