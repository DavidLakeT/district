const { defineConfig } = require('@vue/cli-service')
const MiniCssExtractPlugin = require("mini-css-extract-plugin");


module.exports = defineConfig({
  transpileDependencies: true
});

module.exports = {
  configureWebpack: {
    plugins: [new MiniCssExtractPlugin()],
  }
};