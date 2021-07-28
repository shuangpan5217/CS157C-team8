import React from "react";

export class Login extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      username: "",
      password: "",
    }
  }

  render() {
    return (
      <div className="base-container" ref={this.props.containerRef}>
        <div className="header">Login</div>
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
          </div>
        </div>
        <div className="footer">
          <button type="button" className="btn"
                  onClick={() => {
                    const username = this.state.username;
                    const password = this.state.password;

                    console.log('username', username);
                    console.log('password', password);

                    const signin_api = "http://127.0.0.1:4999/login";
                    fetch(signin_api, {
                      method: 'POST',
                      body: JSON.stringify({
                        username: this.state.username,
                        password: this.state.password,
                      })
                    }).then((response) => {
                        console.log("response", response);
                        return response.json();
                    }).then((data) => {
                        console.log("############# login result############ ");
                        console.log("data", JSON.stringify(data));


                        if (data.StatusCode != 200) {
                            alert("Login failed!!");
                        }

                        else {
                            // alert("Login successfully!");
                            this.props.setIsLoggedIn(true);
                            this.props.setUser(username);
                            this.props.setNickName(data.Body.nickname);
                            this.props.setDescription(data.Body.description);


                        }

                    });
                  }}>
            Login
          </button>
        </div>
      </div>
    );
  }
}
