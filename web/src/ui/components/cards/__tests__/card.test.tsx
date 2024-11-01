import { getByTestId, render } from "@testing-library/react";
import renderer from "react-test-renderer"
import Card from "../Card";

describe("<Card />", () => {
  it("renders card correctly", () => {
    const tree = renderer
      .create(<Card>test content</Card>)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders card child correctly", () => {
    const childElement = (<p data-test-id="child">test content</p>)
    const cardWithChild = renderer.create(<Card data-test-id="parent">{childElement}</Card>).toTree();

    expect(cardWithChild?.props.children).toBe(childElement);
  });
});
