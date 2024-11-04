import renderer from "react-test-renderer"
import { CommonButton } from "../CommonButton";

describe("<CommonButton />", () => {
  it("renders submit button correctly", () => {
    const tree = renderer
      .create(<CommonButton isSubmit={true} type="blueBgWhiteText" onClick={() => { }}>test button text</CommonButton>)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders common action button correctly", () => {
    const tree = renderer
      .create(<CommonButton isSubmit={false} type="blueBgWhiteText" onClick={() => { }}>test button text</CommonButton>)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders dangerous submit button correctly", () => {
    const tree = renderer
      .create(<CommonButton isSubmit={true} type="redBgWhiteText" onClick={() => { }}>test button text</CommonButton>)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders common dangerous button correctly", () => {
    const tree = renderer
      .create(<CommonButton isSubmit={false} type="redBgWhiteText" onClick={() => { }}>test button text</CommonButton>)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });


  it("renders upload and submit button correctly", () => {
    const tree = renderer
      .create(<CommonButton isSubmit={true} type="transparentBgMainText" onClick={() => { }}>test button text</CommonButton>)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders common upload button correctly", () => {
    const tree = renderer
      .create(<CommonButton isSubmit={false} type="transparentBgMainText" onClick={() => { }}>test button text</CommonButton>)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
