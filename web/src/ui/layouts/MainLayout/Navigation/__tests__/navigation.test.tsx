import renderer from "react-test-renderer"
import Navigation from "../Navigation";

describe("<Navigation />", () => {
  it("renders navigation correctly", () => {
    const tree = renderer
      .create(<Navigation />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

});
