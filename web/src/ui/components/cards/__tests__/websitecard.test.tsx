import { getByTestId, render } from "@testing-library/react";
import renderer from "react-test-renderer"
import { WebsiteInterface } from "../../../../types";
import WebsiteCard from "../WebsiteCard";

const websiteTestData: WebsiteInterface = {
  id: 1,
  name: "test name",
  url: "test/url",
  ipCount: 3,
  upstreamsCount: 2,
  status: "active",
}

describe("<WebsiteCard />", () => {
  it("renders website's card correctly", () => {
    const tree = renderer
      .create(<WebsiteCard data={websiteTestData} />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  // it("renders card child correctly", () => {
  //   const cardWithChild = renderer.create(<WebsiteCard data={websiteTestData} />).toJSON();

  //   expect(cardWithChild).toMatch(`*websiteTestData.name*`);
  // });
});