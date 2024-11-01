import renderer from "react-test-renderer"
import { WebsiteIdPage } from "../WebsiteIdPage";

describe("<WebsiteIdPage />", () => {
  it("renders website page layout", () => {
    const tree = renderer
      .create(<WebsiteIdPage />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
