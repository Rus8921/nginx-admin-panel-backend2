import renderer from "react-test-renderer"
import { SidebarItem } from "../SidebarItem";
import { Layout } from "react-feather";
import { createMemoryRouter, RouterProvider, redirect } from "react-router-dom";
import { ProtectedRoute } from "../../../../../routes/ProtectedRoute";


describe("<SidebarItem />", () => {
  it("renders sidebar navigation item correctly", () => {
    const router = createMemoryRouter([
      {
        path: "/",
        element: (<SidebarItem ItemIcon={Layout} title="Websites" route="websites" />),
        children: [{
          index: true,
          element: (<>index page mock</>)
        },
        {
          path: "websites",
          element: (<ProtectedRoute>websites page mock</ProtectedRoute>),
        }]
      }])

    const tree = renderer.create(
      <RouterProvider router={router} />
    ).toJSON();
    expect(tree).toMatchSnapshot();
  });


  it("renders current sidebar navigation item correctly", () => {
    const router = createMemoryRouter([
      {
        path: "/",
        element: (<SidebarItem ItemIcon={Layout} title="Main" route="" />)
      }])

    const tree = renderer.create(
      <RouterProvider router={router} />
    ).toJSON();
    expect(tree).toMatchSnapshot();
  });
});
