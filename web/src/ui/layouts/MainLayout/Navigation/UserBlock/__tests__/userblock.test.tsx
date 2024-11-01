import renderer from "react-test-renderer"
import UserBlock from "../UserBlock";

describe("<UserBlock />", () => {

  it("renders user block correctly", () => {
    const tree = renderer
      .create(<UserBlock />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

});
