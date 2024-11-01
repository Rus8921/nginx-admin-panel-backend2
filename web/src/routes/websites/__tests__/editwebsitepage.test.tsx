import renderer from "react-test-renderer"
import { EditWebsitePage } from "../EditWebsitePage";

describe("<EditWebsitePage />", () => {
  it("renders edit website page layout correctly", () => {
    const tree = renderer
      .create(<EditWebsitePage />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

});
