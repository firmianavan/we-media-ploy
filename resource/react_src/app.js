class AppV extends React.Component{
  constructor(props) {
    super(props);
    this.state = {
        value: '',
        tags: ["go","babel"]
    };
  }
  render(){
    return(
        <div>
        <NavbarV />
        <SingleColBodyV />
        </div>
    )
  }
}