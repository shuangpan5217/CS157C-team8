import React from "react";
import "./left.css";


export default class Left extends React.Component {
    render() {
        return(
            <div className="left">
                <div className="write">
                    write
                </div>
                <div className="read">
                    read
                </div>
                <div className="save">
                    save
                </div>
            </div>
        );
    }
}