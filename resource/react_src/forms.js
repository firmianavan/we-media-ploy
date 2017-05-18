//setValue is function, when value changed, it calls setValue(currentValue)
function FieldGroup({ id, label, help,setValue }) {
  return (
    <FormGroup controlId={id}>
      <ControlLabel>{label}</ControlLabel>
      <FormControl inputRef={setValue}  />
      {help && <HelpBlock>{help}</HelpBlock>}
    </FormGroup>
  );
}

class TextareaV extends React.Component{
  constructor(props) {
    super(props);
    this.state = {
        value: '',
        tags: ["go","babel"]
    };
  }
  render(){
    return
    <textarea class="form-control" rows="3" placeholder="props.placeholder" > aaaccdvvv
    </textarea>
  }
}

class LogInFormV extends React.Component{
  constructor(props) {
    super(props);
    this.state = {
    };
    this.sub=this.sub.bind(this);
  }
  sub(){
    fetch('/a/new',{
      method: 'POST',
      headers: {
        'Content-Type': 'application/json' 
      },
      body:JSON.stringify({
        title:this.state.title.value,
        tags:this.state.tags.value,
        body:ReactDOM.findDOMNode(this.refs.articleBody).value, 
      })
    }).then(function(response) {
      return response.json()
    }).then(function(json) {
      console.log('parsed json', json)
      if (props.callback){
        props.callback()
      }
    }).catch(function(ex) {
      console.log('parsing failed', ex)
    })
  }

  render(){
    return(
        <Form horizontal>
          <FormGroup controlId="formHorizontalEmail">
            <Col componentClass={ControlLabel} sm={2}>
              Email
            </Col>
            <Col sm={10}>
              <FormControl type="email" placeholder="Email" />
            </Col>
          </FormGroup>

          <FormGroup controlId="formHorizontalPassword">
            <Col componentClass={ControlLabel} sm={2}>
              Password
            </Col>
            <Col sm={10}>
              <FormControl type="password" placeholder="Password" />
            </Col>
          </FormGroup>

          <FormGroup>
            <Col smOffset={2} sm={10}>
              <Checkbox>Remember me</Checkbox>
            </Col>
          </FormGroup>

          <FormGroup>
            <Col smOffset={2} sm={10}>
              <Button type="submit">
                Sign in
              </Button>
            </Col>
          </FormGroup>
        </Form>
    )
  }
}

class RegistFormV extends React.Component{
  constructor(props) {
    super(props);
    this.state = {
    };
    this.sub=this.sub.bind(this);
  }
  sub(){
    fetch('/a/new',{
      method: 'POST',
      headers: {
        'Content-Type': 'application/json' 
      },
      body:JSON.stringify({
        title:this.state.title.value,
        tags:this.state.tags.value,
        body:ReactDOM.findDOMNode(this.refs.articleBody).value, 
      })
    }).then(function(response) {
      return response.json()
    }).then(function(json) {
      console.log('parsed json', json)
      if (props.callback){
        props.callback()
      }
    }).catch(function(ex) {
      console.log('parsing failed', ex)
    })
  }

  render(){
    return(
        <Form horizontal>
          <FormGroup controlId="registEmail">
            <Col componentClass={ControlLabel} sm={2}>
              邮箱
            </Col>
            <Col sm={10}>
              <FormControl type="email" placeholder="Email" />
            </Col>
          </FormGroup>
          <hr className="transparent"/>
          <FormGroup controlId="registPassword">
            <Col componentClass={ControlLabel} sm={2}>
              密码
            </Col>
            <Col sm={10}> 
              <FormControl type="password" placeholder="Password" />
            </Col>
          </FormGroup>

          <FormGroup controlId="retryRegistPassword">
            <Col componentClass={ControlLabel} sm={2}>
              重新输入
            </Col>
            <Col sm={10}> 
              <FormControl type="password" placeholder="Password" />
            </Col>
          </FormGroup>

          <FormGroup>
            <Col smOffset={2} sm={10}>
              <Button className="pull-right" type="submit">
                提交
              </Button>
            </Col>
          </FormGroup>
        </Form>
    )
  }
}

class AddArticleFormV extends React.Component{
  constructor(props) {
    super(props);
    this.state = {
    };
    this.sub=this.sub.bind(this);
    this.update=this.update.bind(this);
  }
  sub(){
    fetch('/a/new',{
      method: 'POST',
      headers: {
        'Content-Type': 'application/json' 
      },
      body:JSON.stringify({
        title:this.state.title.value,
        tags:this.state.tags.value,
        body:ReactDOM.findDOMNode(this.refs.articleBody).value, 
      })
    }).then(function(response) {
      return response.json()
    }).then(function(json) {
      console.log('parsed json', json)
    }).catch(function(ex) {
      console.log('parsing failed', ex)
    })
  }
  update(key){
    return ref=>{this.state[key]=ref}    
  }

  render(){
    return(
        <form>
          <FieldGroup id="newAticleTitle" type="text" label="标题" setValue={this.update("title")} />
          <FieldGroup id="newAticleTag" type="text" label="标签" setValue={this.update("tags")} />

          <FormGroup controlId="formControlsTextarea">
            <ControlLabel>正文</ControlLabel>
            <FormControl rows="4" componentClass="textarea" placeholder="仅支持md" ref="articleBody" />
          </FormGroup>

          <Button  onClick={this.sub}>Submit</Button>
        </form>
    )
  }
};


