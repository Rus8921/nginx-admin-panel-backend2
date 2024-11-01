import renderer from "react-test-renderer"
import Logo from "../Logo";

describe("<Logo />", () => {
  it("renders logo correctly", () => {
    const tree = renderer
      .create(<Logo />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

});
