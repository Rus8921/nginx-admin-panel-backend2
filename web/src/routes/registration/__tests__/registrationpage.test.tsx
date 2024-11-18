import { createMemoryRouter } from "react-router";
import { RouterProvider } from "react-router-dom";
import renderer from "react-test-renderer"
import { RegistrationPage, loader as registrationLoader } from "../RegistrationPage";

describe("<RegistrationPage />", () => {
  it("renders registration page layout correctly", () => {
    const router = createMemoryRouter([
      {
        index: true,
        path: "/",
        element: (<RegistrationPage />)
      },
      {
        path: "/registration",
        loader: registrationLoader,
        element: <RegistrationPage />,
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
