import renderer from "react-test-renderer"
import { EditServerPage } from "../EditServerPage";

describe("<EditServerPage />", () => {
  it("renders edit server page layout correctly", () => {
    const tree = renderer
      .create(<EditServerPage />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

});
