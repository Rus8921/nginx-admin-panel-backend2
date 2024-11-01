import renderer from "react-test-renderer"
import Status from "../Status";

describe("<Status />", () => {
  it("renders entity's active status correctly", () => {
    const tree = renderer
      .create(<Status status="active" />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders entity's inactive status correctly", () => {
    const tree = renderer
      .create(<Status status="inactive" />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});

