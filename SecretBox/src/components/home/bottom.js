import React from "react";
import "./bottom.css"
import Left from "./left";
import Center from "./center";
import writeIcon from '../../image/write.jpeg';
import readIcon from '../../image/read.png';
import saveIcon from '../../image/save.png';

export default class Bottom extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            msgPanelState: "",
            msgInEdit: "",
            msgPicked: "",
            secret_id: "",
            secret_owner: "",
            created_time: "",
            secret_nickname: "",
            savedMsgs: [],
        };
    }

    writeMsg() {
        this.setState({msgPanelState: 'write'});

    }

    pickMsg() {

        const api = "http://127.0.0.1:4999/secret?username=" + this.props.user;
        fetch(api, {
            method: 'GET',
        }).then((response) => {
            console.log("response", response);
            return response.json();
        }).then((data) => {
            console.log("############# post secret result############ ");
            console.log("data", JSON.stringify(data));


            this.setState({
                msgPicked: data.Body.content,
                secret_id: data.Body.secret_id,
                secret_owner: data.Body.username,
                secret_nickname: data.Body.nickname,
                created_time: new Date(data.Body.created_time).toISOString(),


            });


        }).catch(err => {
            console.log("picking secret error:", err);
            this.setState({msgPicked: ""});

        });


        this.setState({msgPanelState: 'picked'});

    }

    savedMsg() {

        const api = "http://127.0.0.1:4999/savedsecret?username=" + this.props.user;
        fetch(api, {
            method: 'GET',
        }).then((response) => {
            console.log("response", response);
            return response.json();
        }).then((data) => {
            console.log("############# get saved secrets result############ ");
            console.log("data", JSON.stringify(data));


            if (data.StatusCode != 200)
                alert("Get saved secrets failed!!");
            else {
                this.setState({savedMsgs: data.Body.saved_secrets})
            }

        });


        this.setState({msgPanelState: 'saved'});

    }

    getMsgPanelMap() {
        return {
            "": (<div className={"empty"}>Click a button on the left to perform an action</div>),
            write: (
                <div className={"write"}>
                    <textarea value={this.state.msgInEdit} onChange={(e) => {
                        this.setState({msgInEdit: e.target.value})
                    }}></textarea>
                    <button onClick={() => {
                        if (this.state.msgInEdit.length === 0) {
                            alert("Can not post empty secret!");
                            return;
                        }

                        const api = "http://127.0.0.1:4999/secret";
                        fetch(api, {
                            method: 'POST',
                            body: JSON.stringify({
                                username: this.props.user,
                                nickname: this.props.nickname,
                                content: this.state.msgInEdit,
                            })
                        }).then((response) => {
                            console.log("response", response);
                            return response.json();
                        }).then((data) => {
                            console.log("############# post secret result############ ");
                            console.log("data", JSON.stringify(data));


                            if (data.StatusCode != 201)
                                alert("Post secret failed!!");
                            else
                                alert("Post secret successfully!");
                        });
                    }}>Post
                    </button>
                </div>
            ),
            picked: (
                <div className={
                    "picked"
                }
                >
                    <span>{this.state.msgPicked || "No more secrets, please try again later."}</span>
                    {this.state.msgPicked ? <button onClick={() => {


                        const api = "http://127.0.0.1:4999/savedsecret";
                        fetch(api, {
                            method: 'POST',
                            body: JSON.stringify({
                                "secret_id": this.state.secret_id,
                                "secret_owner": this.state.secret_owner,
                                "nickname": this.state.secret_nickname,
                                "username": this.props.user,
                                // front-end will take care of this field because it is not the same with the response
                                // need some format conversion
                                "created_time": this.state.created_time,
                                "content": this.state.msgPicked
                            })
                        }).then((response) => {
                            console.log("response", response);
                            return response.json();
                        }).then((data) => {
                            console.log("############# save secret result############ ");
                            console.log("data", JSON.stringify(data));


                            if (data.StatusCode != 200)
                                alert("Saving secret failed!!");
                            else
                                alert("Saving secret successfully!");
                        });
                    }}
                    >Save</button> : null}

                </div>
            ),
            saved: (
                <div className={
                    "saved"
                }
                >
                    <ul>
                        {
                            this.state.savedMsgs.map(msgObj => (
                                <li>
                                    <span>{msgObj.nickname} : {msgObj.content} </span>
                                    <span
                                        onClick={() => {
                                            let newSavedMsgs = [];
                                            // newSavedMsgs = newSavedMsgs.concat(this.state.savedMsgs) ;
                                            for (let i = 0; i < this.state.savedMsgs.length; ++i) {
                                                if (this.state.savedMsgs[i].secret_id != msgObj.secret_id)
                                                    newSavedMsgs.push(this.state.savedMsgs[i]);
                                            }

                                            this.setState({savedMsgs: newSavedMsgs});

                                            const api = "http://127.0.0.1:4999/savedsecret?username=" + this.props.user
                                                + "&&secret_id=" + msgObj.secret_id;
                                            fetch(api, {
                                                method: 'DELETE',
                                            }).then((response) => {
                                                console.log("response", response);
                                                return response.json();
                                            }).then((data) => {
                                                if (data.StatusCode != 200)
                                                    alert("Delete failed!!");
                                            });
                                        }}
                                    >🗑️</span>
                                </li>

                            ))
                        }


                    </ul>
                </div>
            ),


        }
            ;
    }

    render() {
        return (
            <div className="bottom">
                <div className="left bottom-item">
                    <div className="write">
                        <button>
                            <img src={writeIcon} onClick={this.writeMsg.bind(this)}></img>

                        </button>
                    </div>
                    <div className="read">
                        <button>
                            <img src={readIcon} onClick={this.pickMsg.bind(this)}></img>
                        </button>
                    </div>
                    <div className="save">
                        <button>
                            <img src={saveIcon} onClick={this.savedMsg.bind(this)}></img>
                        </button>
                    </div>
                </div>
                <div className="center bottom-item">
                    {
                        this.getMsgPanelMap()[this.state.msgPanelState]
                    }


                </div>
            </div>


        );
    }
}