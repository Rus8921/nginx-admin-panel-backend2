import { ServerConfigPage } from "../ServerConfigPage";
import renderer from "react-test-renderer";

describe("<ServerConfigPage />", () => {
  it("renders servers page layout correctly", () => {
    const tree = renderer.create(<ServerConfigPage />).toJSON();
    expect(tree).toMatchSnapshot();
  });
});
