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
      .create(<WebsitesPage datasetId={1} />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders empty website page layout correctly", () => {
    const tree = renderer
      .create(<WebsitesPage datasetId={3} />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
