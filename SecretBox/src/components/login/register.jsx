import React from "react";

export class Register extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      username: "",
      password: "",
      nickname: "",
      description: "",
    }
  }

  render() {
    return (
        <div className="base-container" ref={this.props.containerRef}>
          <div className="header">Register</div>
          <div className="content">
            <div className="form">
              <div className="form-group">
                <label htmlFor="username">Username</label>
                <input type="text" name="username" placeholder="username"
                       value={this.state.username}

                       onChange={
                         (e) => {
                           console.log(this.state.username);

                           let val = e.target.value;
                           this.setState((state, props) => ({
                             username: val
                           }));
                         }
                       }
                />
              </div>
              <div className="form-group">
                <label htmlFor="password">Password</label>
                <input type="password" name="password" placeholder="password"
                       value={this.state.password}

                       onChange={
                         (e) => {
                           console.log(this.state.password);

                           let val = e.target.value;
                           this.setState((state, props) => ({
                             password: val
                           }));
                         }
                       }
                />
              </div>
              <div className="form-group">
                <label htmlFor="nickname">Nickname</label>
                <input type="text" name="nickname" placeholder="nickname"
                       value={this.state.nickname}

                       onChange={
                         (e) => {
                           console.log(this.state.nickname);

                           let val = e.target.value;
                           this.setState((state, props) => ({
                             nickname: val
                           }));
                         }
                       }
                />
              </div>
              <div className="form-group">
                <label htmlFor="description">Personal Description</label>
                <input type="text" name="description" placeholder="personal description"
                       value={this.state.description}

                       onChange={
                         (e) => {
                           console.log(this.state.description);
                           let val = e.target.value;
                           this.setState((state, props) => ({
                             description: val
                           }));
                         }
                       }
                />
              </div>
            </div>
          </div>
          <div className="footer">
            <button type="button" className="btn"
                    onClick={() => {
                      const username = this.state.username;
                      const password = this.state.password;
                      const nickname = this.state.nickname;
                      const description = this.state.description;

                      console.log('username', username);
                      console.log('password', password);
                      console.log('nickname', nickname);
                      console.log('description', description);

                      const signup_api = "http://127.0.0.1:4999/login?signup=true";
                      fetch(signup_api, {
                        method: 'POST',
                        body: JSON.stringify({
                          username: this.state.username,
                          password: this.state.password,
                          nickname: this.state.nickname,
                          description: this.state.description,
                        })
                      }).then((response) => {
                        console.log("response", response);
                        return response.json();
                      }).then((data) => {
                          console.log("############# register result############ ");
                          console.log("data", JSON.stringify(data));


                          if (data.statusCode != 201)
                              alert("Sign up failed!!");
                          else
                              alert("Sign up successfully!");
                      });
                    }}>
              Register
            </button>
          </div>
        </div>
    );
  }
}
