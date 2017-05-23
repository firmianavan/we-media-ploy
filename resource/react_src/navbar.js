/*var Navbar = ReactBootstrap.Navbar,
Nav = ReactBootstrap.Nav,
NavItem = ReactBootstrap.NavItem,
DropdownButton = ReactBootstrap.DropdownButton,
MenuItem = ReactBootstrap.MenuItem,
NavDropdown=ReactBootstrap.NavDropdown;*/
//import { Navbar } from 'react-bootstrap';


//登录modal
class LogInModal extends React.Component {  
  constructor(props) {
    super(props);
    this.state = { showModal: false };
    this.close=this.close.bind(this)
    this.doLogin=this.doLogin.bind(this)
    this.doRegister=this.doRegister.bind(this)
  }
  close() {
    this.setState({ showModal: false });
  }
  doLogin() {
    this.setState({ 
      showModal: true ,
      type: 'l'
    });
  }
  doRegister() {
    this.setState({ 
      showModal: true ,
      type: 'r'
    });
  }
  render() {

    return (
      <Nav pullRight>
      {
        !this.state.loginStatus ?
        <NavItem eventKey={1.1} href="" onClick={this.doLogin}>登录</NavItem>
        :
        <NavItem eventKey={1.2} href="#">新文章</NavItem>
      }
      {
        !this.state.loginStatus ?
        <NavItem eventKey={2} href="" onClick={this.doRegister}>注册</NavItem>
        :
        <NavDropdown eventKey={3} title={this.state.user.email} id="basic-nav-dropdown">
            <MenuItem eventKey={3.1}>Action</MenuItem>
            <MenuItem eventKey={3.2}>Another action</MenuItem>
            <MenuItem eventKey={3.3}>设置</MenuItem>
            <MenuItem divider />
            <MenuItem eventKey={3.3}>登出</MenuItem>
          </NavDropdown>
      }
        <Modal show={this.state.showModal} onHide={this.close}>
          <Modal.Header closeButton>
            <Modal.Title>{this.state.type=='l'?'Log In':'Register'}</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            {this.state.type=='l'?
              //通过onSuccess获取登录成功返回的用户信息
              <LogInFormV onSuccess={(user)=>{this.state.loginStatus=true;this.state.user=user;this.close()}}/>:
              <RegistFormV />
            }  
          </Modal.Body>
          <Modal.Footer>
            <Button onClick={this.close}>Close</Button>
          </Modal.Footer>
        </Modal>
      </Nav>
    );
  }
};


class NavbarV extends React.Component {
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
        <LogInModal />
      </Navbar.Collapse>
    </Navbar>
    )
  }
};