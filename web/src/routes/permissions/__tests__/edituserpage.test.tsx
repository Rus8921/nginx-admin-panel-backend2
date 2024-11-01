import renderer from "react-test-renderer"
import { EditUserPage } from "../users/EditUserPage";

describe("<EditUserPage />", () => {
  it("renders edit user page layout correctly", () => {
    const tree = renderer
      .create(<EditUserPage />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

});
