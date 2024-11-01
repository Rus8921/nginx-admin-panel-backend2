import renderer from "react-test-renderer"
import { ServerIdPage } from "../ServerIdPage";

describe("<ServerIdPage />", () => {
  it("renders server page layout correctly", () => {
    const tree = renderer
      .create(<ServerIdPage />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
