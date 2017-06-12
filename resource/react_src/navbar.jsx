/*var Navbar = ReactBootstrap.Navbar,
Nav = ReactBootstrap.Nav,
NavItem = ReactBootstrap.NavItem,
DropdownButton = ReactBootstrap.DropdownButton,
MenuItem = ReactBootstrap.MenuItem,
NavDropdown=ReactBootstrap.NavDropdown;*/
import  React  from 'react'
import { 
  Navbar,
  Nav,
  NavItem,
  DropdownButton,
  MenuItem,
  NavDropdown,
  Modal,
  Button
  } from 'react-bootstrap';

import {
  RegistFormV,
  LogInFormV
} from './forms.jsx'

function linkto(link){
  return () => {
    location.href=link
  }
}

//登录modal
export class UserItem extends React.Component {  
  constructor(props) {
    super(props);
  }
  render() {

    return (
      <Nav pullRight>
      {
        !this.props.status ?
        <NavItem eventKey={1.1} href="#" onClick={linkto("/sign/in")} >登录</NavItem>
        :
        <NavItem eventKey={1.2} href="#">新文章</NavItem>
      }
      {
        !this.props.status ?
        <NavItem eventKey={2} href="#" onClick={linkto("/sign/up")} >注册</NavItem>
        :
        <NavDropdown eventKey={3} title={this.props.data.email} id="basic-nav-dropdown">
            <MenuItem eventKey={3.1}>Action</MenuItem>
            <MenuItem eventKey={3.2}>Another action</MenuItem>
            <MenuItem eventKey={3.3}>设置</MenuItem>
            <MenuItem divider />
            <MenuItem eventKey={3.3}>登出</MenuItem>
          </NavDropdown>
      }
      </Nav>
    );
  }
};


export class NavbarV extends React.Component {
  constructor(props) {
    super(props);
    this.state = {value: ''};
  }
  render(){
    return (
      <Navbar inverse collapseOnSelect>
      <Navbar.Header>
        <Navbar.Brand>
          <a href="#">马颊河</a>
        </Navbar.Brand>
        <Navbar.Toggle />
      </Navbar.Header>
      <Navbar.Collapse>
        <UserItem status={this.props.status} data={this.props.data}/>
      </Navbar.Collapse>
    </Navbar>
    )
  }
};