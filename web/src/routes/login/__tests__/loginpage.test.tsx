import { createMemoryRouter } from "react-router";
import { RouterProvider } from "react-router-dom";
import renderer from "react-test-renderer"
import { LoginPage, loader as loginLoader } from "../LoginPage";

describe("<LoginPage />", () => {
  it("renders login page layout correctly", () => {
    const router = createMemoryRouter([
      {
        index: true,
        path: "/",
        element: (<LoginPage />)
      },
      {
        path: "/login",
        loader: loginLoader,
        element: <LoginPage />,
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
