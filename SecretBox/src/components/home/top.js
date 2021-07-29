import React from "react";
import "./top.css"
import iconImage from '../../image/user.jpeg';

export default class Top extends React.Component {

    constructor(props) {
        super(props);
    }

    logout() {
        this.props.setIsLoggedIn(false);
        this.props.setUser(null);
        this.props.setNickName(null);
        this.props.setDescription(null);


    }

    render() {
        return (
            <div className="top">
                <div className="icon">
                    <img src={iconImage}
                    />

                </div>
                <div className="info">
                    <div className="nickname info-item">
                        <span>Nickname: {this.props.nickname}</span>
                        <span>🖊</span>
                    </div>
                    <div className="username info-item">
                        <span> UserName: {this.props.user}</span>

                    </div>
                    <div className="description info-item">
                        <span>Description: {this.props.description}</span>
                        <span>🖊</span>
                    </div>
                </div>
                <div className="logout">
                    <button onClick={this.logout.bind(this)}>
                        Logout
                    </button>

                </div>
            </div>
        );
    }
}