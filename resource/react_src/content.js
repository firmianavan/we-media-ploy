class ContentV extends React.Component{
  constructor(props) {
    super(props);
    this.state = {
        value: '',
        tags: ["go","babel"]
    };
  }
  render(){
    return(
        <div className="row">
        <PageHeader>{"Example page header"}<small>Subtext for header</small></PageHeader>
        </div>
    )
  }
}