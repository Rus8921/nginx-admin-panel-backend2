import { ComponentProps } from "react";

function Card({ children, className, isClickable, ...rest }: ComponentProps<"div"> & { isClickable: boolean }) {
  return (
    <div className={"w-full p-8 bg-white border border-scndry-txt-clr rounded-2xl " + (isClickable && "hover:border-main-clr hover:shadow-2xs active:border-scndry-clr active:shadow-none ") + className} tabIndex={isClickable ? 0 : undefined} {...rest}>
      {children}
    </div>
  )
}

export default Card;