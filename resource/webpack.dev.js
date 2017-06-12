const Merge = require('webpack-merge');
var webpack = require('webpack');
const CommonConfig = require('./webpack.common.js');

module.exports = function(env) {
    return Merge(CommonConfig,{
          devtool: 'cheap-module-source-map',
          plugins: [
            new webpack.LoaderOptionsPlugin({
              minimize: true,
              debug: false
            }),
            new webpack.DefinePlugin({
              'process.env': {
                'NODE_ENV': JSON.stringify('development')
              }
            })
        ]
    })
  
}
