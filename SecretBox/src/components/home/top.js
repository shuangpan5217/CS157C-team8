import React from "react";
import "./top.css"
import iconImage from '../../image/user.jpeg';

export default class Top extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            nicknameInEdit: false,
            descriptionInEdit: false,
        }
    }

    logout() {
        this.props.setIsLoggedIn(false);
        this.props.setUser(null);
        this.props.setNickName(null);
        this.props.setDescription(null);


    }

    editNicknameClicked() {
        if (this.state.nicknameInEdit) {
            this.setState({nicknameInEdit: false});
            const api = "http://127.0.0.1:4999/updateuser";
            fetch(api, {
                method: 'PATCH',
                body: JSON.stringify({
                    username: this.props.user,
                    nickname: this.props.nickname,
                    description: this.props.description,
                })
            }).then((response) => {
                console.log("response", response);
                return response.json();
            }).then((data) => {
                console.log("############# post secret result############ ");
                console.log("data", JSON.stringify(data));

                if (data.StatusCode != 200)
                    alert("Post secret failed!!");

            });
        } else {
            this.setState({nicknameInEdit: true});
        }
    }


    editDescriptionClicked() {
        if (this.state.descriptionInEdit) {
            this.setState({descriptionInEdit: false});
            const api = "http://127.0.0.1:4999/updateuser";
            fetch(api, {
                method: 'PATCH',
                body: JSON.stringify({
                    username: this.props.user,
                    nickname: this.props.nickname,
                    description: this.props.description,
                })
            }).then((response) => {
                console.log("response", response);
                return response.json();
            }).then((data) => {
                console.log("############# post secret result############ ");
                console.log("data", JSON.stringify(data));

                if (data.StatusCode != 200)
                    alert("Post secret failed!!");

            });
        } else {
            this.setState({descriptionInEdit: true});
        }
    }


    const
    editNicknameBtn = (<span onClick={() => {
        this.editNicknameClicked()
    }}>🖊</span>);

    const
    editDescriptionBtn = (<span onClick={() => {
        this.editDescriptionClicked()
    }}>🖊</span>);

    render() {
        return (
            <div className="top">
                <div className="icon">
                    <img src={iconImage}
                    />

                </div>
                <div className="info">
                    <div className="nickname info-item">
                        {
                            this.state.nicknameInEdit ?
                                (
                                    <div>
                                        <span>Nickname: </span>
                                        <input type="text" value={this.props.nickname}
                                               onChange={
                                                   (e) => {
                                                       let val = e.target.value;
                                                       this.props.setNickName(val);
                                                   }
                                               }

                                        />
                                        {this.editNicknameBtn}
                                    </div>
                                )
                                :
                                (
                                    <div>
                                        <span>Nickname: {this.props.nickname}</span>
                                        {this.editNicknameBtn}

                                    </div>
                                )
                        }


                    </div>
                    <div className="username info-item">
                        <span> UserName: {this.props.user}</span>

                    </div>
                    <div className="description info-item">
                        {
                            this.state.descriptionInEdit ?
                                (
                                    <div>
                                        <span>Description: </span>
                                        <input type="text" value={this.props.description}
                                               onChange={
                                                   (e) => {
                                                       let val = e.target.value;
                                                       this.props.setDescription(val);
                                                   }
                                               }
                                        />
                                        {this.editDescriptionBtn}
                                    </div>
                                )
                                :
                                (
                                    <div>
                                        <span>Description: {this.props.description}</span>
                                        {this.editDescriptionBtn}

                                    </div>
                                )
                        }
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