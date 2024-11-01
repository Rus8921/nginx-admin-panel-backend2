import renderer from "react-test-renderer"
import { LoginPage } from "../LoginPage";

describe("<LoginPage />", () => {
  it("renders login page layout correctly", () => {
    const tree = renderer
      .create(<LoginPage />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
