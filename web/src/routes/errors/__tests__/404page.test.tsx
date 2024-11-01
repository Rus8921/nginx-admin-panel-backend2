import renderer from "react-test-renderer"
import { NotFoundErrorPage } from "../404";

describe("<NotFoundErrorPage />", () => {
  it("renders 404 page layout correctly", () => {
    const tree = renderer
      .create(<NotFoundErrorPage />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
