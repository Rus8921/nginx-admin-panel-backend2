import "./Card.css"
import { HTMLProps, PropsWithChildren } from "react";

function Card(props: PropsWithChildren<HTMLProps<HTMLDivElement>>) {
  let { width, children } = props

  return (
    <div className="Card" width={width} tabIndex={0} {...props}>
      {children}
    </div>
  )
}

export default Card;