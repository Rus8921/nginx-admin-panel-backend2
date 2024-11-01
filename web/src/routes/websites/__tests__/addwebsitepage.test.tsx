import renderer from "react-test-renderer"
import { AddWebsitePage } from "../AddWebsitePage";

describe("<AddWebsitePage />", () => {
  it("renders add website page layout correctly", () => {
    const tree = renderer
      .create(<AddWebsitePage />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

});
