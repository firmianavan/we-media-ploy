import React from 'react'
import  {NavbarV} from './navbar.jsx'
import  {SingleColBodyV,FooterV} from './content.jsx'

import {
  RegistFormV,
  LogInFormV
} from './forms.jsx'

 
import {
  BrowserRouter as Router,
  Route,
  Link
} from 'react-router-dom'



export class SignV extends React.Component{
  constructor(props) {
    super(props);
  }
  render(){
    return(
        
        <Router >
          <div className="container sign bs-docs-single-col-container">
             <Route exact path={this.props.match.url  } component={LogInFormV}/>
             <Route path={this.props.match.url+"/in"  } component={LogInFormV}/>
             <Route path={this.props.match.url+"/up"  } component={RegistFormV}/>
          </div>
        </Router>
    )
  }
}

export const  RootRouterV =()=> {


    return(
  
          <Router >
            <div>
           
            <Route path="/sign" component={SignV}/>
            <Route exact path="/" component={AppV}/>
            </div>
          </Router>
        )
    }



export class AppV extends React.Component{
  constructor(props) {
    super(props);
    this.state={
      login:{}
    }
  }
  componentDidMount(){
    LogInFormV.AskForLogin(null,null,
      (json) =>{
      if(json.status=="success"){
        json.loginStatus=true
        this.setState({
          login:json
        })
      }
    })
  }
  render(){
    return(
        <div>
        <NavbarV status={this.state.login.status} data={this.state.login.data}/>
        <FooterV/>
        </div>
    )
  }
}

