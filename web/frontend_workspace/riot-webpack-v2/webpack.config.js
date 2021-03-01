const webpack = require('webpack');
const path = require('path');

module.exports = {
    entry: {
        resume: './src/resume.js',
        admin: './src/admin.js'
    },
    output: {
        path: __dirname + '/public/js/',
        filename: '[name].bundle.js'
    },
    module: {
        rules: [
            {
                test: /\.tag$/,
                enforce: 'pre',
                exclude: /node_modules/,
                use: [
                    {
                        loader: 'riot-tag-loader',
                        options: {
                            //template: 'pug',   //template if needed
                            debug: true
                        }
                    }
                ]
            },
            {
                test: /\.js|\.tag$/,
                enforce: 'post',
                exclude: /node_modules/,
                use: [
                    {
                        loader: 'babel-loader',
                        options: {
                            presets: `es2015-riot`
                        }
                    }
                ]
            }
        ]
    },
    resolve: {
        extensions: ['.js', '.tag']
    },
    devServer: {
        contentBase: path.join(__dirname, 'public'),
        port: 3000,
        inline: true
    },
    plugins: [
        new webpack.ProvidePlugin({
            riot: 'riot',
            route: 'riot-route'
        })
    ]
};
