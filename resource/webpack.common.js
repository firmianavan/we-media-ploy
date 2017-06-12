var path = require('path'),
  webpack = require('webpack');

var rootPath = path.resolve(__dirname, '..'), // 项目根目录
  src = path.join(rootPath, 'src'); // 开发源码目录
 // env = process.env.NODE_ENV.trim(); // 当前环境

var commonPath = {
  rootPath: rootPath,
  dist: path.join(rootPath, 'dist'), // build 后输出目录
  staticDir: path.join(rootPath, 'static') // 无需处理的静态资源目录
};

module.exports = {
  entry: {
    'vendor': ["whatwg-fetch","react-dom","react"],
    'app': './index.js'
  },

  output: {
    path: path.join(__dirname, '/dist'),
    filename: '[name].bundle.js',
    sourceMapFilename: '[name].map'
  },

  resolve: {
    extensions: ['.ts', '.js', '.json'],
    modules: [path.join(__dirname, '../react_src'), 'node_modules']
  },

  module: {
    rules: [
      {
        test: /\.css$/,
        use: ['to-string-loader', 'css-loader']
      },
      {
        test: /\.(jpg|png|gif)$/,
        use: 'file-loader'
      },
      {
        test: /\.jsx$/,
        exclude: /(node_modules|dist)/,
        loader: 'babel-loader',
        query: {
          presets: ['es2015','react']
        }
      },
    ]
  },
  
  plugins: [

    new webpack.optimize.CommonsChunkPlugin({
      name: "vendor",

      // filename: "vendor.js"
      // (Give the chunk a different name)

      minChunks: Infinity,
      // (with more entries, this ensures that no other module
      //  goes into the vendor chunk)
    })

  ]
}