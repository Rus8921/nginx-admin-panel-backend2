import renderer from "react-test-renderer"
import { MainLayout } from "../MainLayout";

describe("<MainLayout />", () => {
  it("renders main layout correctly", () => {
    const tree = renderer
      .create(<MainLayout />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

});
