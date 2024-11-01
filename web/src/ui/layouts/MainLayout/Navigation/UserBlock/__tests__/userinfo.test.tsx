import renderer from "react-test-renderer"
import UserInfo from "../UserInfo";
import { IUser } from "../../../../../../types/user";

describe("<UserInfo />", () => {
  const userTestData: IUser = {
    email: "test@email.com",
    login: "testlogin",
    token: "testtoken",
    refreshToken: "testrefreshToken",
    tokenExpiresMilliseconds: 1000,
    role: 1,
  }

  it("renders user info div correctly", () => {
    const tree = renderer
      .create(<UserInfo user={userTestData} />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

});
