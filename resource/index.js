import React from 'react'
import {render} from 'react-dom'
import  {RootRouterV} from './react_src/app.jsx'

var __fetch =fetch;
fetch = function(url,option,on_json,on_error){
    __fetch(url,option)
        .then(function(response){
            if(response.status>=200 && response<300){
                return response.json()
            }else{
                if(on_error){
                    on_error(response,null)
                }
            }
        }).then(function(json){
            if(on_json) on_json(json);
        }).then(function(ex){
                if(on_error){
                    on_error(null,ex)
                }
        })
}

render(React.createElement(RootRouterV), window.document.getElementById("react_root"));