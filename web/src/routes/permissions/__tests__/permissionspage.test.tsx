import renderer from "react-test-renderer"
import { PermissionsPage } from "../PermissionsPage";

describe("<PermissionsPage />", () => {
  it("renders permissions page layout correctly", () => {
    const tree = renderer
      .create(<PermissionsPage />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

});
