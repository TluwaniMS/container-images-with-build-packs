const { app } = require("./app");
const port = process.env.PORT || 3008;

app.listen(port, () => {
  console.log(`server is running on port: http://localhost:${port}/graphql`);
});
