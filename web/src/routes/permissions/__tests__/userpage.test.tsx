import renderer from "react-test-renderer"
import { UserPage } from "../users/UserPage";

describe("<UserPage />", () => {
  it("renders user page layout correctly", () => {
    const tree = renderer
      .create(<UserPage />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
