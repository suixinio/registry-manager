/* eslint-disable @typescript-eslint/no-var-requires */
const path = require('path');
const {
    override,
    addWebpackModuleRule,
    addWebpackPlugin,
    addWebpackAlias,
} = require('customize-cra');
const ArcoWebpackPlugin = require('@arco-plugins/webpack-react');
const addLessLoader = require('customize-cra-less-loader');
const setting = require('./src/settings.json');
const CopyPlugin = require('copy-webpack-plugin');

module.exports = {
    webpack: override(
        addLessLoader({
            lessLoaderOptions: {
                lessOptions: {},
            },
        }),
        addWebpackModuleRule({
            test: /\.svg$/,
            loader: '@svgr/webpack',
        }),
        addWebpackPlugin(
            new ArcoWebpackPlugin({
                theme: '@arco-themes/react-arco-pro',
                modifyVars: {
                    'arcoblue-6': setting.themeColor,
                },
            })
        ),
        addWebpackPlugin(
            new CopyPlugin({
                patterns: [
                    {
                        from: 'package.json',
                        to: 'version.json',
                        transform(content, path) {
                            let contentJson = JSON.parse(content);
                            return JSON.stringify({
                                version: contentJson.version,
                                name: contentJson.name,
                            })
                        },
                    },
                ]
            }),
        ),
        addWebpackAlias({
            '@': path.resolve(__dirname, 'src'),
        })
    ),
};
