import React from 'react'

export class ContentV extends React.Component{
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
        <PageHeader>{"新文章"}<small></small></PageHeader>
        <AddArticleFormV/>
        </div>
    )
  }
}

export class SingleColBodyV extends React.Component{
  constructor(props) {
    super(props);
  }
  render(){
    return(
        <div className="container bs-docs-single-col-container">
           {this.props.children}
        </div>
    )
  }
}

export class FooterV extends React.Component{
  constructor(props) {
    super(props);
    this.state = {
        value: '',
        tags: ["go","babel"]
    };
  }
  render(){
    return(
        <footer className="bs-docs-footer" >
          <div className="container" >
              <p >
                本站内容遵守<a href="http://creativecommons.org/licenses/by-sa/3.0/" target="_blank">CC BY-SA 3.0 License</a>协议
              </p>
              <ul className="bs-docs-footer-links muted" >
                  <li ></li>
                  <li >·</li>
                  <li ><a href="https://github.com/firmianavan/we-media-ploy">GitHub</a>
                  </li>
                  <li >·</li>
                  <li ><a href="https://github.com/firmianavan/we-media-ploy/issues?state=open">Issues</a></li>
                  <li >·</li>
              </ul>
          </div>
      </footer>
    )
  }
}

