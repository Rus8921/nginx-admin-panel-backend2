import { createMemoryRouter, RouterProvider, useNavigate } from "react-router-dom";
import renderer from "react-test-renderer"
import { useUserStore } from "../../../stores/userStore";
import { IUser } from "../../../types/user";
import { LogoutPage } from "../LogoutPage";


describe("<LogoutPage />", () => {
  it("renders logout page correctly", () => {
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
      return (<LogoutPage />)
    }

    const router = createMemoryRouter([
      {
        index: true,
        path: "/logout",
        element: <LoggedInPageMock />,
      },
      {
        path: "/login",
        element: <>login page mock</>,
      }
    ])

    const tree = renderer
      .create(<RouterProvider router={router} />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });
});
