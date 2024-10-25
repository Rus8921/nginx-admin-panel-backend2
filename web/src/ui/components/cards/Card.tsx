import { HTMLProps, PropsWithChildren } from "react";

function Card(props: PropsWithChildren<HTMLProps<HTMLDivElement>>) {
  let { children } = props

  return (
    <div className="w-full p-8 bg-white border border-scndry-txt-clr rounded-2xl hover:border-main-clr hover:shadow-2xs active:border-scndry-clr active:shadow-none" tabIndex={0} {...props}>
      {children}
    </div>
  )
}

export default Card;