import React from "react";
import "./App.scss";
import { Login, Register } from "./components/login/index";
import Top from "./components/home/top";
import Bottom from "./components/home/bottom";
import {BrowserRouter as Router, Link, Route, Switch} from "react-router-dom";


class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isLogginActive: true,
      isLoggedIn: false,
      user: null,
      nickname: null,
      description: null,
    };
  }

  // setIsLoggedIn(loggedIn) {
  //   this.setState({isLoggedIn: loggedIn});
  // }
  //
  // setUser(user) {
  //   this.setState({user: user});
  // }

  setIsLoggedIn(loggedIn) {
    this.setState({isLoggedIn: loggedIn});
  }

  setUser(user) {
    this.setState({user: user});
  }

  setNickName(nickname) {
    this.setState({nickname: nickname});
  }

  setDescription(description) {
    this.setState({description: description});
  }
  componentDidMount() {
    //Add .right by default
    this.rightSide.classList.add("right");
  }

  changeState() {
    const { isLogginActive } = this.state;

    if (isLogginActive) {
      this.rightSide.classList.remove("right");
      this.rightSide.classList.add("left");
    } else {
      this.rightSide.classList.remove("left");
      this.rightSide.classList.add("right");
    }
    this.setState(prevState => ({ isLogginActive: !prevState.isLogginActive }));
  }

  render() {
    const { isLogginActive } = this.state;
    const current = isLogginActive ? "Register" : "Login";
    const currentActive = isLogginActive ? "login" : "register";
    return  this.state.isLoggedIn ?
        (
            <div>
              <Top {...this.state}
                   setIsLoggedIn={this.setIsLoggedIn.bind(this)}
                   setUser={this.setUser.bind(this)}

                   setNickName={this.setNickName.bind(this)}

                   setDescription={this.setDescription.bind(this)}
              ></Top>
              <Bottom {...this.state}
                      setIsLoggedIn={this.setIsLoggedIn.bind(this)}
                      setUser={this.setUser.bind(this)}

                      setNickName={this.setNickName.bind(this)}

                      setDescription={this.setDescription.bind(this)}
              ></Bottom>
            </div>
        ) :
        (


            <Router>
              <div className="App">
                <div className="title">Secret Box</div>
                <div className="login">
                  <div className="container" ref={ref => (this.container = ref)}>
                    {isLogginActive && (
                        <Login containerRef={ref => (this.current = ref)}
                               setIsLoggedIn={this.setIsLoggedIn.bind(this)}
                               setUser={this.setUser.bind(this)}

                               setNickName={this.setNickName.bind(this)}

                               setDescription={this.setDescription.bind(this)}
                        />
                    )}
                    {!isLogginActive && (
                        <Register containerRef={ref => (this.current = ref)}
                                  setIsLoggedIn={this.setIsLoggedIn.bind(this)}
                                  setUser={this.setUser.bind(this)}

                                  setNickName={this.setNickName.bind(this)}

                                  setDescription={this.setDescription.bind(this)}
                        />
                    )}
                  </div>
                  <RightSide
                      current={current}
                currentActive={currentActive}
                containerRef={ref => (this.rightSide = ref)}
                onClick={this.changeState.bind(this)}
            />
          </div>
        </div>
      </Router>
    );
  }
}

const RightSide = props => {
  return (
    <div
      className="right-side"
      ref={props.containerRef}
      onClick={props.onClick}
    >
      <div className="inner-container">
        <div className="text">{props.current}</div>
      </div>
    </div>
  );
};

export default App;
