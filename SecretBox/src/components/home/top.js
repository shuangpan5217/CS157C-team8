import React from "react";
import "./top.css"
import iconImage from '../../image/user.jpeg';

export default class Top extends React.Component {

    logout() {
        this.props.setIsLoggedIn(false);
        this.props.setUser(null);
        this.props.setNickName(null);
        this.props.setDescription(null);


    }
    render() {
        return(
            <div className="top">
                <div className="icon">
                    <img src={iconImage}
                    />

                </div>
                <div className="info">
                    <div className="nickname info-item">
                        Nickname: {this.props.nickname}
                    </div>
                    <div className="username info-item">
                        UserName: {this.props.user}
                    </div>
                    <div className="description info-item">
                         Description: {this.props.description}
                    </div>
                </div>
                <div className="logout">
                    <button onClick={this.logout}>
                        Logout
                    </button>

                </div>
            </div>
        );
    }
}