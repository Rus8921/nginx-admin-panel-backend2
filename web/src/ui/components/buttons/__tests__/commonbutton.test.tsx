import renderer from "react-test-renderer"
import { CommonButton } from "../CommonButton";

describe("<CommonButton />", () => {
  it("renders submit button correctly", () => {
    const tree = renderer
      .create(<CommonButton buttonText="test button text" isSubmit={true} type="blueBgWhiteText" onClick={() => { }} />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders common action button correctly", () => {
    const tree = renderer
      .create(<CommonButton buttonText="test button text" isSubmit={false} type="blueBgWhiteText" onClick={() => { }} />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders dangerous submit button correctly", () => {
    const tree = renderer
      .create(<CommonButton buttonText="test button text" isSubmit={true} type="redBgWhiteText" onClick={() => { }} />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders common dangerous button correctly", () => {
    const tree = renderer
      .create(<CommonButton buttonText="test button text" isSubmit={false} type="redBgWhiteText" onClick={() => { }} />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });


  it("renders upload and submit button correctly", () => {
    const tree = renderer
      .create(<CommonButton buttonText="test button text" isSubmit={true} type="transparentBgMainText" onClick={() => { }} />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders common upload button correctly", () => {
    const tree = renderer
      .create(<CommonButton buttonText="test button text" isSubmit={false} type="transparentBgMainText" onClick={() => { }} />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
