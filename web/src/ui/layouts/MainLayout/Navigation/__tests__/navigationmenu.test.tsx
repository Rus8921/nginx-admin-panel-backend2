import renderer from "react-test-renderer"
import NavigationMenu from "../NavigationMenu";

describe("<NavigationMenu />", () => {
  it("renders navigation sidebar correctly", () => {
    const tree = renderer
      .create(<NavigationMenu />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

});
