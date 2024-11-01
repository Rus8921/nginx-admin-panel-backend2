import renderer from "react-test-renderer"
import { AddServerPage } from "../AddServerPage";

describe("<AddServerPage />", () => {
  it("renders add server page layout correctly", () => {
    const tree = renderer
      .create(<AddServerPage />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

});
