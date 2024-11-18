import renderer from "react-test-renderer"
import { WebsitesPage } from "../WebsitesPage";

describe("<WebsitesPage />", () => {
  it("renders websites page layout with websites correctly", () => {
    const tree = renderer
      .create(<WebsitesPage />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders empty website page layout correctly", () => {
    const tree = renderer
      .create(<WebsitesPage datasetId="test_empty" />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders empty website page layout correctly", () => {
    const tree = renderer
      .create(<WebsitesPage datasetId="undefined" />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
