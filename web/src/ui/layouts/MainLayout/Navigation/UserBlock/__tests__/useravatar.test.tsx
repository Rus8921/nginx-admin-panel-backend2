import renderer from "react-test-renderer"
import UserAvatar from "../UserAvatar";

describe("<UserAvatar />", () => {
  it("renders user avatar correctly", () => {
    const tree = renderer
      .create(<UserAvatar login="tester" />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

});
