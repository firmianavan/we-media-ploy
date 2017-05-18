let ReactRouter = window.ReactRouterDOM;
let Route = ReactRouter.Route;
let Link = ReactRouter.Link;
let Router=ReactRouter.BrowserRouter;

class RouterV extends React.Component{
    constructor(props) {
        super(props);
    }
    render(){
        return(
          <Router >
            <Route path="/" component={SingleColBodyV}/>
          </Router>
        )
  }
}