import renderer from "react-test-renderer"
import TextButton from "../TextButton";

describe("<TextButton />", () => {
  it("renders text button correctly", () => {
    const tree = renderer
      .create(<TextButton onClick={() => { }}>Test Button</TextButton>)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

});
