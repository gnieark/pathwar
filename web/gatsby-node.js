/**
 * Implement Gatsby's Node APIs in this file.
 *
 * See: https://www.gatsbyjs.org/docs/node-apis/
 */

exports.onCreatePage = async ({ page, actions }) => {
  const { createPage } = actions;

  const appIsRoot = process.env.APP_ROOT === "true";
  const matchingRegex = appIsRoot ? /^\/$/ : /^\/app/;
  const matchingPath = appIsRoot ? `/*` : `/app/*`;

  // page.matchPath is a special key that's used for matching pages
  // only on the client.
  if (page.path.match(matchingRegex)) {
    page.matchPath = matchingPath;

    // Update the page.
    createPage(page);
  }
};

exports.onCreateWebpackConfig = ({ stage, actions }) => {
  if (stage === "develop") {
    actions.setWebpackConfig({
      devtool: "eval-source-map"
    });
  }
};
