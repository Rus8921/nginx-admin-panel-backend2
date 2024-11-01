import renderer from "react-test-renderer"
import LangSwitch from "../LangSwitch";

describe("<LangSwitch />", () => {
  it("renders language switcher correctly", () => {
    const tree = renderer
      .create(<LangSwitch />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

});
