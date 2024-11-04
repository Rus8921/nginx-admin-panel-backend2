import renderer from "react-test-renderer"
import UserBlock from "../UserBlock";
import { createMemoryRouter } from "react-router";
import { ProtectedRoute } from "../../../../../../routes/ProtectedRoute";
import { RouterProvider } from "react-router-dom";
import { useUserStore } from "../../../../../../stores/userStore";

describe("<UserBlock />", () => {
  it("renders non-logged in user block correctly", () => {
    const router = createMemoryRouter([
      {
        index: true,
        path: "/",
        element: (<UserBlock />)
      },
      {
        path: "/logout",
        element: (<>logout page mock</>),
      }
    ])

    const tree = renderer
      .create(<RouterProvider router={router} />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

  it("renders user block correctly", () => {
    function LoggedInPageMock() {
      useUserStore.getState().login({
        user: {
          email: "test@email.com",
          login: "testlogin",
          token: "testtoken",
          refreshToken: "testrefreshToken",
          tokenExpiresMilliseconds: 1000,
          role: 1,
        },
      });
      return (<UserBlock />)
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

    const tree = renderer
      .create(<RouterProvider router={router} />)
      .toJSON();
    expect(tree).toMatchSnapshot();
  });

});
