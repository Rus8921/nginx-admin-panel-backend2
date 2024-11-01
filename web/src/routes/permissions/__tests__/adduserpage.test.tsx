import renderer from "react-test-renderer"
import { AddUserPage } from "../users/AddUserPage";

describe("<AddUserPage />", () => {
  it("renders add user page layout correctly", () => {
    const tree = renderer
      .create(<AddUserPage />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

});
