import renderer from "react-test-renderer"
import { ServersPage } from "../ServersPage";

describe("<ServersPage />", () => {
  it("renders servers page layout correctly", () => {
    const tree = renderer
      .create(<ServersPage />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

});
