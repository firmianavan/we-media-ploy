class SingleColBodyV extends React.Component{
  constructor(props) {
    super(props);
    this.state = {
        value: '',
        tags: ["go"]
    };
  }
  render(){
    return(
        <div className="container bs-docs-single-col-container">
            <ContentV />
        </div>
    )
  }
}