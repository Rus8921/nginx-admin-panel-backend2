import { PropsWithChildren } from "react";

function FlatTableCell({ children }: PropsWithChildren) {
  return (<td>{children}</td>);
}

export default FlatTableCell;
