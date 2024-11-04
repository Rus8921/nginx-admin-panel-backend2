import renderer from "react-test-renderer"
import UserInfo from "../UserInfo";
import { IUser } from "../../../../../../types/user";
import { useUserStore } from "../../../../../../stores/userStore";
import { createMemoryRouter, useNavigate } from "react-router";
import { RouterProvider } from "react-router-dom";

describe("<UserInfo />", () => {
  const userTestData: IUser = {
    email: "test@email.com",
    login: "testlogin",
    token: "testtoken",
    refreshToken: "testrefreshToken",
    tokenExpiresMilliseconds: 1000,
    role: 1,
  }

  function LoggedInPageMock() {
    const navigate = useNavigate()
    useUserStore.getState().login({ user: userTestData });
    return (<UserInfo user={userTestData} logout={() => {
      navigate("/logout");
    }} />)
  }

  const router = createMemoryRouter([
    {
      index: true,
      path: "/",
      element: (<LoggedInPageMock />)
    },
    {
      path: "/logout",
      element: (<>logout page mock</>),
    }
  ])

  it("renders user info div correctly", () => {
    const tree = renderer
      .create(<RouterProvider router={router} />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
