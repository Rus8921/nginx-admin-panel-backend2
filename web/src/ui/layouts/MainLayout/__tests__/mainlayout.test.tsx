import { createMemoryRouter } from "react-router";
import { RouterProvider } from "react-router-dom";
import renderer from "react-test-renderer"
import { MainLayout } from "../MainLayout";

describe("<MainLayout />", () => {
  it("renders main layout correctly", () => {
    const router = createMemoryRouter([
      {
        path: "/",
        element: (<MainLayout />),
        children: [{
          index: true,
          element: (<>index page mock</>)
        },
        {
          path: "websites",
          element: (<>websites page mock</>),
        },
        {
          path: "servers",
          element: (<>websites page mock</>),
        },
        {
          path: "permissions",
          element: (<>websites page mock</>),
        }]
      },
      {
        path: "/logout",
        element: (<>logout page mock</>),
      }])

    const tree = renderer.create(
      <RouterProvider router={router} />
    ).toJSON();
    expect(tree).toMatchSnapshot();
  })
});
