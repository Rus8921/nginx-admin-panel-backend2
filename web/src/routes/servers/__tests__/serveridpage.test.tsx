import renderer from "react-test-renderer";
import { ServerConfigPage } from "../ServerConfigPage";

describe("<ServerConfigPage />", () => {
  it("renders server page layout correctly", () => {
    const tree = renderer.create(<ServerConfigPage />).toJSON();
    expect(tree).toMatchSnapshot();
  });
});
